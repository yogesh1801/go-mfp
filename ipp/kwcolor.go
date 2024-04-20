// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Color Names
// PWG5101.1: 4. Color Names

package ipp

import (
	"strings"
)

// KwColor represents a standard color name
// PWG5101.1: 4. Color Names
type KwColor string

// Standard values for KwColor
const (
	KwColorNoColor        KwColor = "no-color"
	KwColorBlack          KwColor = "black"
	KwColorClearBlack     KwColor = "clear-black"
	KwColorLightBlack     KwColor = "light-black"
	KwColorBlue           KwColor = "blue"
	KwColorClearBlue      KwColor = "clear-blue"
	KwColorDarkBlue       KwColor = "dark-blue"
	KwColorLightBlue      KwColor = "light-blue"
	KwColorBrown          KwColor = "brown"
	KwColorClearBrown     KwColor = "clear-brown"
	KwColorDarkBrown      KwColor = "dark-brown"
	KwColorLightBrown     KwColor = "light-brown"
	KwColorBuff           KwColor = "buff"
	KwColorClearBuff      KwColor = "clear-buff"
	KwColorDarkBuff       KwColor = "dark-buff"
	KwColorLightBuff      KwColor = "light-buff"
	KwColorCyan           KwColor = "cyan"
	KwColorClearCyan      KwColor = "clear-cyan"
	KwColorDarkCyan       KwColor = "dark-cyan"
	KwColorLightCyan      KwColor = "light-cyan"
	KwColorGold           KwColor = "gold"
	KwColorClearGold      KwColor = "clear-gold"
	KwColorDarkGold       KwColor = "dark-gold"
	KwColorLightGold      KwColor = "light-gold"
	KwColorGoldenrod      KwColor = "goldenrod"
	KwColorClearGoldenrod KwColor = "clear-goldenrod"
	KwColorDarkGoldenrod  KwColor = "dark-goldenrod"
	KwColorLightGoldenrod KwColor = "light-goldenrod"
	KwColorGray           KwColor = "gray"
	KwColorClearGray      KwColor = "clear-gray"
	KwColorDarkGray       KwColor = "dark-gray"
	KwColorLightGray      KwColor = "light-gray"
	KwColorGreen          KwColor = "green"
	KwColorClearGreen     KwColor = "clear-green"
	KwColorDarkGreen      KwColor = "dark-green"
	KwColorLightGreen     KwColor = "light-green"
	KwColorIvory          KwColor = "ivory"
	KwColorClearIvory     KwColor = "clear-ivory"
	KwColorDarkIvory      KwColor = "dark-ivory"
	KwColorLightIvory     KwColor = "light-ivory"
	KwColorMagenta        KwColor = "magenta"
	KwColorClearMagenta   KwColor = "clear-magenta"
	KwColorDarkMagenta    KwColor = "dark-magenta"
	KwColorLightMagenta   KwColor = "light-magenta"
	KwColorMustard        KwColor = "mustard"
	KwColorClearMustard   KwColor = "clear-mustard"
	KwColorDarkMustard    KwColor = "dark-mustard"
	KwColorLightMustard   KwColor = "light-mustard"
	KwColorOrange         KwColor = "orange"
	KwColorClearOrange    KwColor = "clear-orange"
	KwColorDarkOrange     KwColor = "dark-orange"
	KwColorLightOrange    KwColor = "light-orange"
	KwColorPink           KwColor = "pink"
	KwColorClearPink      KwColor = "clear-pink"
	KwColorDarkPink       KwColor = "dark-pink"
	KwColorLightPink      KwColor = "light-pink"
	KwColorRed            KwColor = "red"
	KwColorClearRed       KwColor = "clear-red"
	KwColorDarkRed        KwColor = "dark-red"
	KwColorLightRed       KwColor = "light-red"
	KwColorSilver         KwColor = "silver"
	KwColorClearSilver    KwColor = "clear-silver"
	KwColorDarkSilver     KwColor = "dark-silver"
	KwColorLightSilver    KwColor = "light-silver"
	KwColorTurquoise      KwColor = "turquoise"
	KwColorClearTurquoise KwColor = "clear-turquoise"
	KwColorDarkTurquoise  KwColor = "dark-turquoise"
	KwColorLightTurquoise KwColor = "light-turquoise"
	KwColorViolet         KwColor = "violet"
	KwColorClearViolet    KwColor = "clear-violet"
	KwColorDarkViolet     KwColor = "dark-violet"
	KwColorLightViolet    KwColor = "light-violet"
	KwColorWhite          KwColor = "white"
	KwColorClearWhite     KwColor = "clear-white"
	KwColorYellow         KwColor = "yellow"
	KwColorClearYellow    KwColor = "clear-yellow"
	KwColorDarkYellow     KwColor = "dark-yellow"
	KwColorLightYellow    KwColor = "light-yellow`"
)

