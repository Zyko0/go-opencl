//go:build windows

package opencl

import (
	"syscall"
	"unsafe"
)

func loadLibrary() (uintptr, error) {
	handle, err := syscall.LoadLibrary("opencl.dll")
	if err != nil {
		return 0, err
	}
	// purego unsupported functions
	dll := syscall.DLL{
		Name:   "opencl.dll",
		Handle: handle,
	}
	// Note: Functions with too many arguments requiring manual loading
	readImg := dll.MustFindProc("clEnqueueReadImage")
	mapImg := dll.MustFindProc("clEnqueueMapImage")
	mapBuffer := dll.MustFindProc("clEnqueueMapBuffer")
	enqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]clSize, row_pitch, slice_pitch clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		r1, _, _ := readImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(unsafe.Pointer(&origin[0])),
			uintptr(unsafe.Pointer(&region[0])),
			uintptr(row_pitch),
			uintptr(slice_pitch),
			uintptr(ptr),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
		)
		return clStatus(r1)
	}
	enqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]clSize, imageRowPitch, imageSlicePitch *clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		r1, _, _ := mapImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(mapFlags),
			uintptr(unsafe.Pointer(&origin[:][0])),
			uintptr(unsafe.Pointer(&region[:][0])),
			uintptr(unsafe.Pointer(imageRowPitch)),
			uintptr(unsafe.Pointer(imageSlicePitch)),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)

		return r1
	}
	enqueueMapBuffer = func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		r1, _, _ := mapBuffer.Call(
			uintptr(queue),
			uintptr(buffer),
			uintptr(block),
			uintptr(mapFlags),
			uintptr(offset),
			uintptr(size),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)

		return r1
	}

	return uintptr(handle), err
}
