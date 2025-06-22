// MFP - Miulti-Function Printers and scanners toolkit
// Print and scam servers with added scriptability.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device model

package scriptable

import (
	"fmt"
	"io"
	"reflect"

	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/proto/ipp"
	"github.com/OpenPrinting/go-mfp/util/uuid"
)

// Model defines the whole characteristics of the MFP device being
// modeled, including the IPP printer attributes, eSCL and WSD
// scanner capabilities, scripting hooks, used to modify device
// behavior and the Python interpreter instance, used to execute
// these hooks.
type Model struct {
	py              *cpython.Python
	ippPrinterAttrs *ipp.PrinterAttributes
	esclScanCaps    *escl.ScannerCapabilities

	pyUUID       *cpython.Object
	pyRange      *cpython.Object
	pyResolution *cpython.Object
}

// NewModel creates a new Model with empty printer/scanner parameters.
func NewModel() (*Model, error) {
	// Create Python interpreter
	py, err := cpython.NewPython()
	if err != nil {
		return nil, err
	}

	// Load startup script
	err = py.Exec(embedPyInit, "init.py")
	if err != nil {
		return nil, err
	}

	// Load some commonly used Python objects
	model := &Model{py: py}
	if err == nil {
		model.pyUUID, err = py.GetGlobal("UUID")
	}
	if err == nil {
		model.pyRange, err = py.GetGlobal("Range")
	}
	if err == nil {
		model.pyResolution, err = py.GetGlobal("Resolution")
	}

	assert.Must(model.pyUUID != nil)
	assert.Must(model.pyUUID.IsCallable())
	assert.Must(model.pyRange != nil)
	assert.Must(model.pyRange.IsCallable())
	assert.Must(model.pyResolution != nil)
	assert.Must(model.pyResolution.IsCallable())

	if err != nil {
		return nil, err
	}

	return model, nil
}

// SetIPPPrinterAttrs sets the [escl.ScannerCapabilities].
func (model *Model) SetIPPPrinterAttrs(attrs *ipp.PrinterAttributes) {
	model.ippPrinterAttrs = attrs
}

// GetIPPPrinterAttrs returns the [escl.ScannerCapabilities].
func (model *Model) GetIPPPrinterAttrs() *ipp.PrinterAttributes {
	return model.ippPrinterAttrs
}

// SetESCLScanCaps sets the [escl.ScannerCapabilities].
func (model *Model) SetESCLScanCaps(caps *escl.ScannerCapabilities) {
	model.esclScanCaps = caps
}

// GetESCLScanCaps returns the [escl.ScannerCapabilities].
func (model *Model) GetESCLScanCaps() *escl.ScannerCapabilities {
	return model.esclScanCaps
}

// structToPython converts the protocol object, represented as Go
// structure, into the Python object.
//
// s MUST be struct or pointer to struct.
func (model *Model) pyExportStruct(s any) (*cpython.Object, error) {
	// Create output cpython.Object (the empty dict).
	dict, err := model.py.NewObject(map[any]any(nil))
	if err != nil {
		return nil, err
	}

	// Normalize input parameter and obtain the reflect.Value for it.
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Pointer && v.Elem().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	assert.Must((v.Kind() == reflect.Struct))

	// Roll over all struct fields
	flgs := reflect.VisibleFields(v.Type())
	for _, fld := range flgs {
		// Skip non-exposed fields
		if !fld.IsExported() {
			continue
		}

		// Obtain and normalize field value
		f := v.FieldByName(fld.Name)
		switch f.Kind() {
		case reflect.Slice:
			// Skip nil slices
			if f.IsNil() {
				continue
			}
		case reflect.Pointer:
			// Skip nil pointers. Dereference others.
			if f.IsNil() {
				continue
			}
			f = f.Elem()
		}

		// Convert into the Python Object and add to the dict,
		item, err := model.pyExportValue(f)
		if err == nil {
			err = dict.Set(fld.Name, item)
		}

		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}

// pyExportArray exports array or slice as the Python object.
func (model *Model) pyExportArray(v reflect.Value) (*cpython.Object, error) {
	list := make([]*cpython.Object, v.Len())
	var err error
	for i := 0; i < v.Len() && err == nil; i++ {
		elem := v.Index(i)
		list[i], err = model.pyExportValue(elem)
	}

	if err != nil {
		return nil, err
	}

	return model.py.NewObject(list)
}

// pyExportArray exports the value as the Python object.
func (model *Model) pyExportValue(v reflect.Value) (*cpython.Object, error) {
	// Handle known types
	data := v.Interface()
	switch v := data.(type) {
	case escl.Version:
		return model.py.NewObject(v.String())

	case escl.Range:
		args := []any{v.Min, v.Max, v.Normal}
		if v.Step != nil {
			args = append(args, *v.Step)
		}
		return model.pyRange.Call(args...)
	case escl.DiscreteResolution:
		return model.pyResolution.Call(v.XResolution, v.YResolution)
	case uuid.UUID:
		return model.pyUUID.Call(v.String())

	case fmt.Stringer:
		return model.py.NewObject(v.String())
	}

	// Switch by reflect.Kind
	switch v.Kind() {
	case reflect.Struct:
		return model.pyExportStruct(data)

	case reflect.Array, reflect.Slice:
		return model.pyExportArray(v)
	}

	// Let Python handle default case
	return model.py.NewObject(data)
}

// pyFormat writes Python object into the io.Writer.
func (model *Model) pyFormat(obj *cpython.Object, w io.Writer) error {
	f := newFormatter(w)
	return f.Format(obj)
}
