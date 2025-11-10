// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// libppd binding (C helpers)

#ifndef libppd_h
#define libppd_h

#include <ppd/ppd.h>

// libppd_init initialized the libppd.
// It returns NULL on success or error message in a case of error.
const char *libppd_init (void);

// libppd/libcups function wrappers.
void            libppd_ippDelete(ipp_t *ipp);
ipp_state_t     libppd_ippWriteFile(int fd, ipp_t *ipp);
void            libppd_ppdClose(ppd_file_t *ppd);
ppd_file_t      *libppd_ppdOpenFd(int fd);
ipp_t           *libipp_ppdLoadAttributes(ppd_file_t *ppd);

#endif

// vim:ts=8:sw=4:et
