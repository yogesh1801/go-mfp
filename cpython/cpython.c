// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CPython glue -- the C side

#include "cpython.h"

#include <dlfcn.h>
#include <stdarg.h>

// py_error is set to non-NULL in a case of any initialization
// error. py_error_buf provides static buffer for that.
//
// This all is thread-safe, because used only at the initialization
// time where only a single thread has access to this stuff.
static const char                               *py_error;
static char                                     py_error_buf[1024];

// py_main_thread keeps reference to the main Python thread state
// between calls to functions that use that.
static PyThreadState                            *py_main_thread;

// The table of the libpython3 symbols.
//
// Most of the programs that use embedded Python interpreter
// link against the libpython3.NN.so.1.0 dynamic library.
//
// As result, they are tied to the particular Python version.
//
// We, instead, link against the libpython3.so and use only the
// Stable Python API (https://docs.python.org/3/c-api/stable.html).
//
// Unfortunately, libpython3.so doesn't expose any symbols and
// only contains a reference (ELF NEEDED) to the proper
// libpython3.NN.so.1.0 library.
//
// So we can't use libpython3.NN.so.1.0 symbols directly. Loading
// the libpython3.so implicitly puts these symbols into the main
// process namespace, but every symbols needs to be manually resolved
// with the explicit dlsym(RTLD_DEFAULT, name) call.
//
// So these pointers keeps these resolved symbols...
static __typeof__(PyBool_FromLong)              *PyBool_FromLong_p;
static __typeof__(PyByteArray_AsString)         *PyByteArray_AsString_p;
static __typeof__(PyByteArray_Size)             *PyByteArray_Size_p;
static __typeof__(PyBytes_AsStringAndSize)      *PyBytes_AsStringAndSize_p;
static __typeof__(PyBytes_FromStringAndSize)    *PyBytes_FromStringAndSize_p;
static __typeof__(PyCallable_Check)             *PyCallable_Check_p;
static __typeof__(Py_CompileString)             *Py_CompileString_p;
static __typeof__(PyComplex_FromDoubles)        *PyComplex_FromDoubles_p;
static __typeof__(PyComplex_ImagAsDouble)       *PyComplex_ImagAsDouble_p;
static __typeof__(PyComplex_RealAsDouble)       *PyComplex_RealAsDouble_p;
static __typeof__(Py_DecRef)                    *Py_DecRef_p;
static __typeof__(PyDict_New)                   *PyDict_New_p;
static __typeof__(PyErr_Clear)                  *PyErr_Clear_p;
static __typeof__(PyErr_Fetch)                  *PyErr_Fetch_p;
static __typeof__(PyErr_NormalizeException)     *PyErr_NormalizeException_p;
static __typeof__(PyErr_Occurred)               *PyErr_Occurred_p;
static __typeof__(PyEval_EvalCode)              *PyEval_EvalCode_p;
static __typeof__(PyEval_RestoreThread)         *PyEval_RestoreThread_p;
static __typeof__(PyEval_SaveThread)            *PyEval_SaveThread_p;
static __typeof__(PyFloat_AsDouble)             *PyFloat_AsDouble_p;
static __typeof__(PyFloat_FromDouble)           *PyFloat_FromDouble_p;
static __typeof__(PyImport_AddModule)           *PyImport_AddModule_p;
static __typeof__(Py_InitializeEx)              *Py_InitializeEx_p;
static __typeof__(PyInterpreterState_Clear)     *PyInterpreterState_Clear_p;
static __typeof__(PyInterpreterState_Delete)    *PyInterpreterState_Delete_p;
static __typeof__(PyList_GetItem)               *PyList_GetItem_p;
static __typeof__(PyList_New)                   *PyList_New_p;
static __typeof__(PyList_SetItem)               *PyList_SetItem_p;
static __typeof__(PyLong_AsLongLong)            *PyLong_AsLongLong_p;
static __typeof__(PyLong_AsUnsignedLongLong)    *PyLong_AsUnsignedLongLong_p;
static __typeof__(PyLong_FromLongLong)          *PyLong_FromLongLong_p;
static __typeof__(PyLong_FromString)            *PyLong_FromString_p;
static __typeof__(PyLong_FromUnsignedLongLong)  *PyLong_FromUnsignedLongLong_p;
static __typeof__(PyModule_GetDict)             *PyModule_GetDict_p;
static __typeof__(Py_NewInterpreter)            *Py_NewInterpreter_p;
static __typeof__(Py_NewRef)                    *Py_NewRef_p;
static __typeof__(PyObject_Call)                *PyObject_Call_p;
static __typeof__(PyObject_DelItem)             *PyObject_DelItem_p;
static __typeof__(PyObject_GetAttrString)       *PyObject_GetAttrString_p;
static __typeof__(PyObject_GetItem)             *PyObject_GetItem_p;
static __typeof__(PyObject_HasAttrString)       *PyObject_HasAttrString_p;
static __typeof__(*PyObject_Length)             *PyObject_Length_p;
static __typeof__(PyObject_Repr)                *PyObject_Repr_p;
static __typeof__(PyObject_SetAttrString)       *PyObject_SetAttrString_p;
static __typeof__(PyObject_SetItem)             *PyObject_SetItem_p;
static __typeof__(PyObject_Str)                 *PyObject_Str_p;
static __typeof__(PyThreadState_Clear)          *PyThreadState_Clear_p;
static __typeof__(PyThreadState_Delete)         *PyThreadState_Delete_p;
static __typeof__(PyThreadState_GetInterpreter) *PyThreadState_GetInterpreter_p;
static __typeof__(PyThreadState_Get)            *PyThreadState_Get_p;
static __typeof__(PyThreadState_New)            *PyThreadState_New_p;
static __typeof__(PyThreadState_Swap)           *PyThreadState_Swap_p;
static __typeof__(PyTuple_GetItem)              *PyTuple_GetItem_p;
static __typeof__(PyTuple_New)                  *PyTuple_New_p;
static __typeof__(PyTuple_SetItem)              *PyTuple_SetItem_p;
static __typeof__(*PyType_GetFlags)             *PyType_GetFlags_p;
static __typeof__(PyType_IsSubtype)             *PyType_IsSubtype_p;
static __typeof__(PyUnicode_AsUCS4)             *PyUnicode_AsUCS4_p;
static __typeof__(PyUnicode_FromStringAndSize)  *PyUnicode_FromStringAndSize_p;

