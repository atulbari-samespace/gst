package gst

/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"

type Event struct {
	GstEvent *C.GstEvent
}

func NewEosEvent() (event *Event) {
	CGstEvent := C.gst_event_new_eos()

	event = &Event{
		GstEvent: CGstEvent,
	}

	return
}
