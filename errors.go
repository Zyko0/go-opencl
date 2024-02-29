package opencl

import "C"
import (
	"fmt"
	"github.com/Zyko0/go-opencl/constants"
)

type CLError struct {
	e clStatus
}

func (e clStatus) getError() error {
	if e == CL_SUCCESS {
		return nil
	}
	return &CLError{e: e}
}

func (e *CLError) Error() string {
	if err, ok := errorMap2[e.e]; ok {
		return err
	}
	return fmt.Sprintf("cl: error %d", e.e)
}

// Common OpenCl errors
const (
	ErrDeviceNotFound                     = "cl: Device Not Found"
	ErrDeviceNotAvailable                 = "cl: Device Not Available"
	ErrCompilerNotAvailable               = "cl: Compiler Not Available"
	ErrMemObjectAllocationFailure         = "cl: Mem Object Allocation Failure"
	ErrOutOfResources                     = "cl: Out Of Resources"
	ErrOutOfHostMemory                    = "cl: Out Of Host Memory"
	ErrProfilingInfoNotAvailable          = "cl: Profiling Info Not Available"
	ErrMemCopyOverlap                     = "cl: Mem Copy Overlap"
	ErrImageFormatMismatch                = "cl: Image Format Mismatch"
	ErrImageFormatNotSupported            = "cl: Image Format Not Supported"
	ErrBuildProgramFailure                = "cl: Build Program Failure"
	ErrMapFailure                         = "cl: Map Failure"
	ErrMisalignedSubBufferOffset          = "cl: Misaligned Sub clBuffer Offset"
	ErrExecStatusErrorForEventsInWaitList = "cl: Exec Status Error For Events In Wait List"
	ErrCompileProgramFailure              = "cl: Compile Program Failure"
	ErrLinkerNotAvailable                 = "cl: Linker Not Available"
	ErrLinkProgramFailure                 = "cl: Link Program Failure"
	ErrDevicePartitionFailed              = "cl: Device Partition Failed"
	ErrKernelArgInfoNotAvailable          = "cl: Kernel Arg Info Not Available"
	ErrInvalidValue                       = "cl: Invalid Value"
	ErrInvalidDeviceType                  = "cl: Invalid Device Type"
	ErrInvalidPlatform                    = "cl: Invalid Platform"
	ErrInvalidDevice                      = "cl: Invalid Device"
	ErrInvalidContext                     = "cl: Invalid clContext"
	ErrInvalidQueueProperties             = "cl: Invalid Queue Properties"
	ErrInvalidCommandQueue                = "cl: Invalid Command Queue"
	ErrInvalidHostPtr                     = "cl: Invalid Host Ptr"
	ErrInvalidMemObject                   = "cl: Invalid Mem Object"
	ErrInvalidImageFormatDescriptor       = "cl: Invalid Image Format Descriptor"
	ErrInvalidImageSize                   = "cl: Invalid Image Size"
	ErrInvalidSampler                     = "cl: Invalid Sampler"
	ErrInvalidBinary                      = "cl: Invalid Binary"
	ErrInvalidBuildOptions                = "cl: Invalid Build Options"
	ErrInvalidProgram                     = "cl: Invalid Program"
	ErrInvalidProgramExecutable           = "cl: Invalid Program Executable"
	ErrInvalidKernelName                  = "cl: Invalid Kernel Name"
	ErrInvalidKernelDefinition            = "cl: Invalid Kernel Definition"
	ErrInvalidKernel                      = "cl: Invalid Kernel"
	ErrInvalidArgIndex                    = "cl: Invalid Arg Index"
	ErrInvalidArgValue                    = "cl: Invalid Arg Value"
	ErrInvalidArgSize                     = "cl: Invalid Arg Size"
	ErrInvalidKernelArgs                  = "cl: Invalid Kernel Args"
	ErrInvalidWorkDimension               = "cl: Invalid Work Dimension"
	ErrInvalidWorkGroupSize               = "cl: Invalid Work Group Size"
	ErrInvalidWorkItemSize                = "cl: Invalid Work Item Size"
	ErrInvalidGlobalOffset                = "cl: Invalid Global Offset"
	ErrInvalidEventWaitList               = "cl: Invalid Event Wait List"
	ErrInvalidEvent                       = "cl: Invalid Event"
	ErrInvalidOperation                   = "cl: Invalid Operation"
	ErrInvalidGlObject                    = "cl: Invalid Gl Object"
	ErrInvalidBufferSize                  = "cl: Invalid clBuffer Size"
	ErrInvalidMipLevel                    = "cl: Invalid Mip Level"
	ErrInvalidGlobalWorkSize              = "cl: Invalid Global Work Size"
	ErrInvalidProperty                    = "cl: Invalid Property"
	ErrInvalidImageDescriptor             = "cl: Invalid Image Descriptor"
	ErrInvalidCompilerOptions             = "cl: Invalid Compiler Options"
	ErrInvalidLinkerOptions               = "cl: Invalid Linker Options"
	ErrInvalidDevicePartitionCount        = "cl: Invalid Device Partition Count"
)

