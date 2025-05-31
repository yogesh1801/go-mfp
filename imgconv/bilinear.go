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

	for s := 0; s < slen; s++ {
		draw.Draw(simg, srect, image.Black, image.ZP, draw.Over)
		simg.Set(s, 0, color.White)
		draw.BiLinear.Scale(dimg, drect, simg, srect, draw.Over, nil)

		for d := 0; d < dlen; d++ {
			w := float32(dimg.Gray16At(d, 0).Y) / 0xffff
			if w != 0 {
				cff := scaleCoeff{S: s, D: d, W: w}
				coeffs = append(coeffs, cff)
			}
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