// Python exceptions (some of them):
static PyObject *PyExc_KeyError_p;
static PyObject *PyExc_OverflowError_p;

// Python build-in (primitive) types:
PyTypeObject *PyBool_Type_p;
PyTypeObject *PyByteArray_Type_p;
PyTypeObject *PyBytes_Type_p;
PyTypeObject *PyCFunction_Type_p;
PyTypeObject *PyComplex_Type_p;
PyTypeObject *PyDictKeys_Type_p;
PyTypeObject *PyDict_Type_p;
PyTypeObject *PyFloat_Type_p;
PyTypeObject *PyFrozenSet_Type_p;
PyTypeObject *PyList_Type_p;
PyTypeObject *PyLong_Type_p;
PyTypeObject *PyMemoryView_Type_p;
PyTypeObject *PyModule_Type_p;
PyTypeObject *PySet_Type_p;
PyTypeObject *PySlice_Type_p;
PyTypeObject *PyTuple_Type_p;
PyTypeObject *PyType_Type_p;
PyTypeObject *PyUnicode_Type_p;

/// Directly exposed libpython functions:
int                             (*Py_IsNone_p)(PyObject *);
int                             (*Py_IsTrue_p)(PyObject *);
int                             (*Py_IsFalse_p)(PyObject *);
Py_ssize_t                      (*PyUnicode_GetLength_p)(PyObject *);

