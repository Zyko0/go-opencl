package opencl

import (
	"syscall"
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	CL_SUCCESS = 0
)

func loadLibrary() (uintptr, error) {
	// RTLD_NOW|RTLD_GLOBAL
	// Windows
	handle, err := syscall.LoadLibrary("opencl.dll")
	return uintptr(handle), err
}

type clSize uint64
type clStatus uint32

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
	createContext           func(properties *contextProperties, numDevices uint32, devices []Device, pfnNotify *createContextNotifyFunc, userData []byte, errCodeRet *clStatus) Context
	releaseContext          func(ctx Context) clStatus
	createProgramWithSource func(ctx Context, count clSize, strings []string, lenghts []clSize, errCodeRet *clStatus) Program
	createBuffer            func(ctx Context, memFlags MemFlag, size clSize, hostPtr *any, errCodeRet *clStatus) Buffer
	// Command queue
	createCommandQueue   func(context Context, device Device, properties commandQueueProperties, errCodeRet *clStatus) CommandQueue
	enqueueNDRangeKernel func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset, globalWorkSize, localWorkSize []clSize, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	enqueueReadBuffer    func(queue CommandQueue, buffer Buffer, blockingRead bool, offset, cb clSize, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) clStatus
	finishCommandQueue   func(queue CommandQueue) clStatus
	flushCommandQueue    func(queue CommandQueue) clStatus
	releaseCommandQueue  func(queue CommandQueue) clStatus
	// Program
	buildProgram   func(program Program, numDevices uint32, devices []Device, options string, pfnNotify *buildProgramNotifyFunc, userData []byte) clStatus
	createKernel   func(program Program, kernelName string, errCodeRet *clStatus) Kernel
	releaseProgram func(program Program) clStatus
	// Kernel
	setKernelArg  func(kernel Kernel, argIndex uint, argSize clSize, argValue unsafe.Pointer) clStatus
	releaseKernel func(kernel Kernel) clStatus
	// Buffer
	getMemObjectInfo func(buffer Buffer, memInfo memInfo, paramValueSize clSize, paramValue *uint, paramValueSizeRet *clSize) clStatus
	releaseMemObject func(buffer Buffer) clStatus
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
	// Command queue
	purego.RegisterLibFunc(&createCommandQueue, handle, "clCreateCommandQueue")
	purego.RegisterLibFunc(&enqueueNDRangeKernel, handle, "clEnqueueNDRangeKernel")
	purego.RegisterLibFunc(&enqueueReadBuffer, handle, "clEnqueueReadBuffer")
	purego.RegisterLibFunc(&finishCommandQueue, handle, "clFinish")
	purego.RegisterLibFunc(&flushCommandQueue, handle, "clFlush")
	purego.RegisterLibFunc(&releaseCommandQueue, handle, "clReleaseCommandQueue")
	// Program
	purego.RegisterLibFunc(&buildProgram, handle, "clBuildProgram")
	purego.RegisterLibFunc(&releaseProgram, handle, "clReleaseProgram")
	purego.RegisterLibFunc(&createKernel, handle, "clCreateKernel")
	// Kernel
	purego.RegisterLibFunc(&setKernelArg, handle, "clSetKernelArg")
	purego.RegisterLibFunc(&releaseKernel, handle, "clReleaseKernel")
	// Buffer
	purego.RegisterLibFunc(&getMemObjectInfo, handle, "clGetMemObjectInfo")
	purego.RegisterLibFunc(&releaseMemObject, handle, "clReleaseMemObject")

	return nil
}
