package opencl

import "errors"

type Program uint

func (p Program) Build(device Device) error {
	st := buildProgram(
		p, 1, []Device{device}, "", nil, nil,
	)
	if st != CL_SUCCESS {
		return errors.New("oops at build program")
	}

	return nil
}

func (p Program) Release() error {
	st := releaseProgram(p)
	if st != CL_SUCCESS {
		return errors.New("oops at release program")
	}

	return nil
}

func (p Program) CreateKernel(name string) (Kernel, error) {
	var st clStatus

	kernel := createKernel(p, name, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create kernel")
	}

	return kernel, nil
}
