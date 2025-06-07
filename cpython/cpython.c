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
static __typeof__(Py_CompileString)             *Py_CompileString_p;
static __typeof__(Py_DecRef)                    *Py_DecRef_p;
static __typeof__(PyEval_EvalCode)              *PyEval_EvalCode_p;
static __typeof__(PyEval_RestoreThread)         *PyEval_RestoreThread_p;
static __typeof__(PyEval_SaveThread)            *PyEval_SaveThread_p;
static __typeof__(PyImport_AddModule)           *PyImport_AddModule_p;
static __typeof__(Py_InitializeEx)              *Py_InitializeEx_p;
static __typeof__(PyInterpreterState_Clear)     *PyInterpreterState_Clear_p;
static __typeof__(PyInterpreterState_Delete)    *PyInterpreterState_Delete_p;
static __typeof__(PyModule_GetDict)             *PyModule_GetDict_p;
static __typeof__(Py_NewInterpreter)            *Py_NewInterpreter_p;
static __typeof__(PyThreadState_Clear)          *PyThreadState_Clear_p;
static __typeof__(PyThreadState_Delete)         *PyThreadState_Delete_p;
static __typeof__(PyThreadState_GetInterpreter) *PyThreadState_GetInterpreter_p;
static __typeof__(PyThreadState_Get)            *PyThreadState_Get_p;
static __typeof__(PyThreadState_New)            *PyThreadState_New_p;
static __typeof__(PyThreadState_Swap)           *PyThreadState_Swap_p;

// Python build-in (primitive) types:
PyTypeObject *PyBool_Type_p;
PyTypeObject *PyByteArray_Type_p;
PyTypeObject *PyBytes_Type_p;
PyTypeObject *PyCFunction_Type_p;
PyTypeObject *PyComplex_Type_p;
PyTypeObject *PyDict_Type_p;
PyTypeObject *PyDictKeys_Type_p;
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
__typeof__(PyUnicode_AsUCS4)    *PyUnicode_AsUCS4_p;

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

// py_load_all loads all Python symbols.
static void py_load_all (void) {
    Py_CompileString_p = py_load("Py_CompileString");
    Py_DecRef_p = py_load("Py_DecRef");
    PyEval_EvalCode_p = py_load("PyEval_EvalCode");
    PyEval_RestoreThread_p = py_load("PyEval_RestoreThread");
    PyEval_SaveThread_p = py_load("PyEval_SaveThread");
    PyImport_AddModule_p = py_load("PyImport_AddModule");
    Py_InitializeEx_p = py_load("Py_InitializeEx");
    PyInterpreterState_Clear_p = py_load("PyInterpreterState_Clear");
    PyInterpreterState_Delete_p = py_load("PyInterpreterState_Delete");
    PyModule_GetDict_p = py_load("PyModule_GetDict");
    Py_NewInterpreter_p = py_load("Py_NewInterpreter");
    PyThreadState_Clear_p = py_load("PyThreadState_Clear");
    PyThreadState_Delete_p = py_load("PyThreadState_Delete");
    PyThreadState_GetInterpreter_p = py_load("PyThreadState_GetInterpreter");
    PyThreadState_Get_p = py_load("PyThreadState_Get");
    PyThreadState_New_p = py_load("PyThreadState_New");
    PyThreadState_Swap_p = py_load("PyThreadState_Swap");

    PyBool_Type_p = py_load("PyBool_Type");
    PyByteArray_Type_p = py_load("PyByteArray_Type");
    PyBytes_Type_p = py_load("PyBytes_Type");
    PyCFunction_Type_p = py_load("PyCFunction_Type");
    PyComplex_Type_p = py_load("PyComplex_Type");
    PyDict_Type_p = py_load("PyDict_Type");
    PyDictKeys_Type_p = py_load("PyDictKeys_Type");
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

    Py_IsNone_p = py_load("Py_IsNone");
    Py_IsTrue_p = py_load("Py_IsTrue");
    Py_IsFalse_p = py_load("Py_IsFalse");
    PyUnicode_GetLength_p = py_load("PyUnicode_GetLength");
    PyUnicode_AsUCS4_p = py_load("PyUnicode_AsUCS4");
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
static PyThreadState *py_enter (PyInterpreterState *interp) {
    PyThreadState *prev, *t = PyThreadState_New_p(interp);
    prev = PyThreadState_Swap_p(t);
    return prev;
}

// py_leave detaches the calling thread from the Python interpreter.
//
// Its parameter must be the value, previously returned by the
// corresponding py_enter call.
static void py_leave (PyThreadState *prev) {
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

// py_interp_eval evaluates string as a Python statement.
// It returns Python value of the executed statement on
// success, NULL in a case of any error.
PyObject *py_interp_eval (PyInterpreterState *interp, const char *s) {
    PyThreadState *prev = py_enter(interp);
    PyObject      *res = NULL;

    // Obtain the __main__ module reference and its namespace
    PyObject *main_module = PyImport_AddModule_p("__main__");
    if (main_module == NULL) {
        goto DONE;
    }

    PyObject *dict = PyModule_GetDict_p(main_module);

    // Compile the statement
    PyObject *code = Py_CompileString_p(s, "__main__", Py_eval_input);
    if (code == NULL) {
        goto DONE;
    }

    // Execute the statement
    res = PyEval_EvalCode_p(code, dict, dict);

    // Release allocated objects
    Py_DecRef_p(code);

    // Cleanup and exit.
DONE:
    py_leave(prev);
    return res;
}

// vim:ts=8:sw=4:et