// py_set_error formats and sets py_error.
static void py_set_error (const char *fmt, ...) {
    va_list ap;

    if (py_error == NULL) {
        va_start(ap, fmt);
        vsnprintf(py_error_buf, sizeof(py_error_buf), fmt, ap);
        va_end(ap);
        py_error = py_error_buf;
    }
}

// py_load loads Python symbol by name.
static void *py_load (const char *name) {
    void *p = NULL;

    if (py_error == NULL) {
        p = dlsym(RTLD_DEFAULT, name);
        if (p == NULL) {
            py_set_error("%s", dlerror());
        }
    }

    return p;
}

// py_load loads and dereferences pointer from the libpython3.so.
static void *py_load_ptr (const char *name) {
    void **pp = py_load(name);

    if (pp != NULL) {
        return *pp;
    }

    return NULL;
}

// py_load_all loads all Python symbols.
static void py_load_all (void) {
    PyBool_FromLong_p = py_load("PyBool_FromLong");
    PyByteArray_AsString_p = py_load("PyByteArray_AsString");
    PyByteArray_Size_p = py_load("PyByteArray_Size");
    PyBytes_AsStringAndSize_p = py_load("PyBytes_AsStringAndSize");
    PyBytes_FromStringAndSize_p = py_load("PyBytes_FromStringAndSize");
    PyCallable_Check_p = py_load("PyCallable_Check");
    Py_CompileString_p = py_load("Py_CompileString");
    PyComplex_FromDoubles_p = py_load("PyComplex_FromDoubles");
    PyComplex_ImagAsDouble_p = py_load("PyComplex_ImagAsDouble");
    PyComplex_RealAsDouble_p = py_load("PyComplex_RealAsDouble");
    Py_DecRef_p = py_load("Py_DecRef");
    PyDict_New_p = py_load("PyDict_New");
    PyErr_Clear_p = py_load("PyErr_Clear");
    PyErr_Fetch_p = py_load("PyErr_Fetch");
    PyErr_NormalizeException_p = py_load("PyErr_NormalizeException");
    PyErr_Occurred_p = py_load("PyErr_Occurred");
    PyEval_EvalCode_p = py_load("PyEval_EvalCode");
    PyEval_RestoreThread_p = py_load("PyEval_RestoreThread");
    PyEval_SaveThread_p = py_load("PyEval_SaveThread");
    PyFloat_AsDouble_p = py_load("PyFloat_AsDouble");
    PyFloat_FromDouble_p = py_load("PyFloat_FromDouble");
    PyImport_AddModule_p = py_load("PyImport_AddModule");
    Py_InitializeEx_p = py_load("Py_InitializeEx");
    PyInterpreterState_Clear_p = py_load("PyInterpreterState_Clear");
    PyInterpreterState_Delete_p = py_load("PyInterpreterState_Delete");
    PyList_GetItem_p = py_load("PyList_GetItem");
    PyList_New_p = py_load("PyList_New");
    PyList_SetItem_p = py_load("PyList_SetItem");
    PyLong_AsLongLong_p = py_load("PyLong_AsLongLong");
    PyLong_AsUnsignedLongLong_p = py_load("PyLong_AsUnsignedLongLong");
    PyLong_FromLongLong_p = py_load("PyLong_FromLongLong");
    PyLong_FromString_p = py_load("PyLong_FromString");
    PyLong_FromUnsignedLongLong_p = py_load("PyLong_FromUnsignedLongLong");
    PyModule_GetDict_p = py_load("PyModule_GetDict");
    Py_NewInterpreter_p = py_load("Py_NewInterpreter");
    Py_NewRef_p = py_load("Py_NewRef");
    PyObject_Call_p = py_load("PyObject_Call");
    PyObject_DelItem_p = py_load("PyObject_DelItem");
    PyObject_GetAttrString_p = py_load("PyObject_GetAttrString");
    PyObject_GetItem_p = py_load("PyObject_GetItem");
    PyObject_HasAttrString_p = py_load("PyObject_HasAttrString");
    PyObject_Length_p = py_load("PyObject_Length");
    PyObject_Repr_p = py_load("PyObject_Repr");
    PyObject_SetAttrString_p = py_load("PyObject_SetAttrString");
    PyObject_SetItem_p = py_load("PyObject_SetItem");
    PyObject_Str_p = py_load("PyObject_Str");
    PyThreadState_Clear_p = py_load("PyThreadState_Clear");
    PyThreadState_Delete_p = py_load("PyThreadState_Delete");
    PyThreadState_GetInterpreter_p = py_load("PyThreadState_GetInterpreter");
    PyThreadState_Get_p = py_load("PyThreadState_Get");
    PyThreadState_New_p = py_load("PyThreadState_New");
    PyThreadState_Swap_p = py_load("PyThreadState_Swap");
    PyTuple_GetItem_p = py_load("PyTuple_GetItem");
    PyTuple_New_p = py_load("PyTuple_New");
    PyTuple_SetItem_p = py_load("PyTuple_SetItem");
    PyType_GetFlags_p = py_load("PyType_GetFlags");
    PyType_IsSubtype_p = py_load("PyType_IsSubtype");
    PyUnicode_AsUCS4_p = py_load("PyUnicode_AsUCS4");
    PyUnicode_FromStringAndSize_p = py_load("PyUnicode_FromStringAndSize");

    PyExc_KeyError_p = py_load_ptr("PyExc_KeyError");
    PyExc_OverflowError_p = py_load_ptr("PyExc_OverflowError");

    PyBool_Type_p = py_load("PyBool_Type");
    PyByteArray_Type_p = py_load("PyByteArray_Type");
    PyBytes_Type_p = py_load("PyBytes_Type");
    PyCFunction_Type_p = py_load("PyCFunction_Type");
    PyComplex_Type_p = py_load("PyComplex_Type");
    PyDictKeys_Type_p = py_load("PyDictKeys_Type");
    PyDict_Type_p = py_load("PyDict_Type");
    PyFloat_Type_p = py_load("PyFloat_Type");
    PyFrozenSet_Type_p = py_load("PyFrozenSet_Type");
    PyList_Type_p = py_load("PyList_Type");
    PyLong_Type_p = py_load("PyLong_Type");
    PyMemoryView_Type_p = py_load("PyMemoryView_Type");
    PyModule_Type_p = py_load("PyModule_Type");
    PySet_Type_p = py_load("PySet_Type");
    PySlice_Type_p = py_load("PySlice_Type");
    PyTuple_Type_p = py_load("PyTuple_Type");
    PyType_Type_p = py_load("PyType_Type");
    PyUnicode_Type_p = py_load("PyUnicode_Type");

    Py_IsFalse_p = py_load("Py_IsFalse");
    Py_IsNone_p = py_load("Py_IsNone");
    Py_IsTrue_p = py_load("Py_IsTrue");
    PyUnicode_GetLength_p = py_load("PyUnicode_GetLength");
}

