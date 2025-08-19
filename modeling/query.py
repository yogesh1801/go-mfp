# MFP - Miulti-Function Printers and scanners toolkit
# Printer and scanner modeling.
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# HTTP-level queries

from dataclasses import dataclass
from http.client import HTTPMessage

# Query represents the HTTP query being processed
@dataclass
class Query:
    request: HTTPMessage = None
    response: HTTPMessage = None

