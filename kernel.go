package opencl

import (
	"errors"
	"unsafe"
)

type Kernel uint

func (k Kernel) SetArg(index uint, arg *Buffer) error {
	// TODO: Support different types than *Buffer
	ptr := unsafe.Pointer(arg)
	st := setKernelArg(k, index, clSize(unsafe.Sizeof(ptr)), ptr)
	if st != CL_SUCCESS {
		return errors.New("oops at set kernel arg")
	}

	return nil
}

func (k Kernel) Release() error {
	st := releaseKernel(k)
	if st != CL_SUCCESS {
		return errors.New("oops at release kernel")
	}

	return nil
}
