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

func (sc scaleCoeff) String() string {
	return fmt.Sprintf("%d->%d: %g", sc.S, sc.D, sc.W)
}

func makeScaleCoefficients(slen, dlen int) []scaleCoeff {
	// Handle special cases
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

	return makeScaleCoefficientsDownscale(slen, dlen)

	// We have (slen - 1) intervals between source pixels and
	// (dlen - 1) intervals between destination pixels.
	//
	// If we bring source and destination coordinates into the
	// [0...(slen-1) * (dlen-1)) space, we can avoid rounding
	// errors here.
	end := uint64(slen-1) * uint64(dlen-1)
	srcstep := uint64(dlen - 1) // end / (slen - 1)
	dststep := uint64(slen - 1) // end / (dlen - 1)

	println("end", end, "srcstep", srcstep, "dststep", dststep)

	coeffs := make([]scaleCoeff, 0, generic.Max(slen, dlen))
	for dst := uint64(0); dst <= end; dst += dststep {
		println("=== dst", dst/dststep)
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

		// Compute range of preceding and succeeding source points
		precLast := generic.LowerDivisibleBy(dst, srcstep)
		precFirst := precLast
		if dst != 0 {
			precFirst = generic.UpperDivisibleBy(dst-dststep, srcstep)
			precFirst = generic.Min(precFirst, precLast)
		}

		succFirst := generic.UpperDivisibleBy(dst, srcstep)
		succLast := succFirst
		if dst != end {
			succLast = generic.LowerDivisibleBy(dst+dststep, srcstep)
			succLast = generic.Max(succFirst, succLast)
		}

		println(precFirst, "-", precLast, dst, succFirst, "-", succLast)

		// Compute weights
		precMiddle := (float32(precFirst) + float32(precLast)) / 2
		precDistance := float32(dst) - precMiddle

		succMiddle := (float32(succFirst) + float32(succLast)) / 2
		succDistance := succMiddle - float32(dst)

		wholeDistance := succMiddle - precMiddle

		println("dist", precDistance/wholeDistance, succDistance/wholeDistance)

		precWeight := succDistance / wholeDistance
		if precFirst != precLast {
			cnt := (precLast - precFirst + srcstep) / srcstep
			precWeight /= float32(cnt)
		}

		succWeight := precDistance / wholeDistance
		if succFirst != succLast {
			cnt := (succLast - succFirst + srcstep) / srcstep
			succWeight /= float32(cnt)
		}

		println("weight", precWeight, succWeight)

		// Generate coefficients
		D := int(dst / dststep)
		for src := precFirst; src <= precLast; src += srcstep {
			S := int(src / srcstep)
			sc := scaleCoeff{S: S, D: D, W: precWeight}
			coeffs = append(coeffs, sc)
			println(S, "->", D, precWeight)
		}

		for src := succFirst; src <= succLast; src += srcstep {
			S := int(src / srcstep)
			sc := scaleCoeff{S: S, D: D, W: succWeight}
			coeffs = append(coeffs, sc)
			println(S, "->", D, succWeight)
		}
	}

	return coeffs
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

	//coeffs := make([]scaleCoeff, 0, slen+dlen)
	for dst := uint64(0); dst <= space; dst += dststep {
		println("=== dst", dst/dststep, "(", dst, ")")
		// Handle the special case when source and destination
		// points matches 1:1
		//if dst%srcstep == 0 {
		//	println("1:1")
		//	S := int(dst / srcstep)
		//	D := int(dst / dststep)
		//	sc := scaleCoeff{S: S, D: D, W: 1.0}
		//	coeffs = append(coeffs, sc)
		//	continue
		//}

		// Compute indices and [0...space] positions of the
		// source first and last source points that contribute
		// into the destination:
		//
		//   firstIdx - index of the first source point
		//   lastIdx  - index of the last source point
		//   firstPos - starting position of the first source point
		//   lastPos  - ending position of the last source point
		//
		// Corner cases require special care.
		firstIdx := 0
		lastIdx := slen - 1
		firstPos := uint64(0)
		lastPos := space - 1

		if dst != 0 {
			firstPos = generic.LowerDivisibleBy(
				dst-dststep/2+srcstep/2, srcstep)
			firstIdx = int(firstPos / srcstep)
			firstPos -= srcstep / 2
		}

		if dst != space {
			lastPos = generic.UpperDivisibleBy(
				dst+dststep/2-srcstep/2, srcstep)
			lastIdx = int(lastPos / srcstep)
			lastPos += srcstep/2 - 1
		}

		println(dst/srcstep, "idx:", firstIdx, "-", lastIdx)
		println(dst/srcstep, "pos:", firstPos, "-", lastPos)
		if false {
			println("firstIdx", firstIdx, "firstPos", firstPos)
			println("lastIdx", lastIdx, "lastPos", lastPos)
		}
	}

	return nil
}
