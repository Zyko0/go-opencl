package opencl

import (
	"errors"
)

type contextProperties uint32

type Context uint

func (c Context) CreateCommandQueue(device Device) (CommandQueue, error) {
	var st clStatus

	queue := createCommandQueue(c, device, 0, &st)
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
		memFlags &= f
	}
	buffer := createBuffer(c, memFlags, clSize(size), nil, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create buffer")
	}

	return buffer, nil
}
