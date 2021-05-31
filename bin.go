package gst


/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

type Bin struct {
	Element
}

func ParseBinFromDescription(binStr string, ghostPads bool) (bin *Bin, err error) {
	var gError *C.GError

	pDesc := (*C.gchar)(unsafe.Pointer(C.CString(binStr)))
	defer C.g_free(C.gpointer(unsafe.Pointer(pDesc)))

	var ghost int
	if ghostPads {
		ghost = 1
	} else {
		ghost = 0
	}

	gstElt := C.gst_parse_bin_from_description(pDesc, C.int(ghost), &gError)

	if gError != nil {
		err = errors.New("create bin error")
		return
	}

	bin = &Bin{}
	bin.GstElement = gstElt

	runtime.SetFinalizer(bin, func(bin *Bin) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bin.GstElement)))
	})

	return
}

func BinNew(name string) (bin *Bin) {

	pName := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(pName)))

	Celement := C.gst_bin_new(pName)
	bin = &Bin{}

	bin.GstElement = Celement

	runtime.SetFinalizer(bin, func(bin *Bin) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bin.GstElement)))
	})

	return
}

func (b *Bin) Add(child *Element) {

	C.X_gst_bin_add(b.GstElement, child.GstElement)
	return
}

func (b *Bin) Remove(child *Element) {

	C.X_gst_bin_remove(b.GstElement, child.GstElement)
	return
}

func (b *Bin) AddMany(elements ...*Element) {
	for _, e := range elements {
		if e != nil {
			C.X_gst_bin_add(b.GstElement, e.GstElement)
		}
	}

	return
}

func (b *Bin) GetByName(name string) (element *Element) {

	n := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	CElement := C.X_gst_bin_get_by_name(b.GstElement, n)

	if CElement == nil {
		return
	}

	element = &Element{
		GstElement: CElement,
	}

	return
}
