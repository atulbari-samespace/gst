package gst

/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type Caps struct {
	caps *C.GstCaps
}

func CapsFromString(caps string) (gstCaps *Caps) {
	c := (*C.gchar)(unsafe.Pointer(C.CString(caps)))
	defer C.g_free(C.gpointer(unsafe.Pointer(c)))
	CCaps := C.gst_caps_from_string(c)
	gstCaps = &Caps{
		caps: CCaps,
	}

	runtime.SetFinalizer(gstCaps, func(gstCaps *Caps) {
		C.gst_caps_unref(gstCaps.caps)
	})

	return
}

func (c *Caps) ToString() (str string) {
	CStr := C.gst_caps_to_string(c.caps)
	defer C.g_free(C.gpointer(unsafe.Pointer(CStr)))
	str = C.GoString((*C.char)(unsafe.Pointer(CStr)))
	return
}

func (c *Caps) String() (str string) {
	CStr := C.gst_caps_to_string(c.caps)
	defer C.g_free(C.gpointer(unsafe.Pointer(CStr)))
	str = C.GoString((*C.char)(unsafe.Pointer(CStr)))
	return
}

func (c *Caps) GetStructure(index int) (structure *Structure) {
	Cstructure := C.gst_caps_get_structure(c.caps, C.uint(index))
	structure = &Structure{
		C: Cstructure,
	}

	return
}