// LocalizedName returns Localized Name for standard colors.
//
// If color is not known, the empty string ("") will be returned.
func (cn KwColor) LocalizedName() string {
	// For all colors in the table, Localized Name can be computed
	// as follows:
	//     "part1-part2" -> "Part1 Part2"
	//
	// The only exception is "no-color" which maps into "Transparent"
	if cn == "no-color" {
		return "Transparent"
	}

	_, found := kwColorByName[cn]
	if found {
		parts := strings.Split(string(cn), "-")
		name := ""
		for n, part := range parts {
			if n != 0 {
				name += " "
			}
			name += strings.ToUpper(part[:1]) + part[1:]
		}
		return name
	}

	return ""
}

// kwColorByName maps standard color names to their RGBA value
var kwColorByName = map[KwColor]Color{
	KwColorNoColor:        0xFFFFFF00,
	KwColorBlack:          0x000000FF,
	KwColorClearBlack:     0x0000007F,
	KwColorLightBlack:     0x808080FF,
	KwColorBlue:           0x0000FFFF,
	KwColorClearBlue:      0x0000FF7F,
	KwColorDarkBlue:       0x00008BFF,
	KwColorLightBlue:      0xADD8E6FF,
	KwColorBrown:          0xA52A2AFF,
	KwColorClearBrown:     0xA52A2A7F,
	KwColorDarkBrown:      0x5C4033FF,
	KwColorLightBrown:     0x9966FFFF,
	KwColorBuff:           0xF0DC82FF,
	KwColorClearBuff:      0xF0DC827F,
	KwColorDarkBuff:       0x976638FF,
	KwColorLightBuff:      0xECD9B0FF,
	KwColorCyan:           0x00FFFFFF,
	KwColorClearCyan:      0x00FFFF7F,
	KwColorDarkCyan:       0x008B8BFF,
	KwColorLightCyan:      0xE0FFFFFF,
	KwColorGold:           0xFFD700FF,
	KwColorClearGold:      0xFFD7007F,
	KwColorDarkGold:       0xEEBC1DFF,
	KwColorLightGold:      0xF1E5ACFF,
	KwColorGoldenrod:      0xDAA520FF,
	KwColorClearGoldenrod: 0xDAA5207F,
	KwColorDarkGoldenrod:  0xB8860BFF,
	KwColorLightGoldenrod: 0xFFEC8BFF,
	KwColorGray:           0x808080FF,
	KwColorClearGray:      0x8080807F,
	KwColorDarkGray:       0x404040FF,
	KwColorLightGray:      0xD3D3D3FF,
	KwColorGreen:          0x008000FF,
	KwColorClearGreen:     0x0080007F,
	KwColorDarkGreen:      0x006400FF,
	KwColorLightGreen:     0x90EE90FF,
	KwColorIvory:          0xFFFFF0FF,
	KwColorClearIvory:     0xFFFFF07F,
	KwColorDarkIvory:      0xF2E58FFF,
	KwColorLightIvory:     0xFFF8C9FF,
	KwColorMagenta:        0xFF00FFFF,
	KwColorClearMagenta:   0xFF00FF7F,
	KwColorDarkMagenta:    0x8B008BFF,
	KwColorLightMagenta:   0xFF77FFFF,
	KwColorMustard:        0xFFDB58FF,
	KwColorClearMustard:   0xFFDB587F,
	KwColorDarkMustard:    0x7C7C40FF,
	KwColorLightMustard:   0xEEDD62FF,
	KwColorOrange:         0xFFA500FF,
	KwColorClearOrange:    0xFFA5007F,
	KwColorDarkOrange:     0xFF8C00FF,
	KwColorLightOrange:    0xD9A465FF,
	KwColorPink:           0xFFC0CBFF,
	KwColorClearPink:      0xFFC0CB7F,
	KwColorDarkPink:       0xE75480FF,
	KwColorLightPink:      0xFFB6C1FF,
	KwColorRed:            0xFF0000FF,
	KwColorClearRed:       0xFF00007F,
	KwColorDarkRed:        0x8B0000FF,
	KwColorLightRed:       0xFF3333FF,
	KwColorSilver:         0xC0C0C0FF,
	KwColorClearSilver:    0xC0C0C07F,
	KwColorDarkSilver:     0xAFAFAFFF,
	KwColorLightSilver:    0xE1E1E1FF,
	KwColorTurquoise:      0x30D5C8FF,
	KwColorClearTurquoise: 0x30D5C87F,
	KwColorDarkTurquoise:  0x00CED1FF,
	KwColorLightTurquoise: 0xAFE4DEFF,
	KwColorViolet:         0xEE82EEFF,
	KwColorClearViolet:    0xEE82EE7F,
	KwColorDarkViolet:     0x9400D3FF,
	KwColorLightViolet:    0x7A5299FF,
	KwColorWhite:          0xFFFFFFFF,
	KwColorClearWhite:     0xFFFFFF7F,
	KwColorYellow:         0xFFFF00FF,
	KwColorClearYellow:    0xFFFF007F,
	KwColorDarkYellow:     0xFFCC00FF,
	KwColorLightYellow:    0xFFFFE0FF,
}
