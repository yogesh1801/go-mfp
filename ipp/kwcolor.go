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

	_, found := colorNameTable[cn]
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

var colorNameTable = map[KwColor]Color{
	"no-color":        0xFFFFFF00,
	"black":           0x000000FF,
	"clear-black":     0x0000007F,
	"light-black":     0x808080FF,
	"blue":            0x0000FFFF,
	"clear-blue":      0x0000FF7F,
	"dark-blue":       0x00008BFF,
	"light-blue":      0xADD8E6FF,
	"brown":           0xA52A2AFF,
	"clear-brown":     0xA52A2A7F,
	"dark-brown":      0x5C4033FF,
	"light-brown":     0x9966FFFF,
	"buff":            0xF0DC82FF,
	"clear-buff":      0xF0DC827F,
	"dark-buff":       0x976638FF,
	"light-buff":      0xECD9B0FF,
	"cyan":            0x00FFFFFF,
	"clear-cyan":      0x00FFFF7F,
	"dark-cyan":       0x008B8BFF,
	"light-cyan":      0xE0FFFFFF,
	"gold":            0xFFD700FF,
	"clear-gold":      0xFFD7007F,
	"dark-gold":       0xEEBC1DFF,
	"light-gold":      0xF1E5ACFF,
	"goldenrod":       0xDAA520FF,
	"clear-goldenrod": 0xDAA5207F,
	"dark-goldenrod":  0xB8860BFF,
	"light-goldenrod": 0xFFEC8BFF,
	"gray":            0x808080FF,
	"clear-gray":      0x8080807F,
	"dark-gray":       0x404040FF,
	"light-gray":      0xD3D3D3FF,
	"green":           0x008000FF,
	"clear-green":     0x0080007F,
	"dark-green":      0x006400FF,
	"light-green":     0x90EE90FF,
	"ivory":           0xFFFFF0FF,
	"clear-ivory":     0xFFFFF07F,
	"dark-ivory":      0xF2E58FFF,
	"light-ivory":     0xFFF8C9FF,
	"magenta":         0xFF00FFFF,
	"clear-magenta":   0xFF00FF7F,
	"dark-magenta":    0x8B008BFF,
	"light-magenta":   0xFF77FFFF,
	"mustard":         0xFFDB58FF,
	"clear-mustard":   0xFFDB587F,
	"dark-mustard":    0x7C7C40FF,
	"light-mustard":   0xEEDD62FF,
	"orange":          0xFFA500FF,
	"clear-orange":    0xFFA5007F,
	"dark-orange":     0xFF8C00FF,
	"light-orange":    0xD9A465FF,
	"pink":            0xFFC0CBFF,
	"clear-pink":      0xFFC0CB7F,
	"dark-pink":       0xE75480FF,
	"light-pink":      0xFFB6C1FF,
	"red":             0xFF0000FF,
	"clear-red":       0xFF00007F,
	"dark-red":        0x8B0000FF,
	"light-red":       0xFF3333FF,
	"silver":          0xC0C0C0FF,
	"clear-silver":    0xC0C0C07F,
	"dark-silver":     0xAFAFAFFF,
	"light-silver":    0xE1E1E1FF,
	"turquoise":       0x30D5C8FF,
	"clear-turquoise": 0x30D5C87F,
	"dark-turquoise":  0x00CED1FF,
	"light-turquoise": 0xAFE4DEFF,
	"violet":          0xEE82EEFF,
	"clear-violet":    0xEE82EE7F,
	"dark-violet":     0x9400D3FF,
	"light-violet":    0x7A5299FF,
	"white":           0xFFFFFFFF,
	"clear-white":     0xFFFFFF7F,
	"yellow":          0xFFFF00FF,
	"clear-yellow":    0xFFFF007F,
	"dark-yellow":     0xFFCC00FF,
	"light-yellow":    0xFFFFE0FF,
}
