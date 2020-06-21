package service

/*
#cgo CFLAGS: -I../libobject
#cgo LDFLAGS: -L../libobject -lobject
#cgo LDFLAGS: -Wl,-rpath="./libobject"
#include <stdlib.h>
#include "object_manager.h"
*/
import "C"

import (
	"counter/common"
	"fmt"
	"sync"
	"unsafe"
)

var gAlgoHandler unsafe.Pointer
var gMutex sync.Mutex

func algoInit(modelPath, tag string) error {
	cpath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cpath))
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))

	ret := C.load_object_manager(cpath, ctag, &gAlgoHandler)
	if ret != 0 {
		return fmt.Errorf("algo init error: %d", ret)
	}
	return nil
}

func algoProcess(img string, rect AlgoRect) (error, []AlgoRect) {
	gMutex.Lock()
	defer gMutex.Unlock()
	// fmt.Printf("cimg:%s\n", img)
	cimg := C.CString(img)
	defer C.free(unsafe.Pointer(cimg))

	var cRect C.HRect
	var cRects *C.HRect
	var cInt C.int

	cRect.x = C.int(rect.X)
	cRect.y = C.int(rect.Y)
	cRect.width = C.int(rect.Width)
	cRect.height = C.int(rect.Height)

	ret := C.detect_objects(gAlgoHandler, cimg, cRect, &cRects, &cInt)
	unsafePtr := unsafe.Pointer(cRects)
	goInt := int(cInt)
	defer C.free(unsafePtr)
	if ret != 0 {
		return fmt.Errorf("algo process error: %d", ret), nil
	}
	if cRects == nil {
		return fmt.Errorf("algo process cRects nullptr"), nil
	}
	if goInt == 0 {
		return fmt.Errorf("algo porocess not detectd"), nil
	}
	common.Log.Debugf("ret: %d, goInt: %d, rects: %v", ret, goInt, cRects)
	arrayPtr := (*[1 << 30]C.HRect)(unsafePtr)
	goSlice := arrayPtr[0:goInt:goInt]
	algoRet := make([]AlgoRect, 0)
	for _, v := range goSlice {
		algoRet = append(algoRet, AlgoRect{
			X:      int(v.x),
			Y:      int(v.y),
			Width:  int(v.width),
			Height: int(v.height),
		})
	}
	return nil, algoRet
}

func algoDestroy() {
	C.release_object_manager(&gAlgoHandler)
}