// py_init initializes Python stuff.
// It returns NULL on success or an error message in a case of errors.
// This function needs to be called only once.
//
// This function MUST be called by the main Python thread only.
const char *py_init (void) {
    py_load_all();

    if (py_error != NULL) {
        return py_error;
    }

    Py_InitializeEx_p(0);
    py_main_thread = PyEval_SaveThread_p();

    return py_error;
}

// py_new_interp returns a new Python interpreter.
//
// This function MUST be called by the main Python thread only.
PyInterpreterState *py_new_interp (void) {
    PyThreadState      *tstate, *prev;
    PyInterpreterState *interp;

    // This stuff is very tricky.
    //
    // We first PyEval_RestoreThread(py_main_thread), to obtain
    // the global interpreter lock.
    //
    // Then Py_NewInterpreter() creates a new PyThreadState for
    // us and attaching it to the newly created sub-interpreter.
    //
    // We don't need this thread state and don't want to leak
    // its memory.
    //
    // So we PyThreadState_Swap back to the py_main_thread,
    // destroy the newly created PyThreadState.
    //
    // Finally we need to PyEval_SaveThread() to release the
    // the global interpreter lock.
    PyEval_RestoreThread_p(py_main_thread);

    tstate = Py_NewInterpreter_p();
    interp = PyThreadState_GetInterpreter_p(tstate);
    PyThreadState_Clear_p(tstate);

    PyThreadState_Swap_p(py_main_thread);

    PyThreadState_Delete_p(tstate);

    py_main_thread = PyEval_SaveThread_p();

    return interp;
}

