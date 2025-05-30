// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image scaling via bi-linear interpolation

package imgconv

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// scaleCoeff defines, how source rows or columns contribute
// into the destination:
//
//	for _, sc := range coeffs {
//	   dst[sc.D] += src[S] * sc.W
//	}
type scaleCoeff struct {
	S, D int     // Source->Destination positions
	W    float32 // Source weight
}

// String returns string representation of the scaleCoeff, for debugging.
func (sc scaleCoeff) String() string {
	return fmt.Sprintf("%d->%d: %g", sc.S, sc.D, sc.W)
}

// makeScaleCoefficients prepares coefficients for image
// scaling (vertical or horizontal) with changing image
// dimensions range from [0...slen) to [0...dlen).
func makeScaleCoefficients(slen, dlen int) []scaleCoeff {
	switch {
	case slen == 0 || dlen == 0:
		return nil

	case slen == dlen:
		coeffs := make([]scaleCoeff, slen)
		for x := 0; x < slen; x++ {
			sc := scaleCoeff{S: x, D: x, W: 1.0}
			coeffs[x] = sc
		}
		return coeffs

	case slen == 1:
		coeffs := make([]scaleCoeff, dlen)
		for x := 0; x < dlen; x++ {
			sc := scaleCoeff{S: 0, D: x, W: 1.0}
			coeffs[x] = sc
		}
		return coeffs

	case dlen == 1:
		coeffs := make([]scaleCoeff, slen)
		W := 1.0 / float32(slen)
		for x := 0; x < slen; x++ {
			sc := scaleCoeff{S: x, D: 0, W: W}
			coeffs[x] = sc
		}
		return coeffs

	case dlen > slen:
		return makeScaleCoefficientsUpscale(slen, dlen)
	}

	// dlen < sleb
	return makeScaleCoefficientsDownscale(slen, dlen)
}

// makeScaleCoefficientsUpscale returns bi-linear coefficients
// for image upscale.
func makeScaleCoefficientsUpscale(slen, dlen int) []scaleCoeff {
	assert.Must(slen < dlen)

	// We have (slen - 1) intervals between source pixels and
	// (dlen - 1) intervals between destination pixels.
	//
	// If we bring source and destination coordinates into the
	// [0...(slen-1) * (dlen-1)) space, we can avoid rounding
	// errors here.
	space := uint64(slen-1) * uint64(dlen-1) * 2
	srcstep := uint64(dlen-1) * 2 // space / (slen - 1)
	dststep := uint64(slen-1) * 2 // space / (dlen - 1)

	println("space", space, "srcstep", srcstep, "dststep", dststep)
	println("steps:", "src", space/srcstep, "dst", space/dststep)

	coeffs := make([]scaleCoeff, 0, slen+dlen)
	for dst := uint64(0); dst <= space; dst += dststep {
		println("=== dst", dst/dststep, dst)
		// Handle the special case when source and destination
		// points matches 1:1
		if dst%srcstep == 0 {
			println("1:1")
			S := int(dst / srcstep)
			D := int(dst / dststep)
			sc := scaleCoeff{S: S, D: D, W: 1.0}
			coeffs = append(coeffs, sc)
			continue
		}

		// Find preceding and succeeding source points
		prec := generic.LowerDivisibleBy(dst, srcstep)
		succ := generic.UpperDivisibleBy(dst, srcstep)

		println(prec, "-", dst, "-", succ)

		// Compute weights
		precDistance := float32(dst - prec)
		succDistance := float32(succ - dst)

		wholeDistance := precDistance + succDistance

		println("dist", precDistance/wholeDistance, succDistance/wholeDistance)

		precWeight := succDistance / wholeDistance
		succWeight := precDistance / wholeDistance

		println("weight", precWeight, succWeight)

		// Generate coefficients
		D := int(dst / dststep)
		coeffs = append(coeffs,
			scaleCoeff{S: int(prec / srcstep), D: D, W: precWeight},
			scaleCoeff{S: int(succ / srcstep), D: D, W: succWeight})
	}

	return coeffs
}

