// +build !arm

package gst

/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"
import "time"

func (e *Element) Seek(duration time.Duration) bool {
	result := C.gst_element_seek_simple(e.GstElement, C.GST_FORMAT_TIME, C.GST_SEEK_FLAG_FLUSH, C.long(duration.Nanoseconds()))
	return result == C.TRUE
}
