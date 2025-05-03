# GO-MFP

[![godoc.org](https://godoc.org/github.com/OpenPrinting/go-mfp?status.svg)](https://godoc.org/github.com/OpenPrinting/go-mfp)
![GitHub](https://img.shields.io/github/license/OpenPrinting/go-mfp)
[![Go Report Card](https://goreportcard.com/badge/github.com/OpenPrinting/go-mfp)](https://goreportcard.com/report/github.com/OpenPrinting/go-mfp)

**Warning:** this project is work in progress and under the intensive
development.

This project contains the foundation of the behavior-accurate simulator for
multi-function printers (MFPs) that supports at least IPP 2.x for
printing, eSCL and WSD for scanning, and DNS-SD and WS-Discovery for
device discovery.

This simulator will consist of a core simulation engine that provides
reference implementations of the aforementioned protocols, along with a
customization engine that allows for the expression of implementation
details specific to individual devices without the need to reimplement
common functionalities repeatedly.

For the purpose of implementation of the MFP core simulation engine it
contains many libraries (Go packages) that are aimed to be a part of the
simulator, but can be useful as a general-purpose libraries as well.

It contains (partially done for now) the fairly complete eSCL and WSD
WS-Scan implementations, DNS-SD and WS-Discovery published and querier,
smart device discovery engine, able to bind various instances of the
same device together, representing them as a sub-units of a single
physical device (IPP/IPPS, AppSocket AKA JetDirect and LPD for printing,
eSCL and WSD for scanning), based on the similar work previously done for
scanners in the [sane-airscan](https://github.com/alexpevzner/sane-airscan)
project and much more.

The ultimate goal of this project, among writing the MFP simulator, is
to implement a comprehensive Go toolkit for printing and scanning on
Linux and UNIX-like systems.

## Building

Building is quite straightforward: just type `make` at the root
directory.

It requires the following development packages to be installed:

  * Golang compiler
  * gcc compiler
  * avahi-libs for the DNS-SD

It also requires some external Go libraries, but if your build computer
is connected to the network, they will be downloaded and cached
automatically at the first build attrmpt:

  * github.com/OpenPrinting/go-avahi - Go binding for Avahi
  * github.com/OpenPrinting/goipp - Low-level IPP implementation
  * github.com/google/go-cmp - used only for testing
  * github.com/peterh/liner - line editor for the interactive shell

<!-- vim:ts=8:sw=4:et:textwidth=72
-->
