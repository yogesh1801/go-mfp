// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Standard media sizes

package ipp

// KwMedia represents standard media size. Used in many places
type KwMedia string

// Standard constants for KwMedia
const (
	KwMediaAsmeF                        KwMedia = "asme_f_28x40in"
	KwMediaChoiceIsoA4210x297mmNaLetter KwMedia = "choice_iso_a4_210x297mm_na_letter_8.5x11in"
	KwMediaIso2a0                       KwMedia = "iso_2a0_1189x1682mm"
	KwMediaIsoA0                        KwMedia = "iso_a0_841x1189mm"
	KwMediaIsoA1                        KwMedia = "iso_a1_594x841mm"
	KwMediaIsoA1x3                      KwMedia = "iso_a1x3_841x1783mm"
	KwMediaIsoA1x4                      KwMedia = "iso_a1x4_841x2378mm"
	KwMediaIsoA2                        KwMedia = "iso_a2_420x594mm"
	KwMediaIsoA2x3                      KwMedia = "iso_a2x3_594x1261mm"
	KwMediaIsoA2x4                      KwMedia = "iso_a2x4_594x1682mm"
	KwMediaIsoA2x5                      KwMedia = "iso_a2x5_594x2102mm"
	KwMediaIsoA3Extra                   KwMedia = "iso_a3-extra_322x445mm"
	KwMediaIsoA3                        KwMedia = "iso_a3_297x420mm"
	KwMediaIsoA0x3                      KwMedia = "iso_a0x3_1189x2523mm"
	KwMediaIsoA3x3                      KwMedia = "iso_a3x3_420x891mm"
	KwMediaIsoA3x4                      KwMedia = "iso_a3x4_420x1189mm"
	KwMediaIsoA3x5                      KwMedia = "iso_a3x5_420x1486mm"
	KwMediaIsoA3x6                      KwMedia = "iso_a3x6_420x1783mm"
	KwMediaIsoA3x7                      KwMedia = "iso_a3x7_420x2080mm"
	KwMediaIsoA4Extra                   KwMedia = "iso_a4-extra_235.5x322.3mm"
	KwMediaIsoA4Tab                     KwMedia = "iso_a4-tab_225x297mm"
	KwMediaIsoA4                        KwMedia = "iso_a4_210x297mm"
	KwMediaIsoA4x3                      KwMedia = "iso_a4x3_297x630mm"
	KwMediaIsoA4x4                      KwMedia = "iso_a4x4_297x841mm"
	KwMediaIsoA4x5                      KwMedia = "iso_a4x5_297x1051mm"
	KwMediaIsoA4x6                      KwMedia = "iso_a4x6_297x1261mm"
	KwMediaIsoA4x7                      KwMedia = "iso_a4x7_297x1471mm"
	KwMediaIsoA4x8                      KwMedia = "iso_a4x8_297x1682mm"
	KwMediaIsoA4x9                      KwMedia = "iso_a4x9_297x1892mm"
	KwMediaIsoA5Extra                   KwMedia = "iso_a5-extra_174x235mm"
	KwMediaIsoA5                        KwMedia = "iso_a5_148x210mm"
	KwMediaIsoA6                        KwMedia = "iso_a6_105x148mm"
	KwMediaIsoA7                        KwMedia = "iso_a7_74x105mm"
	KwMediaIsoA8                        KwMedia = "iso_a8_52x74mm"
	KwMediaIsoA9                        KwMedia = "iso_a9_37x52mm"
	KwMediaIsoA10                       KwMedia = "iso_a10_26x37mm"
	KwMediaIsoB0                        KwMedia = "iso_b0_1000x1414mm"
	KwMediaIsoB1                        KwMedia = "iso_b1_707x1000mm"
	KwMediaIsoB2                        KwMedia = "iso_b2_500x707mm"
	KwMediaIsoB3                        KwMedia = "iso_b3_353x500mm"
	KwMediaIsoB4                        KwMedia = "iso_b4_250x353mm"
	KwMediaIsoB5Extra                   KwMedia = "iso_b5-extra_201x276mm"
	KwMediaIsoB5                        KwMedia = "iso_b5_176x250mm"
	KwMediaIsoB6                        KwMedia = "iso_b6_125x176mm"
	KwMediaIsoB6c4                      KwMedia = "iso_b6c4_125x324mm"
	KwMediaIsoB7                        KwMedia = "iso_b7_88x125mm"
	KwMediaIsoB8                        KwMedia = "iso_b8_62x88mm"
	KwMediaIsoB9                        KwMedia = "iso_b9_44x62mm"
	KwMediaIsoB10                       KwMedia = "iso_b10_31x44mm"
	KwMediaIsoC0                        KwMedia = "iso_c0_917x1297mm"
	KwMediaIsoC1                        KwMedia = "iso_c1_648x917mm"
	KwMediaIsoC2                        KwMedia = "iso_c2_458x648mm"
	KwMediaIsoC3                        KwMedia = "iso_c3_324x458mm"
	KwMediaIsoC4                        KwMedia = "iso_c4_229x324mm"
	KwMediaIsoC5                        KwMedia = "iso_c5_162x229mm"
	KwMediaIsoC6                        KwMedia = "iso_c6_114x162mm"
	KwMediaIsoC6c5                      KwMedia = "iso_c6c5_114x229mm"
	KwMediaIsoC7                        KwMedia = "iso_c7_81x114mm"
	KwMediaIsoC7c6                      KwMedia = "iso_c7c6_81x162mm"
	KwMediaIsoC8                        KwMedia = "iso_c8_57x81mm"
	KwMediaIsoC9                        KwMedia = "iso_c9_40x57mm"
	KwMediaIsoC10                       KwMedia = "iso_c10_28x40mm"
	KwMediaIsoDl                        KwMedia = "iso_dl_110x220mm"
	KwMediaIsoID1                       KwMedia = "iso_id-1_53.98x85.6mm"
	KwMediaIsoRa0                       KwMedia = "iso_ra0_860x1220mm"
	KwMediaIsoRa1                       KwMedia = "iso_ra1_610x860mm"
	KwMediaIsoRa2                       KwMedia = "iso_ra2_430x610mm"
	KwMediaIsoRa3                       KwMedia = "iso_ra3_305x430mm"
	KwMediaIsoRa4                       KwMedia = "iso_ra4_215x305mm"
	KwMediaIsoSra0                      KwMedia = "iso_sra0_900x1280mm"
	KwMediaIsoSra1                      KwMedia = "iso_sra1_640x900mm"
	KwMediaIsoSra2                      KwMedia = "iso_sra2_450x640mm"
	KwMediaIsoSra3                      KwMedia = "iso_sra3_320x450mm"
	KwMediaIsoSra4                      KwMedia = "iso_sra4_225x320mm"
	KwMediaJisB0                        KwMedia = "jis_b0_1030x1456mm"
	KwMediaJisB1                        KwMedia = "jis_b1_728x1030mm"
	KwMediaJisB2                        KwMedia = "jis_b2_515x728mm"
	KwMediaJisB3                        KwMedia = "jis_b3_364x515mm"
	KwMediaJisB4                        KwMedia = "jis_b4_257x364mm"
	KwMediaJisB5                        KwMedia = "jis_b5_182x257mm"
	KwMediaJisB6                        KwMedia = "jis_b6_128x182mm"
	KwMediaJisB7                        KwMedia = "jis_b7_91x128mm"
	KwMediaJisB8                        KwMedia = "jis_b8_64x91mm"
	KwMediaJisB9                        KwMedia = "jis_b9_45x64mm"
	KwMediaJisB10                       KwMedia = "jis_b10_32x45mm"
	KwMediaJisExec                      KwMedia = "jis_exec_216x330mm"
	KwMediaJpnChou2                     KwMedia = "jpn_chou2_111.1x146mm"
	KwMediaJpnChou3                     KwMedia = "jpn_chou3_120x235mm"
	KwMediaJpnChou4                     KwMedia = "jpn_chou4_90x205mm"
	KwMediaJpnChou40                    KwMedia = "jpn_chou40_90x225mm"
	KwMediaJpnHagaki                    KwMedia = "jpn_hagaki_100x148mm"
	KwMediaJpnKahu                      KwMedia = "jpn_kahu_240x322.1mm"
	KwMediaJpnKaku1                     KwMedia = "jpn_kaku1_270x382mm"
	KwMediaJpnKaku2                     KwMedia = "jpn_kaku2_240x332mm"
	KwMediaJpnKaku3                     KwMedia = "jpn_kaku3_216x277mm"
	KwMediaJpnKaku4                     KwMedia = "jpn_kaku4_197x267mm"
	KwMediaJpnKaku5                     KwMedia = "jpn_kaku5_190x240mm"
	KwMediaJpnKaku7                     KwMedia = "jpn_kaku7_142x205mm"
	KwMediaJpnKaku8                     KwMedia = "jpn_kaku8_119x197mm"
	KwMediaJpnOufuku                    KwMedia = "jpn_oufuku_148x200mm"
	KwMediaJpnYou4                      KwMedia = "jpn_you4_105x235mm"
	KwMediaNa5x7                        KwMedia = "na_5x7_5x7in"
	KwMediaNa6x9                        KwMedia = "na_6x9_6x9in"
	KwMediaNa7x9                        KwMedia = "na_7x9_7x9in"
	KwMediaNa9x11                       KwMedia = "na_9x11_9x11in"
	KwMediaNa10x11                      KwMedia = "na_10x11_10x11in"
	KwMediaNa10x13                      KwMedia = "na_10x13_10x13in"
	KwMediaNa10x14                      KwMedia = "na_10x14_10x14in"
	KwMediaNa10x15                      KwMedia = "na_10x15_10x15in"
	KwMediaNa11x12                      KwMedia = "na_11x12_11x12in"
	KwMediaNa11x15                      KwMedia = "na_11x15_11x15in"
	KwMediaNa12x19                      KwMedia = "na_12x19_12x19in"
	KwMediaNaA2                         KwMedia = "na_a2_4.375x5.75in"
	KwMediaNaArchA                      KwMedia = "na_arch-a_9x12in"
	KwMediaNaArchB                      KwMedia = "na_arch-b_12x18in"
	KwMediaNaArchC                      KwMedia = "na_arch-c_18x24in"
	KwMediaNaArchD                      KwMedia = "na_arch-d_24x36in"
	KwMediaNaArchE2                     KwMedia = "na_arch-e2_26x38in"
	KwMediaNaArchE3                     KwMedia = "na_arch-e3_27x39in"
	KwMediaNaArchE                      KwMedia = "na_arch-e_36x48in"
	KwMediaNaBPlus                      KwMedia = "na_b-plus_12x19.17in"
	KwMediaNaC5                         KwMedia = "na_c5_6.5x9.5in"
	KwMediaNaC                          KwMedia = "na_c_17x22in"
	KwMediaNaD                          KwMedia = "na_d_22x34in"
	KwMediaNaE                          KwMedia = "na_e_34x44in"
	KwMediaNaEdp                        KwMedia = "na_edp_11x14in"
	KwMediaNaEurEdp                     KwMedia = "na_eur-edp_12x14in"
	KwMediaNaExecutive                  KwMedia = "na_executive_7.25x10.5in"
	KwMediaNaF                          KwMedia = "na_f_44x68in"
	KwMediaNaFanfoldEur                 KwMedia = "na_fanfold-eur_8.5x12in"
	KwMediaNaFanfoldUs                  KwMedia = "na_fanfold-us_11x14.875in"
	KwMediaNaFoolscap                   KwMedia = "na_foolscap_8.5x13in"
	KwMediaNaGovtLegal                  KwMedia = "na_govt-legal_8x13in"
	KwMediaNaGovtLetter                 KwMedia = "na_govt-letter_8x10in"
	KwMediaNaIndex3x5                   KwMedia = "na_index-3x5_3x5in"
	KwMediaNaIndex4x6Ext                KwMedia = "na_index-4x6-ext_6x8in"
	KwMediaNaIndex4x6                   KwMedia = "na_index-4x6_4x6in"
	KwMediaNaIndex5x8                   KwMedia = "na_index-5x8_5x8in"
	KwMediaNaInvoice                    KwMedia = "na_invoice_5.5x8.5in"
	KwMediaNaLedger                     KwMedia = "na_ledger_11x17in"
	KwMediaNaLegalExtra                 KwMedia = "na_legal-extra_9.5x15in"
	KwMediaNaLegal                      KwMedia = "na_legal_8.5x14in"
	KwMediaNaLetterExtra                KwMedia = "na_letter-extra_9.5x12in"
	KwMediaNaLetterPlus                 KwMedia = "na_letter-plus_8.5x12.69in"
	KwMediaNaLetter                     KwMedia = "na_letter_8.5x11in"
	KwMediaNaMonarch                    KwMedia = "na_monarch_3.875x7.5in"
	KwMediaNaNumber9                    KwMedia = "na_number-9_3.875x8.875in"
	KwMediaNaNumber10                   KwMedia = "na_number-10_4.125x9.5in"
	KwMediaNaNumber11                   KwMedia = "na_number-11_4.5x10.375in"
	KwMediaNaNumber12                   KwMedia = "na_number-12_4.75x11in"
	KwMediaNaNumber14                   KwMedia = "na_number-14_5x11.5in"
	KwMediaNaOficio                     KwMedia = "na_oficio_8.5x13.4in"
	KwMediaNaPersonal                   KwMedia = "na_personal_3.625x6.5in"
	KwMediaNaQuarto                     KwMedia = "na_quarto_8.5x10.83in"
	KwMediaNaSuperA                     KwMedia = "na_super-a_8.94x14in"
	KwMediaNaSuperB                     KwMedia = "na_super-b_13x19in"
	KwMediaNaWideFormat                 KwMedia = "na_wide-format_30x42in"
	KwMediaOe12x16                      KwMedia = "oe_12x16_12x16in"
	KwMediaOe14x17                      KwMedia = "oe_14x17_14x17in"
	KwMediaOe18x22                      KwMedia = "oe_18x22_18x22in"
	KwMediaOeA2plus                     KwMedia = "oe_a2plus_17x24in"
	KwMediaOeBusinessCard               KwMedia = "oe_business-card_2x3.5in"
	KwMediaOePhoto10r                   KwMedia = "oe_photo-10r_10x12in"
	KwMediaOePhoto12r                   KwMedia = "oe_photo-12r_12x15in"
	KwMediaOePhoto14x18                 KwMedia = "oe_photo-14x18_14x18in"
	KwMediaOePhoto16r                   KwMedia = "oe_photo-16r_16x20in"
	KwMediaOePhoto20r                   KwMedia = "oe_photo-20r_20x24in"
	KwMediaOePhoto22r                   KwMedia = "oe_photo-22r_22x29.5in"
	KwMediaOePhoto22x28                 KwMedia = "oe_photo-22x28_22x28in"
	KwMediaOePhoto24r                   KwMedia = "oe_photo-24r_24x31.5in"
	KwMediaOePhoto24x30                 KwMedia = "oe_photo-24x30_24x30in"
	KwMediaOePhoto30r                   KwMedia = "oe_photo-30r_30x40in"
	KwMediaOePhotoL                     KwMedia = "oe_photo-l_3.5x5in"
	KwMediaOePhotoS8r                   KwMedia = "oe_photo-s8r_8x12in"
	KwMediaOeSquarePhoto4x4in           KwMedia = "oe_square-photo_4x4in"
	KwMediaOeSquarePhoto5x5in           KwMedia = "oe_square-photo_5x5in"
	KwMediaOm16k184x260mm               KwMedia = "om_16k_184x260mm"
	KwMediaOm16k195x270mm               KwMedia = "om_16k_195x270mm"
	KwMediaOmBusinessCard55x85mm        KwMedia = "om_business-card_55x85mm"
	KwMediaOmBusinessCard55x91mm        KwMedia = "om_business-card_55x91mm"
	KwMediaOmCard                       KwMedia = "om_card_54x86mm"
	KwMediaOmDaiPaKai                   KwMedia = "om_dai-pa-kai_275x395mm"
	KwMediaOmDscPhoto                   KwMedia = "om_dsc-photo_89x119mm"
	KwMediaOmFolioSp                    KwMedia = "om_folio-sp_215x315mm"
	KwMediaOmFolio                      KwMedia = "om_folio_210x330mm"
	KwMediaOmInvite                     KwMedia = "om_invite_220x220mm"
	KwMediaOmItalian                    KwMedia = "om_italian_110x230mm"
	KwMediaOmJuuroKuKai                 KwMedia = "om_juuro-ku-kai_198x275mm"
	KwMediaOmLargePhoto                 KwMedia = "om_large-photo_200x300mm"
	KwMediaOmMediumPhoto                KwMedia = "om_medium-photo_130x180mm"
	KwMediaOmPaKai                      KwMedia = "om_pa-kai_267x389mm"
	KwMediaOmPhoto30x40                 KwMedia = "om_photo-30x40_300x400mm"
	KwMediaOmPhoto30x45                 KwMedia = "om_photo-30x45_300x450mm"
	KwMediaOmPhoto35x46                 KwMedia = "om_photo-35x46_350x460mm"
	KwMediaOmPhoto40x60                 KwMedia = "om_photo-40x60_400x600mm"
	KwMediaOmPhoto50x75                 KwMedia = "om_photo-50x75_500x750mm"
	KwMediaOmPhoto50x76                 KwMedia = "om_photo-50x76_500x760mm"
	KwMediaOmPhoto60x90                 KwMedia = "om_photo-60x90_600x900mm"
	KwMediaOmSmallPhoto                 KwMedia = "om_small-photo_100x150mm"
	KwMediaOmSquarePhoto                KwMedia = "om_square-photo_89x89mm"
	KwMediaOmWidePhoto                  KwMedia = "om_wide-photo_100x200mm"
	KwMediaPrc1                         KwMedia = "prc_1_102x165mm"
	KwMediaPrc2                         KwMedia = "prc_2_102x176mm"
	KwMediaPrc4                         KwMedia = "prc_4_110x208mm"
	KwMediaPrc6                         KwMedia = "prc_6_120x320mm"
	KwMediaPrc7                         KwMedia = "prc_7_160x230mm"
	KwMediaPrc8                         KwMedia = "prc_8_120x309mm"
	KwMediaPrc16k                       KwMedia = "prc_16k_146x215mm"
	KwMediaPrc32k                       KwMedia = "prc_32k_97x151mm"
	KwMediaRoc8k                        KwMedia = "roc_8k_10.75x15.5in"
	KwMediaRoc16k                       KwMedia = "roc_16k_7.75x10.75in"
)

