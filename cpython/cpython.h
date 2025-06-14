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

// py_enter temporary attaches the calling thread to the
// Python interpreter.
//
// It must be called before any operations with the interpreter
// are performed and must be paired with the py_leave.
//
// The value it returns must be passed to the corresponding
// py_leave call.
PyThreadState *py_enter (PyInterpreterState *interp);

// py_leave detaches the calling thread from the Python interpreter.
//
// Its parameter must be the value, previously returned by the
// corresponding py_enter call.
void py_leave (PyThreadState *prev);

// py_interp_eval evaluates string as a Python statement or expression.
// It returns Python value of the executed statement on
// success, NULL in a case of any error.
//
// The name parameter is used for diagnostics messages and
// indicated the input file name.
//
// If expr is true, this function evaluates Python expression and
// saves its result into res. Otherwise, it evaluates a multi-line
// Python script and don't return any PyObject (sets *res to NULL)
bool py_interp_eval (const char *s, const char *file,
                     bool expr, PyObject **res);

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

// py_obj_ref increments the PyObject's reference count.
void py_obj_ref (PyObject *x);

// py_obj_unref decrements the PyObject's reference count.
void py_obj_unref (PyObject *x);

// py_obj_str returns a string representation of the PyObject.
// This is the equivalent of the Python expression str(x).
//
// There is very subtle difference between py_obj_str and py_obj_repr.
// In general:
//   - Use py_obj_str if you want to print the string
//   - Use py_obj_repr if you want to process the string
PyObject *py_obj_str (PyObject *x);

// py_obj_repr returns a string representation of the PyObject.
// This is the equivalent of the Python expression repr(x).
//
// There is very subtle difference between py_obj_str and py_obj_repr.
// In general:
//   - Use py_obj_str if you want to print the string
//   - Use py_obj_repr if you want to process the string
PyObject *py_obj_repr (PyObject *x);

// py_obj_hasattr reports if PyObject has the attribute with the
// specified name.
//
// It returns true on success, false on error and puts answer into
// its third parameter.
bool py_obj_hasattr(PyObject *x, const char *name, bool *answer);

// py_obj_hasattr deletes the attribute with the specified name.
// It returns true on success, false on error.
bool py_obj_delattr(PyObject *x, const char *name);

// py_obj_hasattr retrieves the attribute with the specified name.
// The returned answer, on success, contains a string reference to PyObject.
// It returns true on success, false on error.
bool py_obj_getattr(PyObject *x, const char *name, PyObject **answer);

// py_obj_hasattr sets the attribute with the specified name.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_obj_setattr(PyObject *x, const char *name, PyObject *value);

// py_err_fetch fetches and clears last error.
// If there is no pending error, all pointers will be set to NULL.
void py_err_fetch (PyObject **etype, PyObject **evalue, PyObject **trace);

// py_bool_make makes a new PyBool_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_bool_make(bool val);

// py_bytes_get obtains content of the Python bytes object.
// It returns true on success, false on error.
bool py_bytes_get (PyObject *x, void **data, size_t *size);

// py_bytearray_get obtains content of the Python bytearray object.
// It returns true on success, false on error.
bool py_bytearray_get (PyObject *x, void **data, size_t *size);

// py_complex_get obtains content of the Python complex object.
// It returns true on success, false on error.
bool py_complex_get (PyObject *x, double *real, double *imag);

// py_complex_make makes a new PyComlex_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_complex_make(double real, double imag);

// py_float_get obtains content of the Python float object.
// It returns true on success, false on error.
bool py_float_get (PyObject *x, double *val);

// py_float_make makes a new PyFloat_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_float_make(double val);

// py_list_make makes a new PyList_Type object of the specified size.
// The newly created list MUST be fully populated with the py_list_set
// calls before it can be safely passed to Python interpreter.
// It returns strong object reference on success, NULL on an error.
PyObject *py_list_make(size_t len);

// py_list_set retrieves value of the list item at the given position.
// The returned answer, on success, contains a string reference to PyObject.
// It returns true on success, false on error.
bool py_list_get(PyObject *list, int index, PyObject **answer);

// py_list_set sets value of the list item at the given position.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_list_set(PyObject *list, int index, PyObject *val);

// py_long_get obtains PyObject's value as C long.
// If value doesn't fit C long, overflow flag is set.
//
// It returns true on success, false on error.
bool py_long_get (PyObject *x, long *val, bool *overflow);

// py_long_from_int64 makes a new PyLong_Type object from int64_t value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_int64(int64_t val);
//
// py_long_from_uint64 makes a new PyLong_Type object from uint64_t value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_uint64(uint64_t val);

// py_long_from_string makes a new PyLong_Type object from string value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_string(const char *val);

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
Py_UCS4 *py_str_get (PyObject *str, Py_UCS4 *buf, size_t len);

// py_str_make makes a new PyLong_Type object from string value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_str_make(const char *val, size_t len);

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
