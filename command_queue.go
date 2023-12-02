package opencl

import (
	"errors"
	"unsafe"
)

type commandQueueProperties uint32

type CommandQueue uint

func (cq CommandQueue) EnqueueNDRangeKernel(kernel Kernel, workDim uint, globalWorkSizes []uint64) error {
	sizes := make([]clSize, len(globalWorkSizes))
	for i := range globalWorkSizes {
		sizes[i] = clSize(globalWorkSizes[i])
	}
	st := enqueueNDRangeKernel(
		cq, kernel, workDim, nil, sizes, nil, 0, nil, nil,
	)
	if st != CL_SUCCESS {
		return errors.New("oops at enqueue nd range kernel")
	}

	return nil
}

func (cq CommandQueue) EnqueueReadBuffer(buffer Buffer, blockingRead bool, dataPtr any) error {
	slice := dataPtr.([]float32)
	ptr := unsafe.Pointer(&slice[0])
	dataLen := clSize(len(slice)) * clSize(unsafe.Sizeof(slice[0]))
	st := enqueueReadBuffer(
		cq, buffer, blockingRead, 0, dataLen, ptr, 0, nil, nil,
	)
	if st != CL_SUCCESS {
		return errors.New("oops at enqueue nd range kernel")
	}

	return nil
}

func (cq CommandQueue) Finish() error {
	st := finishCommandQueue(cq)
	if st != CL_SUCCESS {
		return errors.New("oops at finish command queue")
	}

	return nil
}

func (cq CommandQueue) Flush() error {
	st := flushCommandQueue(cq)
	if st != CL_SUCCESS {
		return errors.New("oops at flush command queue")
	}

	return nil
}

func (cq CommandQueue) Release() error {
	st := releaseCommandQueue(cq)
	if st != CL_SUCCESS {
		return errors.New("oops at release command queue")
	}

	return nil
}
