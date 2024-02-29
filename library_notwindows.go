//go:build darwin || freebsd || linux || android

package opencl

import (
	"errors"
	"github.com/ebitengine/purego"
	"runtime"
)

func getOpenCLPath() []string {
	if runtime.GOOS == "linux" {
		return []string{
			"/usr/lib/libOpenCL.so",
			"/usr/local/lib/libOpenCL.so",
			"/usr/local/lib/libpocl.so",
			"/usr/lib64/libOpenCL.so",
			"/usr/lib32/libOpenCL.so",
			"libOpenCL.so"}
	} else if runtime.GOOS == "darwin" {
		return []string{
			"libOpenCL.so",
			"/System/Library/Frameworks/OpenCL.framework/OpenCL"}
	} else if runtime.GOOS == "android" {
		return []string{
			"/system/lib64/libOpenCL.so",
			"/system/vendor/lib64/libOpenCL.so",
			"/system/vendor/lib64/egl/libGLES_mali.so",
			"/system/vendor/lib64/libPVROCL.so",
			"/data/data/org.pocl.libs/files/lib64/libpocl.so",
			"/system/lib/libOpenCL.so",
			"/system/vendor/lib/libOpenCL.so",
			"/system/vendor/lib/egl/libGLES_mali.so",
			"/system/lib64/egl/libGLES_mali.so",
			"/system/vendor/lib/libPVROCL.so",
			"/data/data/org.pocl.libs/files/lib/libpocl.so",
			"libOpenCL.so"}
	}
	return make([]string, 0)
}

func loadLibrary() (uintptr, error) {
	paths := getOpenCLPath()
	if len(paths) == 0 {
		return 0, errors.New("unknown system paths")
	}
	for i := 0; i < len(paths); i++ {
		libOpenCl, err := purego.Dlopen(paths[i], purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err == nil {
			return libOpenCl, err
		}
	}
	return 0, errors.New("no path has passed")
}