// py_enter temporary attaches the calling thread to the
// Python interpreter.
//
// It must be called before any operations with the interpreter
// are performed and must be paired with the py_leave.
//
// The value it returns must be passed to the corresponding
// py_leave call.
PyThreadState *py_enter (PyInterpreterState *interp) {
    PyThreadState *prev, *t = PyThreadState_New_p(interp);
    prev = PyThreadState_Swap_p(t);
    return prev;
}

// py_leave detaches the calling thread from the Python interpreter.
//
// Its parameter must be the value, previously returned by the
// corresponding py_enter call.
void py_leave (PyThreadState *prev) {
    PyThreadState *t = PyThreadState_Get_p();
    PyThreadState_Clear_p(t);
    PyThreadState_Swap_p(prev);
    PyThreadState_Delete_p(t);
}

// py_interp_close closes the Python interpreter.
void py_interp_close (PyInterpreterState *interp) {
    PyThreadState *prev = py_enter(interp);
    PyInterpreterState_Clear_p(interp);
    py_leave(prev);
    PyInterpreterState_Delete_p(interp);
}

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
                     bool expr, PyObject **res) {
    // Obtain the __main__ module reference and its namespace
    PyObject *main_module = PyImport_AddModule_p("__main__");
    if (main_module == NULL) {
        return false;
    }

    PyObject *dict = PyModule_GetDict_p(main_module);

    // Compile the statement
    int      mode = expr ? Py_eval_input : Py_file_input;
    PyObject *code = Py_CompileString_p(s, file, mode);
    if (code == NULL) {
        return false;
    }

    // Execute the statement, release code object
    PyObject *ret = PyEval_EvalCode_p(code, dict, dict);
    Py_DecRef_p(code);

    // Now interpret the result
    if (ret == NULL) {
        return false;
    }

    if (!expr) {
        Py_DecRef_p(ret);
        ret = NULL;
    }

    *res = ret;
    return true;
}

// py_obj_is_byte_array reports if PyObject is PyByteArray_Type or its subclass.
bool py_obj_is_byte_array (PyObject *x) {
    return PyType_IsSubtype_p(Py_TYPE(x), PyByteArray_Type_p) != 0;
}

// py_obj_is_bytes reports if PyObject is PyBytes_Type or its subclass.
bool py_obj_is_bytes (PyObject *x) {
    return PyType_IsSubtype_p(Py_TYPE(x), PyBytes_Type_p) != 0;
}

// py_obj_is_complex reports if PyObject is PyComplex_Type or its subclass.
bool py_obj_is_complex (PyObject *x) {
    return PyType_IsSubtype_p(Py_TYPE(x), PyComplex_Type_p) != 0;
}

// py_obj_is_float reports if PyObject is PyFloat_Type or its subclass.
bool py_obj_is_float (PyObject *x) {
    return PyType_IsSubtype_p(Py_TYPE(x), PyFloat_Type_p) != 0;
}

// py_obj_is_long reports if PyObject is PyLong_Type or its subclass.
bool py_obj_is_long (PyObject *x) {
    unsigned long flags = PyType_GetFlags_p(Py_TYPE(x));
    return (flags & Py_TPFLAGS_LONG_SUBCLASS) != 0;
}

