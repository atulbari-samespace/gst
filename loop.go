package gst

/*
#cgo CFLAGS: -I/usr/include/gstreamer-1.0/ -I/usr/include/glib-2.0/ -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -l:libgthread-2.0.a -l:libgio-2.0.a   -l:libgobject-2.0.a -l:libffi.a  -l:libgstreamer-full-1.0.a -l:libgmodule-2.0.a -lm -lunwind -ldl
#include "gst.h"
*/
import "C"

import (
	"runtime"
)

type GMainLoop struct {
	C *C.GMainLoop
}

func MainLoopNew() (loop *GMainLoop) {
	CLoop := C.g_main_loop_new(nil, C.gboolean(0))
	loop = &GMainLoop{C: CLoop}
	runtime.SetFinalizer(loop, func(loop *GMainLoop) {
		C.g_main_loop_unref(loop.C)
	})

	return
}

func (l *GMainLoop) Run() {
	C.g_main_loop_run(l.C)
}

func (l *GMainLoop) Quit() {
	C.g_main_loop_quit(l.C)
}

func (l *GMainLoop) IsRuning() bool {
	Cbool := C.g_main_loop_is_running(l.C)
	if Cbool == 1 {
		return true
	}
	return false
}
