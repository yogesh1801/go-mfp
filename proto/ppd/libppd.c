// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// libppd binding (C helpers)

#define _GNU_SOURCE
#include <dlfcn.h>
#include <stdlib.h>

#include "libppd.h"

// libppd_error set to non-NULL string in a case of any initialization
// error.
static const char       *libppd_error;

static void             *libppd_handle;

// Table of dynamically loaded libppd symbols.
//
// Note: Some symbols are defined in both libppd and libcups,
// but because their internal structures differ, these definitions
// are not compatible. If a program uses both libppd and libcups,
// objects might be constructed in libppd while destructed by the
// libcups-provided destructor, causing the application to crash.
// Also note that libppd itself dynamically links to libcups,
// so the exact behavior depends on the linking order.
//
// More details can be found here:
//   https://github.com/OpenPrinting/libppd/issues/52
//
// It makes things very fragile and unreliable.
//
// We address this problem by loading libppd dynamically into
// a private namespace using dlmopen(LM_ID_NEWLM, ...).
//
// Unfortunately, this approach is not portable. This code was
// tested on Linux and will probably work on FreeBSD, but it
// won't compile on NetBSD and OpenBSD. A better long-term
// solution is required.
static __typeof__(ippDelete)            *ippDelete_p;
static __typeof__(ippWriteFile)         *ippWriteFile_p;
static __typeof__(ppdClose)             *ppdClose_p;
static __typeof__(ppdCreatePPDFromIPP)  *ppdCreatePPDFromIPP_p;
static __typeof__(ppdLoadAttributes)    *ppdLoadAttributes_p;
static __typeof__(ppdOpenFd)            *ppdOpenFd_p;

// libppd_set_error formats the error message text and sets libppd_error.
static void libppd_set_error (const char *fmt, ...) {
    va_list     ap;
    static char error_buf[1024];

    if (libppd_error == NULL) {
        va_start(ap, fmt);
        vsnprintf(error_buf, sizeof(error_buf), fmt, ap);
        va_end(ap);
        libppd_error = error_buf;
    }

}

// libppd_load dynamically loads a symbol from libppd.
static void *libppd_load (const char *name) {
    void *p = NULL;

    if (libppd_error == NULL) {
        p = dlsym(libppd_handle, name);
        if (p == NULL) {
            libppd_set_error("%s", dlerror());
        }
    }

    return p;
}

// libppd_init initialized the libppd.
// It returns NULL on success or error message in a case of error.
const char *libppd_init (void) {
    libppd_handle = dlmopen(LM_ID_NEWLM, "libppd.so",
        RTLD_NOW | RTLD_LOCAL);

    if (libppd_handle == NULL) {
        libppd_set_error("%s", dlerror());
        return libppd_error;
    }

    ippDelete_p = libppd_load("ippDelete");
    ippWriteFile_p = libppd_load("ippWriteFile");
    ppdClose_p = libppd_load("ppdClose");
    ppdCreatePPDFromIPP_p = libppd_load("ppdCreatePPDFromIPP");
    ppdLoadAttributes_p = libppd_load("ppdLoadAttributes");
    ppdOpenFd_p = libppd_load("ppdOpenFd");

    return libppd_error;
}

// libppd_ippWriteFile wraps the ippWriteFile function.
ipp_state_t libppd_ippWriteFile(int fd, ipp_t *ipp) {
    return ippWriteFile_p(fd, ipp);
}

// libppd_ippDelete wraps the ippDelete function.
void libppd_ippDelete(ipp_t *ipp) {
    ippDelete_p(ipp);
}

// libppd_ppdClose wraps the ppdClose function.
void libppd_ppdClose(ppd_file_t *ppd) {
    ppdClose_p(ppd);
}

// libppd_ppdOpenFd wraps the ppdOpenFd function.
ppd_file_t *libppd_ppdOpenFd(int fd) {
    return ppdOpenFd_p(fd);
}

// libipp_ppdLoadAttributes wraps the ppdLoadAttributes function,
ipp_t *libipp_ppdLoadAttributes(ppd_file_t *ppd) {
    return ppdLoadAttributes_p(ppd);
}

// vim:ts=8:sw=4:et