// py_obj_is_unicode reports if PyObject is PyUnicode_Type or its subclass.
bool py_obj_is_unicode (PyObject *x) {
    unsigned long flags = PyType_GetFlags_p(Py_TYPE(x));
    return (flags & Py_TPFLAGS_UNICODE_SUBCLASS) != 0;
}

// py_obj_ref increments the PyObject's reference count.
void py_obj_ref (PyObject *x) {
        Py_NewRef_p(x);
}

// py_obj_unref decrements the PyObject's reference count.
void py_obj_unref (PyObject *x) {
    Py_DecRef_p(x);
}

// py_obj_str returns a string representation of the PyObject.
// This is the equivalent of the Python expression str(x).
//
// There is very subtle difference between py_obj_str and py_obj_repr.
// In general:
//   - Use py_obj_str if you want to print the string
//   - Use py_obj_repr if you want to process the string
PyObject *py_obj_str (PyObject *x) {
    return PyObject_Str_p(x);
}

// py_obj_repr returns a string representation of the PyObject.
// This is the equivalent of the Python expression repr(x).
//
// There is very subtle difference between py_obj_str and py_obj_repr.
// In general:
//   - Use py_obj_str if you want to print the string
//   - Use py_obj_repr if you want to process the string
PyObject *py_obj_repr (PyObject *x) {
    return PyObject_Repr_p(x);
}

// py_obj_length returns PyObject, in items. It works with any
// container objects (lists, tuples, dictionaries, ...).
// Returns -1 on a error.
ssize_t py_obj_length (PyObject *x) {
    return PyObject_Length_p(x);
}

// py_obj_hasattr reports if PyObject has the attribute with the
// specified name.
//
// It returns true on success, false on error and puts answer into
// its third parameter.
bool py_obj_hasattr(PyObject *x, const char *name, bool *answer) {
    *answer = PyObject_HasAttrString_p(x, name) != 0;
    return true;
}

// py_obj_delattr deletes the attribute with the specified name.
// It returns true on success, false on error.
bool py_obj_delattr(PyObject *x, const char *name) {
    return py_obj_setattr(x, name, NULL);
}

// py_obj_getattr retrieves the attribute with the specified name.
// The returned answer, on success, contains a string reference to PyObject.
// It returns true on success, false on error.
bool py_obj_getattr(PyObject *x, const char *name, PyObject **answer) {
    PyObject *attr = PyObject_GetAttrString_p(x, name);
    *answer = attr;
    return attr != NULL;
}

// py_obj_getattr sets the attribute with the specified name.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_obj_setattr(PyObject *x, const char *name, PyObject *value) {
    return PyObject_SetAttrString_p(x, name, value) == 0;
}

// py_obj_hasitem reports if PyObject contains the item with the
// specified key.
//
// It returns true on success, false on error and puts answer into
// its third parameter.
bool py_obj_hasitem(PyObject *x, PyObject *key, bool *answer) {
    PyObject *item;
    bool     ok = py_obj_getitem(x, key, &item);

    *answer = false;
    if (item != NULL) {
        *answer = true;
        Py_DecRef_p(item);
    }

    return ok;
}

// py_obj_delitem deletes the item with the specified key.
// It returns true on success, false on error.
bool py_obj_delitem(PyObject *x, PyObject *key) {
    return PyObject_DelItem_p(x, key) == 0;
}

// py_obj_getitem retrieves the item with the specified key.
// The returned answer, on success, contains a string reference to PyObject.
// It returns true on success, false on error.
bool py_obj_getitem(PyObject *x, PyObject *key, PyObject **answer) {
    *answer = PyObject_GetItem_p(x, key);

    if (*answer != NULL) {
        return true;
    }

    PyObject *err = PyErr_Occurred_p();
    if (err == NULL || err == PyExc_KeyError_p) {
        PyErr_Clear_p();
        return true;
    }

    return false;
}

// py_obj_setitem sets the item with the specified key.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_obj_setitem(PyObject *x, PyObject *key, PyObject *value) {
    return PyObject_SetItem_p(x, key, value) == 0;
}

