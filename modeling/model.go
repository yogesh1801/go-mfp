// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device model

package modeling

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/proto/ipp"
	"github.com/OpenPrinting/go-mfp/transport"
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
// Use [Model.Close] to release resources owned by the Model.
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

// Close closes the Model and releases all resources associated
// with it.
func (model *Model) Close() {
	model.py.Close()
	model.py = nil
}

// Reset resets the Modal into its initial state.
func (model *Model) Reset() error {
	model2, err := NewModel()
	if err != nil {
		return err
	}

	model.py.Close()
	*model = *model2
	return nil
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

// Write writes model into the [io.Writer]
func (model *Model) Write(w io.Writer) error {
	f := newFormatter(w)

	if model.esclScanCaps != nil {
		obj, err := model.pyExportStruct(model.esclScanCaps)
		if err != nil {
			return err
		}

		f.Printf("# eSCL scanner parameters:\n")
		f.Printf("escl.caps = ")
		err = f.Format(obj)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read reads model from the [io.Reader]
// The filename parameter required for the diagnostics messages.
func (model *Model) Read(filename string, r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = model.py.Exec(string(data), filename)
	if err != nil {
		return err
	}

	obj, err := model.py.Eval("escl.caps")
	if err != nil {
		err = fmt.Errorf("escl.caps: %s", err)
		return err
	}

	if !obj.IsNone() {
		var caps *escl.ScannerCapabilities
		err = model.pyImportStruct(&caps, obj)
		if err != nil {
			err = fmt.Errorf("escl.caps: %s", err)
			return err
		}

		model.esclScanCaps = caps
	}

	return nil
}

// Save writes model into the disk file.
func (model *Model) Save(file string) error {
	// Open the file
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	fp, err := os.OpenFile(file, flags, 0644)
	if err != nil {
		return err
	}

	// Write model data
	err = model.Write(fp)
	err2 := fp.Close()

	if err == nil {
		err = err2
	}

	return err
}

// Load reads model from the disk file.
func (model *Model) Load(file string) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}

	defer fp.Close()

	return model.Read(file, fp)
}

// pyExportStruct converts the protocol object, represented as Go
// structure or pointer to structure, into the Python object.
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
	if v.Kind() == reflect.Pointer && v.Elem().Kind() == reflect.Struct {
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
			err = dict.Set(keywordNormalize(fld.Name), item)
		}

		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}

// pyExportSlice exports slice of values as the Python object.
func (model *Model) pyExportSlice(v reflect.Value) (*cpython.Object, error) {
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

// pyExportValue exports a value as the Python object.
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

	case reflect.Slice:
		return model.pyExportSlice(v)
	}

	// Let Python handle default case
	return model.py.NewObject(data)
}

// pyImportStruct the Python object into the Go struucture, that expected
// to be the protocol object.
//
// p MUST be pointer to struct or pointer to pointer to struct.
func (model *Model) pyImportStruct(p any, obj *cpython.Object) error {

	// Validate argument
	t := reflect.TypeOf(p)

	msg := fmt.Sprintf("%s: invalid type", t)
	assert.MustMsg(t.Kind() == reflect.Pointer, msg)
	assert.MustMsg(p != nil, "nil pointer dereference")

	t = t.Elem()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	assert.MustMsg(t.Kind() == reflect.Struct, msg)

	// Create a new instance of the target structure
	v := reflect.New(t).Elem()

	// Import, field by field
	for _, fld := range reflect.VisibleFields(t) {
		// Lookup python dictionary
		kw := keywordNormalize(fld.Name)
		item, err := obj.Get(kw)
		if err != nil {
			err = fmt.Errorf("%s: %s", fld.Name, err)
			return err
		}

		// Decode the item, if found
		if item != nil {
			fldval := v.FieldByIndex(fld.Index)
			err := model.pyImportValue(fldval, item)
			if err != nil {
				return err
			}
		}
	}

	// Save output
	out := reflect.ValueOf(p).Elem()
	if out.Type().Kind() == reflect.Pointer {
		out.Set(v.Addr())
	} else {
		out.Set(v)
	}

	return nil
}

// pyImportSlice imports slice of values from the Python object.
func (model *Model) pyImportSlice(v reflect.Value, obj *cpython.Object) error {
	// Obtain Python object items
	slice, err := obj.Slice()
	if err != nil {
		return err
	}

	// Allocate output memory
	v.Set(reflect.MakeSlice(v.Type(), len(slice), len(slice)))

	// Decode item by item
	for i, item := range slice {
		err = model.pyImportValue(v.Index(i), item)
		if err != nil {
			return err
		}
	}

	return nil
}

