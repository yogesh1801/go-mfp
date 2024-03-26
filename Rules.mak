# MFP - comprehensive toolkit for multi-function printers and scanners
#
# Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
# See LICENSE for license terms and conditions
#
# Common part for Makefiles

# ----- Environment -----

TOPDIR  := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
GOFILES := $(wildcard *.go)

# ----- Parameters -----

GO	:= go
GOTAGS	:= $(shell which gotags 2>/dev/null)

# ----- Common targets -----

.PHONY: all
.PHONY: clean
.PHONY: cover
.PHONY: tags
.PHONY: test
.PHONY: vet

# Recursive targets
all:	subdirs_all do_all
clean:	subdirs_clean do_clean
test:	subdirs_test do_test
vet:	subdirs_vet do_vet

# Non-recursive targets
cover:	do_cover

# Dependencies
do_all:	tags

# Default actions
tags:
do_all:
do_clean:
do_cover:
do_test:
do_vet:

# Conditional actions
ifneq   ($(GOFILES),)

do_all:
	$(GO) build

do_cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out

do_test:
	$(GO) test

do_vet:
	$(GO) vet

endif


ifneq	($(GOTAGS),)
tags:
	cd $(TOPDIR); gotags -R . > tags
endif

# ----- Subdirs handling

subdirs_all subdirs_test subdirs_vet subdirs_clean:
	@for i in $(SUBDIRS); do \
                $(MAKE) -C $$i $(subst subdirs_,,$@) || exit 1; \
        done