// py_obj_call calls callable object (i.e., function, method, ...)
// with the specified arguments.
//
// The args parameter must be of PyTuple_Type object and it specified
// the function parameters. It must not be NULL.
//
// The kwargs must be PyDict_Type and it specifies keyword arguments.
// It can be NULL, if keyword arguments are not used.
//
// It returns strong object reference on success, NULL on an error.
PyObject *py_obj_call(PyObject *x, PyObject *args, PyObject *kwargs) {
    return PyObject_Call_p(x, args, kwargs);
}

// py_obj_callable reports if object is callable.
// This function always succeeds.
bool py_obj_callable(PyObject *x) {
    return PyCallable_Check_p(x) != 0;
}

// py_err_fetch fetches and clears last error.a
// If there is no pending error, all pointers will be set to NULL.
void py_err_fetch (PyObject **etype, PyObject **evalue, PyObject **trace) {
    PyObject *exc, *val, *tb;

    PyErr_Fetch_p(&exc, &val, &tb);
    if ((exc != NULL) || (val != NULL) || (tb != NULL)) {
        PyErr_NormalizeException_p(&exc, &val, &tb);
    }

    *etype = exc;
    *evalue = val;
    *trace = tb;
}

// py_bool_make makes a new PyBool_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_bool_make(bool val) {
    return PyBool_FromLong_p((long) val);
}

// py_bytes_get obtains content of the Python bytes object.
// It returns true on success, false on error.
bool py_bytes_get (PyObject *x, void **data, size_t *size) {
    Py_ssize_t sz = 0;
    int        rc = PyBytes_AsStringAndSize_p(x, (char**) data, &sz);
    *size = (size_t) sz;
    return rc == 0;
}

// py_bytes_make makes a new PyBytes_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_bytes_make(const void *data, size_t size) {
    return PyBytes_FromStringAndSize_p(data, size);
}

// py_bytearray_get obtains content of the Python bytearray object.
// It returns true on success, false on error.
bool py_bytearray_get (PyObject *x, void **data, size_t *size) {
    Py_ssize_t sz = PyByteArray_Size_p(x);
    char       *bytes = PyByteArray_AsString_p(x);

    if (sz >= 0 && bytes != NULL) {
        *data = (void*) bytes;
        *size = (size_t) sz;
        return true;
    }

    return false;
}

// py_complex_get obtains content of the Python complex object.
// It returns true on success, false on error.
bool py_complex_get (PyObject *x, double *real, double *imag) {
    double r = PyComplex_RealAsDouble_p(x);
    double i = PyComplex_ImagAsDouble_p(x);

    if ((r == -1.0 || i == -1.0) && PyErr_Occurred_p() != NULL) {
        return false;
    }

    *real = r;
    *imag = i;

    return true;
}

// py_complex_make makes a new PyComlex_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_complex_make(double real, double imag) {
    return PyComplex_FromDoubles_p(real, imag);
}

// py_dict_make makes a new PyDict_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_dict_make(void) {
    return PyDict_New_p();
}

// py_float_get obtains content of the Python float object.
// It returns true on success, false on error.
bool py_float_get (PyObject *x, double *val) {
    double v = PyFloat_AsDouble_p(x);

    if (v == -1.0 && PyErr_Occurred_p() != NULL) {
        return false;
    }

    *val = v;

    return true;
}

// py_float_make makes a new PyFloat_Type object.
// It returns strong object reference on success, NULL on an error.
PyObject *py_float_make(double val) {
    return PyFloat_FromDouble_p(val);
}

// py_list_make makes a new PyList_Type object of the specified size.
// The newly created list MUST be fully populated with the py_list_set
// calls before it can be safely passed to Python interpreter.
// It returns strong object reference on success, NULL on an error.
PyObject *py_list_make(size_t len) {
    return PyList_New_p(len);
}