// Size returns media size, width and height, in 1/100 mm.
// If size is not known, (-1, -1) is returned.
func (kw KwMedia) Size() (wid, hei int) {
	if size, ok := kwMediaByName[kw]; ok {
		return size.wid, size.hei
	}

	return -1, -1
}

// kwMediaSize represents media size, associated with the media name.
type kwMediaSize struct {
	wid, hei int // in 1/100 mm
}

// kwMediaByName maps media name to the corresponding media size
var kwMediaByName = map[KwMedia]kwMediaSize{
	KwMediaAsmeF:                        {71120, 101600},
	KwMediaChoiceIsoA4210x297mmNaLetter: {21590, 27940},
	KwMediaIso2a0:                       {118900, 168200},
	KwMediaIsoA0:                        {84100, 118900},
	KwMediaIsoA1:                        {59400, 84100},
	KwMediaIsoA1x3:                      {84100, 178300},
	KwMediaIsoA1x4:                      {84100, 237800},
	KwMediaIsoA2:                        {42000, 59400},
	KwMediaIsoA2x3:                      {59400, 126100},
	KwMediaIsoA2x4:                      {59400, 168200},
	KwMediaIsoA2x5:                      {59400, 210200},
	KwMediaIsoA3Extra:                   {32200, 44500},
	KwMediaIsoA3:                        {29700, 42000},
	KwMediaIsoA0x3:                      {118900, 252300},
	KwMediaIsoA3x3:                      {42000, 89100},
	KwMediaIsoA3x4:                      {42000, 118900},
	KwMediaIsoA3x5:                      {42000, 148600},
	KwMediaIsoA3x6:                      {42000, 178300},
	KwMediaIsoA3x7:                      {42000, 208000},
	KwMediaIsoA4Extra:                   {23550, 32230},
	KwMediaIsoA4Tab:                     {22500, 29700},
	KwMediaIsoA4:                        {21000, 29700},
	KwMediaIsoA4x3:                      {29700, 63000},
	KwMediaIsoA4x4:                      {29700, 84100},
	KwMediaIsoA4x5:                      {29700, 105100},
	KwMediaIsoA4x6:                      {29700, 126100},
	KwMediaIsoA4x7:                      {29700, 147100},
	KwMediaIsoA4x8:                      {29700, 168200},
	KwMediaIsoA4x9:                      {29700, 189200},
	KwMediaIsoA5Extra:                   {17400, 23500},
	KwMediaIsoA5:                        {14800, 21000},
	KwMediaIsoA6:                        {10500, 14800},
	KwMediaIsoA7:                        {7400, 10500},
	KwMediaIsoA8:                        {5200, 7400},
	KwMediaIsoA9:                        {3700, 5200},
	KwMediaIsoA10:                       {2600, 3700},
	KwMediaIsoB0:                        {100000, 141400},
	KwMediaIsoB1:                        {70700, 100000},
	KwMediaIsoB2:                        {50000, 70700},
	KwMediaIsoB3:                        {35300, 50000},
	KwMediaIsoB4:                        {25000, 35300},
	KwMediaIsoB5Extra:                   {20100, 27600},
	KwMediaIsoB5:                        {17600, 25000},
	KwMediaIsoB6:                        {12500, 17600},
	KwMediaIsoB6c4:                      {12500, 32400},
	KwMediaIsoB7:                        {8800, 12500},
	KwMediaIsoB8:                        {6200, 8800},
	KwMediaIsoB9:                        {4400, 6200},
	KwMediaIsoB10:                       {3100, 4400},
	KwMediaIsoC0:                        {91700, 129700},
	KwMediaIsoC1:                        {64800, 91700},
	KwMediaIsoC2:                        {45800, 64800},
	KwMediaIsoC3:                        {32400, 45800},
	KwMediaIsoC4:                        {22900, 32400},
	KwMediaIsoC5:                        {16200, 22900},
	KwMediaIsoC6:                        {11400, 16200},
	KwMediaIsoC6c5:                      {11400, 22900},
	KwMediaIsoC7:                        {8100, 11400},
	KwMediaIsoC7c6:                      {8100, 16200},
	KwMediaIsoC8:                        {5700, 8100},
	KwMediaIsoC9:                        {4000, 5700},
	KwMediaIsoC10:                       {2800, 4000},
	KwMediaIsoDl:                        {11000, 22000},
	KwMediaIsoID1:                       {5398, 8560},
	KwMediaIsoRa0:                       {86000, 122000},
	KwMediaIsoRa1:                       {61000, 86000},
	KwMediaIsoRa2:                       {43000, 61000},
	KwMediaIsoRa3:                       {30500, 43000},
	KwMediaIsoRa4:                       {21500, 30500},
	KwMediaIsoSra0:                      {90000, 128000},
	KwMediaIsoSra1:                      {64000, 90000},
	KwMediaIsoSra2:                      {45000, 64000},
	KwMediaIsoSra3:                      {32000, 45000},
	KwMediaIsoSra4:                      {22500, 32000},
	KwMediaJisB0:                        {103000, 145600},
	KwMediaJisB1:                        {72800, 103000},
	KwMediaJisB2:                        {51500, 72800},
	KwMediaJisB3:                        {36400, 51500},
	KwMediaJisB4:                        {25700, 36400},
	KwMediaJisB5:                        {18200, 25700},
	KwMediaJisB6:                        {12800, 18200},
	KwMediaJisB7:                        {9100, 12800},
	KwMediaJisB8:                        {6400, 9100},
	KwMediaJisB9:                        {4500, 6400},
	KwMediaJisB10:                       {3200, 4500},
	KwMediaJisExec:                      {21600, 33000},
	KwMediaJpnChou2:                     {11110, 14600},
	KwMediaJpnChou3:                     {12000, 23500},
	KwMediaJpnChou4:                     {9000, 20500},
	KwMediaJpnChou40:                    {9000, 22500},
	KwMediaJpnHagaki:                    {10000, 14800},
	KwMediaJpnKahu:                      {24000, 32210},
	KwMediaJpnKaku1:                     {27000, 38200},
	KwMediaJpnKaku2:                     {24000, 33200},
	KwMediaJpnKaku3:                     {21600, 27700},
	KwMediaJpnKaku4:                     {19700, 26700},
	KwMediaJpnKaku5:                     {19000, 24000},
	KwMediaJpnKaku7:                     {14200, 20500},
	KwMediaJpnKaku8:                     {11900, 19700},
	KwMediaJpnOufuku:                    {14800, 20000},
	KwMediaJpnYou4:                      {10500, 23500},
	KwMediaNa5x7:                        {12700, 17780},
	KwMediaNa6x9:                        {15240, 22860},
	KwMediaNa7x9:                        {17780, 22860},
	KwMediaNa9x11:                       {22860, 27940},
	KwMediaNa10x11:                      {25400, 27940},
	KwMediaNa10x13:                      {25400, 33020},
	KwMediaNa10x14:                      {25400, 35560},
	KwMediaNa10x15:                      {25400, 38100},
	KwMediaNa11x12:                      {27940, 30480},
	KwMediaNa11x15:                      {27940, 38100},
	KwMediaNa12x19:                      {30480, 48260},
	KwMediaNaA2:                         {11113, 14605},
	KwMediaNaArchA:                      {22860, 30480},
	KwMediaNaArchB:                      {30480, 45720},
	KwMediaNaArchC:                      {45720, 60960},
	KwMediaNaArchD:                      {60960, 91440},
	KwMediaNaArchE2:                     {66040, 96520},
	KwMediaNaArchE3:                     {68580, 99060},
	KwMediaNaArchE:                      {91440, 121920},
	KwMediaNaBPlus:                      {30480, 48692},
	KwMediaNaC5:                         {16510, 24130},
	KwMediaNaC:                          {43180, 55880},
	KwMediaNaD:                          {55880, 86360},
	KwMediaNaE:                          {86360, 111760},
	KwMediaNaEdp:                        {27940, 35560},
	KwMediaNaEurEdp:                     {30480, 35560},
	KwMediaNaExecutive:                  {18415, 26670},
	KwMediaNaF:                          {111760, 172720},
	KwMediaNaFanfoldEur:                 {21590, 30480},
	KwMediaNaFanfoldUs:                  {27940, 37783},
	KwMediaNaFoolscap:                   {21590, 33020},
	KwMediaNaGovtLegal:                  {20320, 33020},
	KwMediaNaGovtLetter:                 {20320, 25400},
	KwMediaNaIndex3x5:                   {7620, 12700},
	KwMediaNaIndex4x6Ext:                {15240, 20320},
	KwMediaNaIndex4x6:                   {10160, 15240},
	KwMediaNaIndex5x8:                   {12700, 20320},
	KwMediaNaInvoice:                    {13970, 21590},
	KwMediaNaLedger:                     {27940, 43180},
	KwMediaNaLegalExtra:                 {24130, 38100},
	KwMediaNaLegal:                      {21590, 35560},
	KwMediaNaLetterExtra:                {24130, 30480},
	KwMediaNaLetterPlus:                 {21590, 32233},
	KwMediaNaLetter:                     {21590, 27940},
	KwMediaNaMonarch:                    {9843, 19050},
	KwMediaNaNumber9:                    {9843, 22543},
	KwMediaNaNumber10:                   {10478, 24130},
	KwMediaNaNumber11:                   {11430, 26353},
	KwMediaNaNumber12:                   {12065, 27940},
	KwMediaNaNumber14:                   {12700, 29210},
	KwMediaNaOficio:                     {21590, 34036},
	KwMediaNaPersonal:                   {9208, 16510},
	KwMediaNaQuarto:                     {21590, 27508},
	KwMediaNaSuperA:                     {22708, 35560},
	KwMediaNaSuperB:                     {33020, 48260},
	KwMediaNaWideFormat:                 {76200, 106680},
	KwMediaOe12x16:                      {30480, 40640},
	KwMediaOe14x17:                      {35560, 43180},
	KwMediaOe18x22:                      {45720, 55880},
	KwMediaOeA2plus:                     {43180, 60960},
	KwMediaOeBusinessCard:               {5080, 8890},
	KwMediaOePhoto10r:                   {25400, 30480},
	KwMediaOePhoto12r:                   {30480, 38100},
	KwMediaOePhoto14x18:                 {35560, 45720},
	KwMediaOePhoto16r:                   {40640, 50800},
	KwMediaOePhoto20r:                   {50800, 60960},
	KwMediaOePhoto22r:                   {55880, 74930},
	KwMediaOePhoto22x28:                 {55880, 71120},
	KwMediaOePhoto24r:                   {60960, 80010},
	KwMediaOePhoto24x30:                 {60960, 76200},
	KwMediaOePhoto30r:                   {76200, 101600},
	KwMediaOePhotoL:                     {8890, 12700},
	KwMediaOePhotoS8r:                   {20320, 30480},
	KwMediaOeSquarePhoto4x4in:           {10160, 10160},
	KwMediaOeSquarePhoto5x5in:           {12700, 12700},
	KwMediaOm16k184x260mm:               {18400, 26000},
	KwMediaOm16k195x270mm:               {19500, 27000},
	KwMediaOmBusinessCard55x85mm:        {5500, 8500},
	KwMediaOmBusinessCard55x91mm:        {5500, 9100},
	KwMediaOmCard:                       {5400, 8600},
	KwMediaOmDaiPaKai:                   {27500, 39500},
	KwMediaOmDscPhoto:                   {8900, 11900},
	KwMediaOmFolioSp:                    {21500, 31500},
	KwMediaOmFolio:                      {21000, 33000},
	KwMediaOmInvite:                     {22000, 22000},
	KwMediaOmItalian:                    {11000, 23000},
	KwMediaOmJuuroKuKai:                 {19800, 27500},
	KwMediaOmLargePhoto:                 {20000, 30000},
	KwMediaOmMediumPhoto:                {13000, 18000},
	KwMediaOmPaKai:                      {26700, 38900},
	KwMediaOmPhoto30x40:                 {30000, 40000},
	KwMediaOmPhoto30x45:                 {30000, 45000},
	KwMediaOmPhoto35x46:                 {35000, 46000},
	KwMediaOmPhoto40x60:                 {40000, 60000},
	KwMediaOmPhoto50x75:                 {50000, 75000},
	KwMediaOmPhoto50x76:                 {50000, 76000},
	KwMediaOmPhoto60x90:                 {60000, 90000},
	KwMediaOmSmallPhoto:                 {10000, 15000},
	KwMediaOmSquarePhoto:                {8900, 8900},
	KwMediaOmWidePhoto:                  {10000, 20000},
	KwMediaPrc1:                         {10200, 16500},
	KwMediaPrc2:                         {10200, 17600},
	KwMediaPrc4:                         {11000, 20800},
	KwMediaPrc6:                         {12000, 32000},
	KwMediaPrc7:                         {16000, 23000},
	KwMediaPrc8:                         {12000, 30900},
	KwMediaPrc16k:                       {14600, 21500},
	KwMediaPrc32k:                       {9700, 15100},
	KwMediaRoc8k:                        {27305, 39370},
	KwMediaRoc16k:                       {19685, 27305},
}
