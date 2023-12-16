package opencl

import (
	"errors"
	"strings"
	"unsafe"
)

type Program uint

type programBuildInfo uint32

const (
	programBuildStatus  programBuildInfo = 0x1181
	programBuildOptions programBuildInfo = 0x1182
	programBuildLog     programBuildInfo = 0x1183
)

type Version string

const (
	Version1_0 Version = "CL1.0"
	Version1_1 Version = "CL1.1"
	Version1_2 Version = "CL1.2"
	Version2_0 Version = "CL2.0"
	Version3_0 Version = "CL3.0"
)

type ProgramBuildOptions struct {
	// Preprocessor options
	Warnings          bool
	Macros            map[string]string
	DirectoryIncludes []string
	Version           Version
	// Math intrinsics options
	SinglePrecisionConstant bool
	MadEnable               bool
	NoSignedZeros           bool
	FastRelaxedMaths        bool
	// Extensions
	NvidiaVerbose bool
}

func (po *ProgramBuildOptions) String() string {
	if po == nil {
		return ""
	}

	var sb strings.Builder

	// Preprocessor
	if po.Warnings {
		sb.WriteString("-w")
		sb.WriteRune(' ')
	}
	if po.Version != "" {
		sb.WriteString("-cl-std=" + string(po.Version))
		sb.WriteRune(' ')
	}
	// Math intrinsics
	if po.SinglePrecisionConstant {
		sb.WriteString("-cl-single-precision-constant")
		sb.WriteRune(' ')
	}
	if po.MadEnable {
		sb.WriteString("-cl-mad-enable")
		sb.WriteRune(' ')
	}
	if po.NoSignedZeros {
		sb.WriteString("-cl-no-signed-zeros")
		sb.WriteRune(' ')
	}
	if po.FastRelaxedMaths {
		sb.WriteString("-cl-fast-relaxed-math")
		sb.WriteRune(' ')
	}
	// Extensions
	if po.NvidiaVerbose {
		sb.WriteString("-cl-nv-verbose")
		sb.WriteRune(' ')
	}

	return sb.String()
}

func (p Program) Build(device Device, opts *ProgramBuildOptions) (string, error) {
	var err error

	st := buildProgram(
		p, 1, []Device{device}, opts.String(), nil, nil,
	)
	if st != CL_SUCCESS {
		err = errors.New("oops at build program")
	}

	var logsSize clSize
	st = getProgramBuildInfo(
		p, device, programBuildLog, 0, nil, &logsSize,
	)
	if st != CL_SUCCESS {
		return "", errors.New("oops at 1st get program build info")
	}

	var logs = make([]byte, logsSize)
	st = getProgramBuildInfo(
		p, device, programBuildLog, logsSize, unsafe.Pointer(&logs[0]), nil,
	)
	if st != CL_SUCCESS {
		return "", errors.New("oops at 2nd get program build info")
	}

	return string(logs), err
}

func (p Program) CreateKernel(name string) (Kernel, error) {
	var st clStatus

	kernel := createKernel(p, name, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create kernel")
	}

	return kernel, nil
}

func (p Program) Release() error {
	st := releaseProgram(p)
	if st != CL_SUCCESS {
		return errors.New("oops at release program")
	}

	return nil
}