// makeScaleCoefficientsUpscale returns scaling coefficients
// for image downscale.
func makeScaleCoefficientsDownscale(slen, dlen int) []scaleCoeff {
	assert.Must(slen > dlen)

	// Here we have the following picture:
	//   <-Sn-> - range, covered by the n-th source point
	//   <Dm>   - range, covered by the m-th destination point
	//
	//   S0->|<-S1->|<-S2->|<-S3->|<-S4->|<-S5
	//   D9>|<D1>|>D2>|>D3>|>D3>|>D4>|>D5>|>D6
	//
	// We align first and last source and destination points
	// to each other.
	//
	// Each destination point may receive contributions from
	// one or many source point, and we compute each source
	// point contribution proportionally to the range it
	// covers in the whole range taken by the corresponging
	// destination point.
	//
	// We use the [0...(slen-1) * (dlen-1) * 2) coordinate
	// space, which allows to represent source and destination
	// points coordinates and coordinates of the middle point
	// between succeeding source points or succeeding destination
	// point without rounding errors (notice * 2 - it allows
	// to represent middles).
	space := uint64(slen-1) * uint64(dlen-1) * 2
	srcstep := uint64(dlen-1) * 2 // space / (slen - 1)
	dststep := uint64(slen-1) * 2 // space / (dlen - 1)

	println("slen", slen, "dlen", dlen)
	println("space", space, "srcstep", srcstep, "dststep", dststep)
	println("steps:", "src", space/srcstep, "dst", space/dststep)

	coeffs := make([]scaleCoeff, 0, slen+dlen)
	for dst := uint64(0); dst <= space; dst += dststep {
		println("=== dst", dst/dststep, "(", dst, ")")

		// Compute indices and [0...space] positions of the
		// source first and last source points that contribute
		// into the destination:
		//
		//   firstIdx   - index of the first source point
		//   lastIdx    - index of the last source point
		//   firstStart - starting position of the first source point
		//   lastStart  - starting position of the last source point
		//
		// Corner cases require special care.
		firstIdx := 0
		lastIdx := slen - 1
		firstStart := uint64(0)
		lastStart := space - srcstep/2

		if dst != 0 {
			firstStart = generic.LowerDivisibleBy(
				dst-dststep/2+srcstep/2, srcstep)
			firstIdx = int(firstStart / srcstep)
			firstStart -= srcstep / 2
		}

		if dst != space {
			lastStart = generic.UpperDivisibleBy(
				dst+dststep/2-srcstep/2, srcstep)
			lastIdx = int(lastStart / srcstep)
			lastStart -= srcstep / 2
		}

		println(dst/dststep, "firstIdx", firstIdx, "lastIdx", lastIdx)
		println(dst/dststep, "firstStart", firstStart, "lastStart", lastStart)

		// Compute overlap ranges
		dstStart := dst
		dstRange := dststep
		if dst == 0 || dst == space {
			dstRange /= 2
		}

		firstRange := srcstep / 2
		if dst != 0 {
			dstStart -= dststep / 2
			firstRange = firstStart + srcstep - dstStart
		}

		lastRange := srcstep / 2
		if dst != space {
			lastRange = dstStart + dstRange - lastStart
		}

		println(dst/dststep, "dstStart", dstStart)
		println(dst/dststep, "firstRange", firstRange, "lastRange", lastRange, "dstRange", dstRange)

		assert.Must(dstRange ==
			firstRange+lastRange+
				srcstep*uint64(lastIdx-firstIdx-1))

		// Compute weights
		firstWeight := float32(firstRange) / float32(dstRange)
		lastWeight := float32(lastRange) / float32(dstRange)
		midWeight := float32(srcstep) / float32(dstRange)

		// Generate coefficients
		D := int(dst / dststep)
		coeffs = append(coeffs,
			scaleCoeff{S: firstIdx, D: D, W: firstWeight})

		for S := firstIdx + 1; S < lastIdx; S++ {
			coeffs = append(coeffs,
				scaleCoeff{S: S, D: D, W: midWeight})
		}

		coeffs = append(coeffs,
			scaleCoeff{S: lastIdx, D: D, W: lastWeight})
	}

	return coeffs
}
