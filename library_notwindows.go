//go:build darwin || freebsd || linux

package opencl

import (
	"errors"
	"github.com/ebitengine/purego"
	"runtime"
)

func getOpenCLPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return "libOpenCL.so", nil
	default:
		return "", errors.New("unsupported operating system")
	}
}

func loadLibrary() (uintptr, error) {
	path, err := getOpenCLPath()
	if err != nil {
		return 0, err
	}
	handle, err := purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, err
	}

	return handle, err
}
