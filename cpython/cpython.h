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
// It returns Python value of the executed statement on
// success, NULL in a case of any error.
PyObject *py_interp_eval (PyInterpreterState *interp, const char *s);

// py_obj_is_none reports if PyObject is None.
static inline int py_obj_is_none (PyObject *x) {
    extern int (*Py_IsNone_p)(PyObject *);
    return Py_IsNone_p(x);
}

// py_obj_is_true reports if PyObject is True.
static inline int py_obj_is_true (PyObject *x) {
    extern int (*Py_IsTrue_p)(PyObject *);
    return Py_IsTrue_p(x);
}

// py_obj_is_true reports if PyObject is False.
static inline int py_obj_is_false (PyObject *x) {
    extern int (*Py_IsFalse_p)(PyObject *);
    return Py_IsFalse_p(x);
}

// py_str_len returns length of Unicode string, in code points.
// If PyObject is not Unicode, it returns -1.
static inline ssize_t py_str_len (PyObject *str) {
    extern Py_ssize_t (*PyUnicode_GetLength_p)(PyObject *);
    return (ssize_t) PyUnicode_GetLength_p(str);
}

// py_str_get copies Unicode string data as a sequence of the Py_UCS4
// characters.
//
// On success it returns buf, otherwise returns NULL.
// The function may fail if PyObject is not Unicode of if buffer
// is too short.
//
// The trailing '\0' is not copied.
//
// Use py_str_len to obtain the correct string length.
Py_UCS4 *py_str_get (PyInterpreterState *interp, PyObject *str,
                     Py_UCS4 *buf, size_t len);

// Python build-in (primitive) types:
extern PyTypeObject *PyBool_Type_p;
extern PyTypeObject *PyByteArray_Type_p;
extern PyTypeObject *PyBytes_Type_p;
extern PyTypeObject *PyCFunction_Type_p;
extern PyTypeObject *PyComplex_Type_p;
extern PyTypeObject *PyDict_Type_p;
extern PyTypeObject *PyDictKeys_Type_p;
extern PyTypeObject *PyFloat_Type_p;
extern PyTypeObject *PyFrozenSet_Type_p;
extern PyTypeObject *PyList_Type_p;
extern PyTypeObject *PyLong_Type_p;
extern PyTypeObject *PyMemoryView_Type_p;
extern PyTypeObject *PyModule_Type_p;
extern PyTypeObject *PySet_Type_p;
extern PyTypeObject *PySlice_Type_p;
extern PyTypeObject *PyTuple_Type_p;
extern PyTypeObject *PyType_Type_p;
extern PyTypeObject *PyUnicode_Type_p;

#endif

// vim:ts=8:sw=4:et
