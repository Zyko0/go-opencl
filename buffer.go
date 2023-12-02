package opencl

import (
	"errors"
	"unsafe"
)

type MemFlag uint32

const (
	MemFlagsReadWrite    MemFlag = 1 << 0
	MemFlagsWriteOnly    MemFlag = 1 << 1
	MemFlagsReadOnly     MemFlag = 1 << 2
	MemFlagsUseHostPtr   MemFlag = 1 << 3
	MemFlagsAllocHostPtr MemFlag = 1 << 4
	MemFlagsCopyHostPtr  MemFlag = 1 << 5
	// reserved
	MemFlagsHostWriteOnly MemFlag = 1 << 7
	MemFlagsHostReadOnly  MemFlag = 1 << 8
	MemFlagsHostNoAccess  MemFlag = 1 << 9
	// next ones.. (lazy)
)

type Buffer uint

type memInfo uint32

const (
	MemInfoSize memInfo = 0x1102
)

func (b Buffer) getInfo(name memInfo) (uint, error) {
	/*size := clSize(0)
	st := getPlatformInfo(p, name, clSize(0), nil, &size)
	if st != CL_SUCCESS {
		return "", errors.New("oops at 1st get platform info")
	}*/

	info := uint(0)
	ptr := &info
	st := getMemObjectInfo(b, name, clSize(unsafe.Sizeof(ptr)), ptr, nil)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at 2nd get platform info")
	}

	return info, nil
}

func (b Buffer) Size() (uint, error) {
	return b.getInfo(MemInfoSize)
}

func (b Buffer) Release() error {
	st := releaseMemObject(b)
	if st != CL_SUCCESS {
		return errors.New("oops at release buffer")
	}

	return nil
}