// pyImportValue imports a value from the Python object.
func (model *Model) pyImportValue(v reflect.Value, obj *cpython.Object) error {
	// If we are decoding pointer to value, create a new
	// value instance and shift to it.
	if v.Kind() == reflect.Pointer {
		v2 := reflect.New(v.Type().Elem())
		v.Set(v2)
		v = v2.Elem()
	}

	// Handle known types
	switch v.Interface().(type) {
	case escl.ADFOption:
		opt, err := esclDecodeADFOption(obj)
		if err == nil {
			v.Set(reflect.ValueOf(opt))
		}
		return err

	case escl.ADFState:
		st, err := esclDecodeADFState(obj)
		if err == nil {
			v.Set(reflect.ValueOf(st))
		}
		return err

	case escl.BinaryRendering:
		rnd, err := esclDecodeBinaryRendering(obj)
		if err == nil {
			v.Set(reflect.ValueOf(rnd))
		}
		return err

	case escl.CCDChannel:
		ccd, err := esclDecodeCCDChannel(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ccd))
		}
		return err

	case escl.ColorMode:
		cm, err := esclDecodeColorMode(obj)
		if err == nil {
			v.Set(reflect.ValueOf(cm))
		}
		return err

	case escl.ColorSpace:
		sps, err := esclDecodeColorSpace(obj)
		if err == nil {
			v.Set(reflect.ValueOf(sps))
		}
		return err

	case escl.ContentType:
		ct, err := esclDecodeContentType(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ct))
		}
		return err

	case escl.DiscreteResolution:
		r, err := esclDecodeDiscreteResolution(obj)
		if err == nil {
			v.Set(reflect.ValueOf(r))
		}
		return err

	case escl.FeedDirection:
		feed, err := esclDecodeFeedDirection(obj)
		if err == nil {
			v.Set(reflect.ValueOf(feed))
		}
		return err

	case escl.ImagePosition:
		pos, err := esclDecodeImagePosition(obj)
		if err == nil {
			v.Set(reflect.ValueOf(pos))
		}
		return err

	case escl.InputSource:
		src, err := esclDecodeInputSource(obj)
		if err == nil {
			v.Set(reflect.ValueOf(src))
		}
		return err

	case escl.Intent:
		intent, err := esclDecodeIntent(obj)
		if err == nil {
			v.Set(reflect.ValueOf(intent))
		}
		return err

	case escl.JobState:
		st, err := esclDecodeJobState(obj)
		if err == nil {
			v.Set(reflect.ValueOf(st))
		}
		return err

	case escl.JobStateReason:
		rsn, err := esclDecodeJobStateReason(obj)
		if err == nil {
			v.Set(reflect.ValueOf(rsn))
		}
		return err

	case escl.Units:
		un, err := esclDecodeUnits(obj)
		if err == nil {
			v.Set(reflect.ValueOf(un))
		}
		return err

	case escl.Range:
		r, err := esclDecodeRange(obj)
		if err == nil {
			v.Set(reflect.ValueOf(r))
		}
		return err

	case escl.Version:
		ver, err := esclDecodeVersion(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ver))
		}
		return err

	case uuid.UUID:
		s, err := obj.Str()
		if err != nil {
			return err
		}

		u, err := uuid.Parse(s)
		if err == nil {
			v.Set(reflect.ValueOf(u))
		}

		return err
	}

	// Switch by reflect.Kind
	switch v.Kind() {
	case reflect.Struct:
		return model.pyImportStruct(v.Addr().Interface(), obj)

	case reflect.Slice:
		return model.pyImportSlice(v, obj)

	case reflect.Int:
		i, err := obj.Int()
		if err == nil {
			v.Set(reflect.ValueOf(int(i)))
		}
		return err

	case reflect.String:
		s, err := obj.Str()
		if err == nil {
			v.Set(reflect.ValueOf(s))
		}
		return err
	}

	return nil
}

// NewESCLServer creates a virtual eSCL server on a top of
// the existent abstract.Scanner implementation.
//
// It will return nil, if model doesn't have the eSCL scanner capabilities.
func (model *Model) NewESCLServer(
	scanner abstract.Scanner) *escl.AbstractServer {

	// Obtain scanner capabilities
	caps := model.GetESCLScanCaps()
	if caps == nil {
		return nil
	}

	// Setup options
	options := escl.AbstractServerOptions{
		Version:  caps.Version,
		Scanner:  scanner,
		BasePath: "/eSCL",
		Hooks: escl.ServerHooks{
			OnScannerCapabilitiesResponse: model.esclOnScannerCapabilitiesResponse,
		},
	}

	// Create the eSCL server
	return escl.NewAbstractServer(options)
}

// esclOnScannerCapabilitiesResponse implements the
// [escl.ServerHooks.OnScannerCapabilitiesResponse] hook
// for the modeled eSCL scanner.
func (model *Model) esclOnScannerCapabilitiesResponse(
	query *transport.ServerQuery,
	caps *escl.ScannerCapabilities) *escl.ScannerCapabilities {

	caps2 := model.GetESCLScanCaps()
	if caps2 == nil {
		query.Reject(http.StatusServiceUnavailable, nil)
	}

	return caps2
}
