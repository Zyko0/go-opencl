package opencl

import (
	"errors"
	"strconv"
	"unsafe"
)

type ContextProperty uint

const (
	ContextPropertyPlatform        ContextProperty = 0x1084
	ContextPropertyInteropUserSync ContextProperty = 0x1085 // >= 1.2
	// GL
	ContextPropertyGLContextKHR ContextProperty = 0x2008
	ContextPropertyWGL_HDC_KHR  ContextProperty = 0x200B
)

type Context uint

// TODO: make properties into a struct instead of weird map<uint32>

func (d Device) CreateContext(properties map[ContextProperty]ContextProperty) (Context, error) {
	var st clStatus

	flatten := make([]ContextProperty, 0, len(properties)*2)
	for k, v := range properties {
		flatten = append(flatten, k, ContextProperty(v))
	}
	flatten = append(flatten, 0) // End of list
	ctx := createContext(unsafe.Pointer(&flatten[0]), 1, []Device{d}, nil, nil, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create context: " + strconv.FormatInt(int64(st), 10))
	}

	return ctx, nil
}

func (c Context) CreateCommandQueue(device Device) (CommandQueue, error) {
	var st clStatus

	queue := createCommandQueue(c, device, 0, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create command queue")
	}

	return queue, nil
}

type CommandQueueProperty uint32

const (
	CommandQueueOutOfOrderExecModeEnable CommandQueueProperty = (1 << 0)
	CommandQueueProfilingEnable          CommandQueueProperty = (1 << 1)
	CommandQueueOnDevice                 CommandQueueProperty = (1 << 2)
	CommandQueueOnDeviceDefault          CommandQueueProperty = (1 << 3)
)

func (c Context) CreateCommandQueueWithProperties(device Device, properties []CommandQueueProperty) (CommandQueue, error) {
	var st clStatus

	property := CommandQueueProperty(0)
	for _, p := range properties {
		property |= p
	}
	queue := createCommandQueueWithProperties(c, device, property, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create command queue")
	}

	return queue, nil
}

func (c Context) Release() error {
	st := releaseContext(c)
	if st != CL_SUCCESS {
		return errors.New("oops at release context")
	}

	return nil
}

func (c Context) CreateProgram(source string) (Program, error) {
	var st clStatus

	program := createProgramWithSource(
		c, 1, []string{source}, []clSize{clSize(len(source))}, &st,
	)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create program with source")
	}

	return program, nil
}

func (c Context) CreateBuffer(flags []MemFlag, size uint) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := createBuffer(c, memFlags, clSize(size), nil, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create buffer: " + strconv.FormatInt(int64(st), 10))
	}

	return buffer, nil
}

func (c Context) CreateImage2D(flags []MemFlag, format ImageFormat, width, height int) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	w, h := clSize(width), clSize(height)
	buffer := createImage2D(c, memFlags, &format, w, h, 0, nil, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create image 2D: " + strconv.FormatInt(int64(st), 10))
	}

	return buffer, nil
}

// GL

func (c Context) CreateFromGLTexture(flags []MemFlag, target GLEnum, texture GLUint) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := createFromGLTexture(c, memFlags, target, 0, texture, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create from gl texture: " + strconv.FormatInt(int64(st), 10))
	}

	return buffer, nil
}
