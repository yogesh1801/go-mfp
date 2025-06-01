// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image scaling via bi-linear interpolation

package imgconv

import (
	"image"
	"image/color"
	"sort"

	"github.com/OpenPrinting/go-mfp/util/generic"
	"golang.org/x/image/draw"
)

// scaleCoeff defines, how source rows or columns contribute
// into the destination:
//
//	for _, sc := range coeffs {
//	   dst[sc.D] += src[sc.S] * sc.W
//	}
type scaleCoeff struct {
	S, D int     // Source->Destination positions
	W    float32 // Source weight
}

// makeScaleCoefficients prepares coefficients for image
// scaling (vertical or horizontal) with changing image
// dimensions range from [0...slen) to [0...dlen).
func makeScaleCoefficients(slen, dlen int) []scaleCoeff {
	coeffs := make([]scaleCoeff, 0, slen+dlen)

	srect := image.Rect(0, 0, slen, 1)
	drect := image.Rect(0, 0, dlen, 1)

	simg := image.NewGray16(srect)
	dimg := image.NewGray16(drect)

	// Here we use draw.BiLinear.Scale to obtain its
	// source to destination contributions coefficients.
	//
	// We do it, using the following algorithm:
	//   for S := 0; S < slen; S ++ {
	//     - create slen * 1 image
	//     - set point at S to White, all others to Black
	//     - scale this image into the dlen * 1 size
	//     - points at the scaled image gives S->D coefficients
	//   }
	//
	// Unfortunately, this is quadratic algorithm and for the
	// big dimensions it runs very slow.
	//
	// To optimize, we assume that source points placed at the
	// some big enough distance are completely independent, so
	// we can simultaneously and independently use source
	// index N, N+distance, N+distance*2 and so on.
	//
	// Unfortunately, we don't know the safe distance in advance.
	//
	// To estimate, we make the following assumptions:
	//   - for downscaling, (slen/dlen) * 3 is safe assumption
	//   - for upscaling, 4 is safe assumption
	distance := 0

	if slen > dlen {
		distance = (slen + dlen - 1) / dlen
		distance *= 3
	} else {
		distance = 4
	}

	groups := (slen + distance - 1) / distance
	for off := 0; off < distance; off++ {
		// Fill image with Black color, once for all groups
		draw.Draw(simg, srect, image.Black, image.ZP, draw.Over)

		// Set one pixel at the each group to While
		for group := 0; group < groups; group++ {
			s := group*distance + off
			if s >= slen {
				continue
			}

			simg.Set(s, 0, color.White)
		}

		// Scale slen*1 -> dlen*1
		draw.BiLinear.Scale(dimg, drect, simg, srect, draw.Over, nil)

		// Scan the scaled image. D-coordinate matches
		// position at that image, S-coordinates we must
		// recover by counting Black gaps between colored
		// groups of points.
		s := off
		prev := uint16(0)
		for d := 0; d < dlen; d++ {
			c := dimg.Gray16At(d, 0).Y
			if c != 0 {
				w := float32(c) / 0xffff
				cff := scaleCoeff{S: s, D: d, W: w}
				coeffs = append(coeffs, cff)
			} else if prev != 0 {
				s += distance
			}

			prev = c
		}
	}

	sort.SliceStable(coeffs, func(i, j int) bool {
		return coeffs[i].D < coeffs[j].D
	})

	return coeffs
}

// scaleCoefficientsHistorySize computes how much history is used by
// the scale coefficients.
//
// It returns:
//   - 0 if the coefficients never use sources in decreasing order,
//   - 1 if they need to seek one step back from the maximum reached source index,
//   - and so on.
func scaleCoefficientsHistorySize(coeffs []scaleCoeff) int {
	Smax := 0
	hist := 0

	for _, sc := range coeffs {
		Smax = generic.Max(Smax, sc.S)
		hist = generic.Max(hist, Smax-sc.S)
	}

	return hist
}
