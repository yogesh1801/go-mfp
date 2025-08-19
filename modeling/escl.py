# MFP - Miulti-Function Printers and scanners toolkit
# Printer and scanner modeling.
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# eSCL-related definitions

from uuid import UUID
from typing import TypedDict

# Range represents range of integers
class Range (TypedDict):
    Min: int
    Max: int
    Normal: int
    Step: int

# Resolution represents printing or scanning resolution
class Resolution(TypedDict):
    XResolution: int
    YResolution: int

# Region represents scanning region
class Region (TypedDict):
    Width: int
    Height: int
    XOffset: int
    YOffset: int

# Intent keywords:
Document = 'Document'
TextAndGraphic = 'TextAndGraphic'
Photo = 'Photo'
Preview = 'Preview'
Object = 'Object'
BusinessCard = 'BusinessCard'

# ContentType keywords:
Photo = 'Photo'
Text = 'Text'
TextAndPhoto = 'TextAndPhoto'
LineArt = 'LineArt'
Magazine = 'Magazine'
Halftone = 'Halftone'
Auto = 'Auto'

# InputSource keywords:
Platen = 'Platen'
Feeder = 'Feeder'
Camera = 'Camera'

# ColorMode keywords:
BlackAndWhite1 = 'BlackAndWhite1'
Grayscale8 = 'Grayscale8'
Grayscale16 = 'Grayscale16'
RGB24 = 'RGB24'
RGB48 = 'RGB48'

# CcdChannel keywords:
Red = 'Red'
Green = 'Green'
Blue = 'Blue'
NTSC = 'NTSC'
GrayCcd = 'GrayCcd'
GrayCcdEmulated = 'GrayCcdEmulated'

# BinaryRendering keywords:
Halftone = 'Halftone'
Threshold = 'Threshold'

# FeedDirection keywords:
LongEdgeFeed = 'LongEdgeFeed'
ShortEdgeFeed = 'ShortEdgeFeed'

# ScanSettings is the eSCL scan request:
#
# POST /{root}/ScanJobs       - to start scanning
# PUT  /{root}/ScanBufferInfo - to estimate actual scanning parameters
class ScanSettings(TypedDict):
    Version: str
    Intent: str
    ScanRegions: list
    DocumentFormat: str
    DocumentFormatExt: str
    ContentType: str
    InputSource: str
    XResolution: Resolution
    YResolution: Resolution
    ColorMode: str
    ColorSpace: str
    CcdChannel: str
    BinaryRendering: str
    Duplex: str
    FeedDirection: str
    Brightness: int
    CompressionFactor: int
    Contrast: int
    Gamma: int
    Highlight: int
    NoiseRemoval: int
    Shadow: int
    Sharpen: int
    Threshold: int
    BlankPageDetection: bool
    BlankPageDetectionAndRemoval: bool

# caps is the model-settable variable that defines the
# eSCL scanner capabilities
caps: dict = None

