package opencl

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	CL_SUCCESS = 0
)

type clSize uint
type clStatus int32

type createContextNotifyFunc func(errinfo, privateInfo []byte, cb clSize, userData []byte)
type buildProgramNotifyFunc func(program Program, userData []byte)

var (
	// Platform
	getPlatformIDs  func(numEntries uint32, platforms []Platform, numPlatforms *uint32) clStatus
	getPlatformInfo func(platform Platform, platformInfo platformInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *clSize) clStatus
	// Device
	getDeviceIDs  func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) clStatus
	getDeviceInfo func(device Device, deviceInfo DeviceInfo, paramValueSize clSize, paramValue []byte, paramValueSizeRet *uint32) clStatus
	// Context
	createContext           func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *createContextNotifyFunc, userData []byte, errCodeRet *clStatus) Context
	releaseContext          func(ctx Context) clStatus
	createProgramWithSource func(ctx Context, count clSize, strings []string, lengths []clSize, errCodeRet *clStatus) Program
	createBuffer            func(ctx Context, memFlags MemFlag, size clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer
	createImage2D           func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth, imageHeight, imageRowPitch clSize, hostPtr unsafe.Pointer, errCodeRet *clStatus) Buffer
	// Command queue
	createCommandQueue               func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue
	createCommandQueueWithProperties func(context Context, device Device, properties CommandQueueProperty, errCodeRet *clStatus) CommandQueue
	enqueueBarrier                   func(queue CommandQueue) clStatus
	enqueueNDRangeKernel             func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset, globalWorkSize, localWorkSize []clSize, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueReadBuffer                func(queue CommandQueue, buffer Buffer, blockingRead bool, offset, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueReadImage                 func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]clSize, row_pitch, slice_pitch clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueWriteBuffer               func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueMapBuffer                 func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr
	enqueueUnmapMemObject            func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueMapImage                  func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]clSize, imageRowPitch, imageSlicePitch *clSize, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *clStatus) uintptr
	enqueueAcquireGLObjects          func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) clStatus
	enqueueReleaseGLObjects          func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) clStatus
	finishCommandQueue               func(queue CommandQueue) clStatus
	flushCommandQueue                func(queue CommandQueue) clStatus
	releaseCommandQueue              func(queue CommandQueue) clStatus
	// Program
	buildProgram        func(program Program, numDevices uint32, devices []Device, options string, pfnNotify *buildProgramNotifyFunc, userData []byte) clStatus
	getProgramBuildInfo func(program Program, device Device, info programBuildInfo, paramSize clSize, paramValue unsafe.Pointer, paramSizeRet *clSize) clStatus
	createKernel        func(program Program, kernelName string, errCodeRet *clStatus) Kernel
	releaseProgram      func(program Program) clStatus
	// Kernel
	setKernelArg  func(kernel Kernel, argIndex uint, argSize clSize, argValue unsafe.Pointer) clStatus
	releaseKernel func(kernel Kernel) clStatus
	// Buffer
	getMemObjectInfo func(buffer Buffer, memInfo memInfo, paramValueSize clSize, paramValue unsafe.Pointer, paramValueSizeRet *clSize) clStatus
	releaseMemObject func(buffer Buffer) clStatus
	// GL
	createFromGLTexture func(ctx Context, memFlags MemFlag, textureTarget GLEnum, mipLevel GLInt, texture GLUint, errCodeRet *clStatus) Buffer
	getGLObjectInfo     func(memObj Buffer, objectType *CLGLObjectType, objectName *GLUint) clStatus
	getGLTextureInfo    func(memObj Buffer, paramName CLGLTextureInfo, paramValueSize clSize, paramValue unsafe.Pointer, paramValueSizeRet *clSize) clStatus
)

func Initialize() error {
	handle, err := loadLibrary()
	if err != nil {
		return err
	}

	// Platform
	purego.RegisterLibFunc(&getPlatformIDs, handle, "clGetPlatformIDs")
	purego.RegisterLibFunc(&getPlatformInfo, handle, "clGetPlatformInfo")
	// Device
	purego.RegisterLibFunc(&getDeviceIDs, handle, "clGetDeviceIDs")
	purego.RegisterLibFunc(&getDeviceInfo, handle, "clGetDeviceInfo")
	// Context
	purego.RegisterLibFunc(&createContext, handle, "clCreateContext")
	purego.RegisterLibFunc(&releaseContext, handle, "clReleaseContext")
	purego.RegisterLibFunc(&createProgramWithSource, handle, "clCreateProgramWithSource")
	purego.RegisterLibFunc(&createBuffer, handle, "clCreateBuffer")
	purego.RegisterLibFunc(&createImage2D, handle, "clCreateImage2D")
	// Command queue
	purego.RegisterLibFunc(&createCommandQueue, handle, "clCreateCommandQueue")
	purego.RegisterLibFunc(&createCommandQueueWithProperties, handle, "clCreateCommandQueueWithProperties")
	purego.RegisterLibFunc(&enqueueBarrier, handle, "clEnqueueBarrier")
	purego.RegisterLibFunc(&enqueueNDRangeKernel, handle, "clEnqueueNDRangeKernel")
	purego.RegisterLibFunc(&enqueueReadBuffer, handle, "clEnqueueReadBuffer")
	//TODO: purego: broken too many arguments //purego.RegisterLibFunc(&enqueueReadImage, handle, "clEnqueueReadImage")
	purego.RegisterLibFunc(&enqueueWriteBuffer, handle, "clEnqueueWriteBuffer")
	//TODO: purego: broken too many arguments // purego.RegisterLibFunc(&enqueueMapImage, handle, "clEnqueueMapImage")
	//TODO: purego: broken too many arguments // purego.RegisterLibFunc(&enqueueMapBuffer, handle, "clEnqueueMapBuffer")
	purego.RegisterLibFunc(&enqueueUnmapMemObject, handle, "clEnqueueUnmapMemObject")
	purego.RegisterLibFunc(&finishCommandQueue, handle, "clFinish")
	purego.RegisterLibFunc(&flushCommandQueue, handle, "clFlush")
	purego.RegisterLibFunc(&releaseCommandQueue, handle, "clReleaseCommandQueue")
	// Program
	purego.RegisterLibFunc(&buildProgram, handle, "clBuildProgram")
	purego.RegisterLibFunc(&getProgramBuildInfo, handle, "clGetProgramBuildInfo")
	purego.RegisterLibFunc(&createKernel, handle, "clCreateKernel")
	purego.RegisterLibFunc(&releaseProgram, handle, "clReleaseProgram")
	// Kernel
	purego.RegisterLibFunc(&setKernelArg, handle, "clSetKernelArg")
	purego.RegisterLibFunc(&releaseKernel, handle, "clReleaseKernel")
	// Buffer
	purego.RegisterLibFunc(&getMemObjectInfo, handle, "clGetMemObjectInfo")
	purego.RegisterLibFunc(&releaseMemObject, handle, "clReleaseMemObject")
	// GL
	purego.RegisterLibFunc(&createFromGLTexture, handle, "clCreateFromGLTexture")
	purego.RegisterLibFunc(&enqueueAcquireGLObjects, handle, "clEnqueueAcquireGLObjects")
	purego.RegisterLibFunc(&enqueueReleaseGLObjects, handle, "clEnqueueReleaseGLObjects")
	purego.RegisterLibFunc(&getGLObjectInfo, handle, "clGetGLObjectInfo")
	purego.RegisterLibFunc(&getGLTextureInfo, handle, "clGetGLTextureInfo")

	return nil
}
