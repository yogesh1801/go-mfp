// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Object identifiers

package cpython

import (
	"sync/atomic"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// objid uniquely identified each *C.PyObject owned by *C.PyInterpreterState
type objid uint64

// inc atomically increments the objid and returns the new value.
func (oid *objid) inc() objid {
	return objid(atomic.AddUint64((*uint64)(oid), 1))
}

// objmap maintains the mapping between *C.PyObject-s and assigned objid-s.
type objmap struct {
	next   objid
	mapped map[objid]pyObject
}

// newObjmap creates a new objmap
func newObjmap() *objmap {
	return &objmap{
		mapped: make(map[objid]pyObject, 512),
	}
}

// put adds *C.PyObject to the map and returns assigned objid.
func (omap *objmap) put(gate pyGate, obj pyObject) objid {
	oid := omap.next.inc()
	assert.Must(omap.mapped[oid] == nil)
	omap.mapped[oid] = obj
	return oid
}

// get returns *C.PyObject by objid.
func (omap *objmap) get(gate pyGate, oid objid) pyObject {
	return omap.mapped[oid]
}

// del removes the *C.PyObject from the map and deletes its strong referenced.
func (omap *objmap) del(gate pyGate, oid objid) {
	obj := omap.mapped[oid]
	delete(omap.mapped, oid)
	gate.unref(obj)
}

// purge removes all objects from the map.
func (omap *objmap) purge(gate pyGate) {
	objects := make([]pyObject, 0, len(omap.mapped))

	for oid, obj := range omap.mapped {
		objects = append(objects, obj)
		delete(omap.mapped, oid)
	}

	for _, obj := range objects {
		gate.unref(obj)
	}
}

// count returns count of currently mapped objid-s
func (omap *objmap) count(gate pyGate) int {
	return len(omap.mapped)
}
