# MFP - Miulti-Function Printers and scanners toolkit
# Printer and scanner modeling.
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# IPP-related definitions

from dataclasses import dataclass
from enum import Enum

# IPP tags
class TAG(Enum):
    # Delimiters
    ZERO = 0x00
    END = 0x03

    # Groups of attributes
    OPERATION = 0x01
    JOB = 0x02
    PRINTER = 0x04
    UNSUPPORTED_GROUP = 0x05
    SUBSCRIPTION = 0x06
    EVENT_NOTIFICATION = 0x07
    RESOURCE = 0x08
    DOCUMENT = 0x09
    SYSTEM = 0x0a

    # Special values
    UNSUPPORTED_VALUE = 0x10
    DEFAULT = 0x11
    UNKNOWN = 0x12
    NOVALUE = 0x13
    NOTSETTABLE = 0x15
    DELETEATTR = 0x16
    ADMINDEFINE = 0x17

    # Values
    INTEGER = 0x21
    BOOLEAN = 0x22
    ENUM = 0x23
    STRING = 0x30
    DATE = 0x31
    RESOLUTION = 0x32
    RANGE = 0x33
    TEXTLANG = 0x35
    NAMELANG = 0x36
    TEXT = 0x41
    NAME = 0x42
    KEYWORD = 0x44
    URI = 0x45
    URISCHEME = 0x46
    CHARSET = 0x47
    LANGUAGE = 0x48
    MIMETYPE = 0x49
    EXTENSION = 0x7f

    # Collections
    BEGIN_COLLECTION = 0x34
    END_COLLECTION = 0x37
    MEMBERNAME = 0x4a

    # Formatting
    def __repr__ (self):
        return 'ipp.TAG.' + self.name

# IPP_TAG_UNSUPPORTED_VALUE
class UNSUPPORTED_VALUE:
    def __repr__ (self):
        return 'ipp.UNSUPPORTED_VALUE'

# IPP_TAG_DEFAULT
class DEFAULT:
    def __repr__ (self):
        return 'ipp.DEFAULT'

# IPP_TAG_UNKNOWN
class UNKNOWN:
    def __repr__ (self):
        return 'ipp.UNKNOWN()'

# IPP_TAG_NOVALUE
class NOVALUE:
    def __repr__ (self):
        return 'ipp.NOVALUE()'

# IPP_TAG_NOTSETTABLE
class NOTSETTABLE:
    def __repr__ (self):
        return 'ipp.NOTSETTABLE()'

# IPP_TAG_DELETEATTR
class DELETEATTR:
    def __repr__ (self):
        return 'ipp.DELETEATTR()'

# IPP_TAG_ADMINDEFINE
class ADMINDEFINE:
    def __repr__ (self):
        return 'ipp.ADMINDEFINE()'

# IPP_TAG_INTEGER
class INTEGER(int):
    def __repr__ (self):
        return 'ipp.INTEGER(' + repr(int(self)) + ')'

# IPP_TAG_BOOLEAN
@dataclass
class BOOLEAN:
    Val : bool

    def __bool__ (self):
        return self.Val

    def __repr__ (self):
        return 'ipp.BOOLEAN(' + repr(bool(self.Val)) + ')'

# IPP_TAG_ENUM
class ENUM(int):
    def __repr__ (self):
        return 'ipp.ENUM(' + repr(int(self)) + ')'

# IPP_TAG_STRING
class STRING(str):
    def __repr__ (self):
        return 'ipp.STRING(' + repr(str(self)) + ')'

# IPP_TAG_DATE
class DATE(str):
    def __repr__ (self):
        return 'ipp.DATE(' + repr(str(self)) + ')'

# IPP_TAG_RESOLUTION
@dataclass
class RESOLUTION:
    X: int
    Y: int
    Units: str

    def __repr__ (self):
        return 'ipp.RESOLUTION(' + repr(self.X) + ', ' + repr(self.Y) + ', ' + repr(self.Units) + ')'

# IPP_TAG_RANGE
@dataclass
class RANGE:
    Min: int
    Max: int

    def __repr__ (self):
        return 'ipp.RANGE(' + repr(self.Min) + ', ' + repr(self.Max) + ')'

# IPP_TAG_TEXTLANG
@dataclass
class TEXTLANG:
    Text: str
    Lang: str

    def __repr__ (self):
        return 'ipp.TEXTLANG(' + repr(self.Text) + ', ' + repr(self.Lang) + ')'

# IPP_TAG_NAMELANG
@dataclass
class NAMELANG:
    Name: str
    Lang: str

    def __repr__ (self):
        return 'ipp.NAMELANG(' + repr(self.Name) + ', ' + repr(self.Lang) + ')'

# IPP_TAG_TEXT
class TEXT(str):
    def __repr__ (self):
        return 'ipp.TEXT(' + repr(str(self)) + ')'

# IPP_TAG_NAME
class NAME(str):
    def __repr__ (self):
        return 'ipp.NAME(' + repr(str(self)) + ')'

# IPP_TAG_KEYWORD
class KEYWORD(str):
    def __repr__ (self):
        return 'ipp.KEYWORD(' + repr(str(self)) + ')'

# IPP_TAG_URI
class URI(str):
    def __repr__ (self):
        return 'ipp.URI(' + repr(str(self)) + ')'

# IPP_TAG_URISCHEME
class URISCHEME(str):
    def __repr__ (self):
        return 'ipp.URISCHEME(' + repr(str(self)) + ')'

# IPP_TAG_CHARSET
class CHARSET(str):
    def __repr__ (self):
        return 'ipp.CHARSET(' + repr(str(self)) + ')'

# IPP_TAG_LANGUAGE
class LANGUAGE(str):
    def __repr__ (self):
        return 'ipp.LANGUAGE(' + repr(str(self)) + ')'

# IPP_TAG_MIMETYPE
class MIMETYPE(str):
    def __repr__ (self):
        return 'ipp.MIMETYPE(' + repr(str(self)) + ')'

# attrs is the model-settable variable that defines the
# IPP printer attributes
attrs = None
