package opencl

type GLInt int32
type GLEnum uint32
type GLUint uint32

// Buffer

type CLGLObjectType uint32

const (
	CLGLObjectBuffer       CLGLObjectType = 0x2000
	CLGLObjectTexture2D    CLGLObjectType = 0x2001
	CLGLObjectTexture3D    CLGLObjectType = 0x2002
	CLGLObjectRenderBuffer CLGLObjectType = 0x2003
	// ...
	CLGLObjectTextureBuffer CLGLObjectType = 0x2011
)

type CLGLTextureInfo uint32

const (
	CLGLTextureTarget      CLGLTextureInfo = 0x2004
	CLGLTextureMipmapLevel CLGLTextureInfo = 0x2005
)

// Context

const (
	GLTextureTarget2D GLEnum = 0x0DE1
)
