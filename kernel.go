package opencl

import (
	"errors"
	"strconv"
	"unsafe"
)

type Kernel uint

type KernelArg struct {
	ptr  unsafe.Pointer
	size clSize
}

func NewKernelArg[T any](arg *T) KernelArg {
	return KernelArg{
		ptr:  unsafe.Pointer(arg),
		size: clSize(unsafe.Sizeof(*arg)),
	}
}

func (k Kernel) SetArg(index uint, arg KernelArg) error {
	st := setKernelArg(k, index, arg.size, arg.ptr)
	if st != CL_SUCCESS {
		return errors.New("oops at set kernel arg: " + strconv.FormatInt(int64(st), 10))
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
