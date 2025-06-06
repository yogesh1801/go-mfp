// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CPython glue -- the C side

#ifndef cpython_h
#define cpython_h

#define Py_LIMITED_API  0x030A0000
#define PY_SSIZE_T_CLEAN

#include <Python.h>
#include <stdbool.h>

// py_init initializes Python stuff.
// It returns NULL on success or an error message in a case of errors.
// This function needs to be called only once.
const char *py_init (void);

// py_new_interp returns a new Python interpreter.
PyInterpreterState *py_new_interp (void);

// py_interp_close closes the Python interpreter.
void py_interp_close (PyInterpreterState *interp);

// py_interp_eval evaluates string as a Python statement.
// It returns true on success, false in a case of any error.
bool py_interp_eval (PyInterpreterState *interp, const char *s);

#endif

// vim:ts=8:sw=4:et
