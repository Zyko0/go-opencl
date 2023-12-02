package opencl

import "errors"

type DeviceType uint32

const (
	DeviceTypeDefault     DeviceType = 1 << 0
	DeviceTypeCPU         DeviceType = 1 << 1
	DeviceTypeGPU         DeviceType = 1 << 2
	DeviceTypeAccelerator DeviceType = 1 << 3
	DeviceTypeCustom      DeviceType = 1 << 4
	DeviceTypeAll         DeviceType = 0xFFFFFFFF
)

type DeviceInfo uint32

const (
	DeviceInfoType              DeviceInfo = 0x1000
	DeviceInfoAddressBits       DeviceInfo = 0x100D
	DeviceInfoAvailable         DeviceInfo = 0x1027
	DeviceInfoCompilerAvailable DeviceInfo = 0x1028
	DeviceInfoBuiltInKernels    DeviceInfo = 0x103F
	DeviceInfoVendor            DeviceInfo = 0x102C
	DeviceInfoDriverVersion     DeviceInfo = 0x102D
)

type Device uint

func (d Device) CreateContext() (Context, error) {
	var st clStatus

	ctx := createContext(nil, 1, []Device{d}, nil, nil, &st)
	if st != CL_SUCCESS {
		return 0, errors.New("oops at create context")
	}

	return ctx, nil
}
