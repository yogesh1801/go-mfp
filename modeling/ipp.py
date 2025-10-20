# MFP - Miulti-Function Printers and scanners toolkit
# Printer and scanner modeling.
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# IPP-related definitions

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

# attrs is the model-settable variable that defines the
# IPP printer attributes
attrs = None
