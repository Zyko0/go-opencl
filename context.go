package opencl

import (
	"unsafe"
)

type contextProperty uint

type ContextProperties struct {
	Platform *Platform
	// Interop
	InteropUserSync *bool
	// OpenGL
	GLContextKHR *uint
	// Windows
	WGL_HDC_KHR *uint
}

func (cp *ContextProperties) compile() []contextProperty {
	if cp == nil {
		return []contextProperty{0}
	}

	const (
		ContextPropertyPlatform        contextProperty = 0x1084
		ContextPropertyInteropUserSync contextProperty = 0x1085 // >= 1.2
		// GL
		ContextPropertyGLContextKHR contextProperty = 0x2008
		ContextPropertyWGL_HDC_KHR  contextProperty = 0x200B
	)

	var properties []contextProperty
	if cp.Platform != nil {
		properties = append(properties, ContextPropertyPlatform, contextProperty(*cp.Platform))
	}
	if cp.InteropUserSync != nil {
		b := contextProperty(0)
		if *cp.InteropUserSync {
			b = 1
		}
		properties = append(properties, ContextPropertyInteropUserSync, b)
	}
	if cp.GLContextKHR != nil {
		properties = append(properties, ContextPropertyGLContextKHR, contextProperty(*cp.GLContextKHR))
	}
	if cp.WGL_HDC_KHR != nil {
		properties = append(properties, ContextPropertyWGL_HDC_KHR, contextProperty(*cp.WGL_HDC_KHR))
	}
	// End of list should be marked as an extra zero
	return append(properties, 0)
}

type Context uint

// TODO: make properties into a struct instead of weird map<uint32>

func (d Device) CreateContext(properties *ContextProperties) (Context, error) {
	var st clStatus

	flattened := properties.compile()
	ctx := createContext(unsafe.Pointer(&flattened[0]), 1, []Device{d}, nil, nil, &st)
	return ctx, st.getError()
}

func (c Context) CreateCommandQueue(device Device) (CommandQueue, error) {
	var st clStatus
	queue := createCommandQueue(c, device, 0, &st)
	return queue, st.getError()
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
	return queue, st.getError()
}

func (c Context) Release() error {
	return releaseContext(c).getError()
}

func (c Context) CreateProgram(source string) (Program, error) {
	var st clStatus

	program := createProgramWithSource(
		c, 1, []string{source}, []clSize{clSize(len(source))}, &st,
	)
	return program, st.getError()
}

func (c Context) CreateBuffer(flags []MemFlag, size uint) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := createBuffer(c, memFlags, clSize(size), nil, &st)
	return buffer, st.getError()
}

func (c Context) CreateImage2D(flags []MemFlag, format ImageFormat, width, height int) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	w, h := clSize(width), clSize(height)
	buffer := createImage2D(c, memFlags, &format, w, h, 0, nil, &st)
	return buffer, st.getError()
}

// GL

func (c Context) CreateFromGLTexture(flags []MemFlag, target GLEnum, texture GLUint) (Buffer, error) {
	var st clStatus

	memFlags := MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := createFromGLTexture(c, memFlags, target, 0, texture, &st)
	return buffer, st.getError()
}
