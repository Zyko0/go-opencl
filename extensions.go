package opencl

type Extension = string

// TODO: Not exhaustive
const (
	// OpenCL
	Extension_khr_gl_sharing Extension = "cl_khr_gl_sharing"
	Extension_khr_fp64       Extension = "cl_khr_fp64"
	// Nvidia
	Extension_nv_pragma_unroll    Extension = "cl_nv_pragma_unroll"
	Extension_nv_compiler_options Extension = "cl_nv_compiler_options"
)