// py_list_set retrieves value of the list item at the given position.
// It returns strong object reference on success, NULL on an error.
PyObject *py_list_get(PyObject *list, int index) {
    PyObject *item = PyList_GetItem_p(list, index);
    if (item != NULL) {
        Py_NewRef_p(item);
    }
    return item;
}

// py_list_set sets value of the list item at the given position.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_list_set(PyObject *list, int index, PyObject *val) {
    Py_NewRef_p(val);
    return PyList_SetItem_p(list, index, val) == 0;
}

// py_long_get obtains PyObject's value as int64_t.
// If value doesn't fit C long, overflow flag is set.
//
// It returns true on success, false on error.
bool py_long_get_int64 (PyObject *x, int64_t *val, bool *overflow) {
    long long tmp;
    bool      ok = true, ovf = false;

    tmp = (int64_t) PyLong_AsLongLong_p(x);
    if (tmp == -1) {
        PyObject *err = PyErr_Occurred_p();
        if (err == PyExc_OverflowError_p) {
            ovf = true;
        } else {
            ok = false;
        }
    }

    if (!(INT64_MIN <= tmp && tmp <= INT64_MAX)) {
        ovf = true;
    }

    *val = (int64_t) tmp;
    *overflow = ovf;

    return ok;
}

// py_long_get obtains PyObject's value as uint64_t.
// If value doesn't fit C long, overflow flag is set.
//
// It returns true on success, false on error.
bool py_long_get_uint64 (PyObject *x, uint64_t *val, bool *overflow) {
    unsigned long long  tmp;
    bool                ok = true, ovf = false;

    tmp = (int64_t) PyLong_AsUnsignedLongLong_p(x);
    if (tmp == (unsigned long long) -1) {
        PyObject *err = PyErr_Occurred_p();
        if (err == PyExc_OverflowError_p) {
            ovf = true;
        } else {
            ok = false;
        }
    }

    if (!(tmp <= UINT64_MAX)) {
        ovf = true;
    }

    *val = (int64_t) tmp;
    *overflow = ovf;

    return ok;
}

// py_long_from_int64 makes a new PyLong_Type object from int64_t value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_int64(int64_t val) {
    return PyLong_FromLongLong_p((long long) val);
}

// py_long_from_uint64 makes a new PyLong_Type object from uint64_t value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_uint64(uint64_t val) {
    return PyLong_FromUnsignedLongLong_p((unsigned long long) val);
}

// py_long_from_string makes a new PyLong_Type object from string value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_long_from_string(const char *val) {
    return PyLong_FromString_p(val, NULL, 0);
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
Py_UCS4 *py_str_get (PyObject *str, Py_UCS4 *buf, size_t len) {
    return PyUnicode_AsUCS4_p(str, buf, len, 0);
}

// py_str_make makes a new PyLong_Type object from string value.
// It returns strong object reference on success, NULL on an error.
PyObject *py_str_make(const char *val, size_t len) {
    return PyUnicode_FromStringAndSize_p(val, len);
}

// py_tuple_make makes a new PyTuple_Type object of the specified size.
// The newly created tuple MUST be fully populated with the py_tuple_set
// calls before it can be safely passed to Python interpreter.
// It returns strong object reference on success, NULL on an error.
PyObject *py_tuple_make(size_t len) {
    return PyTuple_New_p(len);
}

// py_tuple_set retrieves value of the tuple item at the given position.
// It returns strong object reference on success, NULL on an error.
PyObject *py_tuple_get(PyObject *tuple, int index) {
    PyObject *item = PyTuple_GetItem_p(tuple, index);
    if (item != NULL) {
        Py_NewRef_p(item);
    }
    return item;
}

// py_tuple_set sets value of the tuple item at the given position.
// Internally, it creates a new strong reference to the object.
// It returns true on success, false on error.
bool py_tuple_set(PyObject *tuple, int index, PyObject *val) {
    Py_NewRef_p(val);
    return PyTuple_SetItem_p(tuple, index, val) == 0;
}

// vim:ts=8:sw=4:et
