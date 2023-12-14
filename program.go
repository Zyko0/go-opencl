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

type CLVersion string

const (
	CLVersion1_0 CLVersion = "CL1.0"
	CLVersion1_1 CLVersion = "CL1.1"
	CLVersion1_2 CLVersion = "CL1.2"
	CLVersion2_0 CLVersion = "CL2.0"
	CLVersion3_0 CLVersion = "CL3.0"
)

type ProgramBuildOptions struct {
	// Preprocessor options
	Warnings          bool
	Macros            map[string]string
	DirectoryIncludes []string
	CLVersion         CLVersion
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
	if po.CLVersion != "" {
		sb.WriteString("-cl-std=" + string(po.CLVersion))
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
