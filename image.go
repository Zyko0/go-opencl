package opencl

type ImageChannelOrder uint32

const (
	ImageChannelOrder_R         ImageChannelOrder = 0x10B0
	ImageChannelOrder_A         ImageChannelOrder = 0x10B1
	ImageChannelOrder_RG        ImageChannelOrder = 0x10B2
	ImageChannelOrder_RA        ImageChannelOrder = 0x10B3
	ImageChannelOrder_RGB       ImageChannelOrder = 0x10B4
	ImageChannelOrder_RGBA      ImageChannelOrder = 0x10B5
	ImageChannelOrder_BGRA      ImageChannelOrder = 0x10B6
	ImageChannelOrder_ARGB      ImageChannelOrder = 0x10B7
	ImageChannelOrder_Intensity ImageChannelOrder = 0x10B8
	ImageChannelOrder_Luminance ImageChannelOrder = 0x10B9
)

type ImageChannelType uint32

const (
	ImageChannelType_SNormInt8      ImageChannelType = 0x10D0
	ImageChannelType_SNormInt16     ImageChannelType = 0x10D1
	ImageChannelType_UNormInt8      ImageChannelType = 0x10D2
	ImageChannelType_UNormInt16     ImageChannelType = 0x10D3
	ImageChannelType_UNormShort565  ImageChannelType = 0x10D4
	ImageChannelType_UNormShort555  ImageChannelType = 0x10D5
	ImageChannelType_UNormInt101010 ImageChannelType = 0x10D6
	ImageChannelType_SignedInt8     ImageChannelType = 0x10D7
	ImageChannelType_SignedInt16    ImageChannelType = 0x10D8
	ImageChannelType_SignedInt32    ImageChannelType = 0x10D9
	ImageChannelType_UnsignedInt8   ImageChannelType = 0x10DA
	ImageChannelType_UnsignedInt16  ImageChannelType = 0x10DB
	ImageChannelType_UnsignedInt32  ImageChannelType = 0x10DC
	ImageChannelType_HalfFloat      ImageChannelType = 0x10DD
	ImageChannelType_Float          ImageChannelType = 0x10DE
	// CL_VERSION 2.1
	ImageChannelType_UNormInt101010_2 ImageChannelType = 0x10E0
)

type ImageFormat struct {
	ChannelOrder ImageChannelOrder
	ChannelType  ImageChannelType
}

type ImageData struct {
	*BufferData
	Origin     [3]uint
	Region     [3]uint
	RowPitch   uint
	SlicePitch uint
}
