package gst

/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	C.X_gst_shim_init()
}

type Sample struct {
	Data     []byte
	Duration uint64
	Pts      uint64
	Dts      uint64
	Offset   uint64
}

func CheckPlugins(plugins []string) error {

	var plugin *C.GstPlugin
	var registry *C.GstRegistry

	registry = C.gst_registry_get()

	for _, pluginstr := range plugins {
		plugincstr := C.CString(pluginstr)
		plugin = C.gst_registry_find_plugin(registry, plugincstr)
		C.free(unsafe.Pointer(plugincstr))
		if plugin == nil {
			return fmt.Errorf("Required gstreamer plugin %s not found", pluginstr)
		}
	}

	return nil
}
