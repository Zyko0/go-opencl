package opencl

import (
	"errors"
	"strconv"
	"unsafe"
)

type BufferData struct {
	TypeSize uintptr
	DataSize uintptr
	Pointer  unsafe.Pointer
}

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
	info := uint(0)
	st := getMemObjectInfo(b, name, clSize(unsafe.Sizeof(info)), unsafe.Pointer(&info), nil)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at get buffer info: " + strconv.FormatInt(int64(st), 10))
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

// GL

func (b Buffer) GetGLObjectInfo() (CLGLObjectType, error) {
	var objectType CLGLObjectType

	st := getGLObjectInfo(b, &objectType, nil)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at get gl object info")
	}

	return objectType, nil
}

func (b Buffer) GetGLTextureInfo(info CLGLTextureInfo) (uint32, error) {
	var results = []uint32{0}

	st := getGLTextureInfo(
		b, info, clSize(unsafe.Sizeof(&results[0])), unsafe.Pointer(&results[0]), nil,
	)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at get gl texture info")
	}

	return results[0], nil
}
