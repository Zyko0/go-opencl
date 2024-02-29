package opencl

import (
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
	return setKernelArg(k, index, arg.size, arg.ptr).getError()
}

func (k Kernel) Release() error {
	return releaseKernel(k).getError()
}
