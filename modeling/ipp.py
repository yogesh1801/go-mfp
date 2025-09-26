# MFP - Miulti-Function Printers and scanners toolkit
# Printer and scanner modeling.
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# IPP-related definitions

# Range represents range of integers
class Range (TypedDict):
    Lower: int
    Upper: int

# Resolution represents image resolution
    Xres: int
    Yres: int
    Units: str

# attrs is the model-settable variable that defines the
# IPP printer attributes
attrs = None