var errorMap2 = map[clStatus]string{
	constants.CL_SUCCESS:                                   "",
	constants.CL_DEVICE_NOT_FOUND:                          ErrDeviceNotFound,
	constants.CL_DEVICE_NOT_AVAILABLE:                      ErrDeviceNotAvailable,
	constants.CL_COMPILER_NOT_AVAILABLE:                    ErrCompilerNotAvailable,
	constants.CL_MEM_OBJECT_ALLOCATION_FAILURE:             ErrMemObjectAllocationFailure,
	constants.CL_OUT_OF_RESOURCES:                          ErrOutOfResources,
	constants.CL_OUT_OF_HOST_MEMORY:                        ErrOutOfHostMemory,
	constants.CL_PROFILING_INFO_NOT_AVAILABLE:              ErrProfilingInfoNotAvailable,
	constants.CL_MEM_COPY_OVERLAP:                          ErrMemCopyOverlap,
	constants.CL_IMAGE_FORMAT_MISMATCH:                     ErrImageFormatMismatch,
	constants.CL_IMAGE_FORMAT_NOT_SUPPORTED:                ErrImageFormatNotSupported,
	constants.CL_BUILD_PROGRAM_FAILURE:                     ErrBuildProgramFailure,
	constants.CL_MAP_FAILURE:                               ErrMapFailure,
	constants.CL_MISALIGNED_SUB_BUFFER_OFFSET:              ErrMisalignedSubBufferOffset,
	constants.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST: ErrExecStatusErrorForEventsInWaitList,
	constants.CL_INVALID_VALUE:                             ErrInvalidValue,
	constants.CL_INVALID_DEVICE_TYPE:                       ErrInvalidDeviceType,
	constants.CL_INVALID_PLATFORM:                          ErrInvalidPlatform,
	constants.CL_INVALID_DEVICE:                            ErrInvalidDevice,
	constants.CL_INVALID_CONTEXT:                           ErrInvalidContext,
	constants.CL_INVALID_QUEUE_PROPERTIES:                  ErrInvalidQueueProperties,
	constants.CL_INVALID_COMMAND_QUEUE:                     ErrInvalidCommandQueue,
	constants.CL_INVALID_HOST_PTR:                          ErrInvalidHostPtr,
	constants.CL_INVALID_MEM_OBJECT:                        ErrInvalidMemObject,
	constants.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR:           ErrInvalidImageFormatDescriptor,
	constants.CL_INVALID_IMAGE_SIZE:                        ErrInvalidImageSize,
	constants.CL_INVALID_SAMPLER:                           ErrInvalidSampler,
	constants.CL_INVALID_BINARY:                            ErrInvalidBinary,
	constants.CL_INVALID_BUILD_OPTIONS:                     ErrInvalidBuildOptions,
	constants.CL_INVALID_PROGRAM:                           ErrInvalidProgram,
	constants.CL_INVALID_PROGRAM_EXECUTABLE:                ErrInvalidProgramExecutable,
	constants.CL_INVALID_KERNEL_NAME:                       ErrInvalidKernelName,
	constants.CL_INVALID_KERNEL_DEFINITION:                 ErrInvalidKernelDefinition,
	constants.CL_INVALID_KERNEL:                            ErrInvalidKernel,
	constants.CL_INVALID_ARG_INDEX:                         ErrInvalidArgIndex,
	constants.CL_INVALID_ARG_VALUE:                         ErrInvalidArgValue,
	constants.CL_INVALID_ARG_SIZE:                          ErrInvalidArgSize,
	constants.CL_INVALID_KERNEL_ARGS:                       ErrInvalidKernelArgs,
	constants.CL_INVALID_WORK_DIMENSION:                    ErrInvalidWorkDimension,
	constants.CL_INVALID_WORK_GROUP_SIZE:                   ErrInvalidWorkGroupSize,
	constants.CL_INVALID_WORK_ITEM_SIZE:                    ErrInvalidWorkItemSize,
	constants.CL_INVALID_GLOBAL_OFFSET:                     ErrInvalidGlobalOffset,
	constants.CL_INVALID_EVENT_WAIT_LIST:                   ErrInvalidEventWaitList,
	constants.CL_INVALID_EVENT:                             ErrInvalidEvent,
	constants.CL_INVALID_OPERATION:                         ErrInvalidOperation,
	constants.CL_INVALID_GL_OBJECT:                         ErrInvalidGlObject,
	constants.CL_INVALID_BUFFER_SIZE:                       ErrInvalidBufferSize,
	constants.CL_INVALID_MIP_LEVEL:                         ErrInvalidMipLevel,
	constants.CL_INVALID_GLOBAL_WORK_SIZE:                  ErrInvalidGlobalWorkSize,
	constants.CL_INVALID_PROPERTY:                          ErrInvalidProperty,
}
