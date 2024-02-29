package opencl

import (
	"strings"
)

type DeviceType uint32

const (
	DeviceTypeDefault     DeviceType = 1 << 0
	DeviceTypeCPU         DeviceType = 1 << 1
	DeviceTypeGPU         DeviceType = 1 << 2
	DeviceTypeAccelerator DeviceType = 1 << 3
	DeviceTypeCustom      DeviceType = 1 << 4
	DeviceTypeAll         DeviceType = 0xFFFFFFFF
)

type Device uint

type deviceInfo uint32

const (
	deviceInfoType              deviceInfo = 0x1000
	deviceInfoAddressBits       deviceInfo = 0x100D
	deviceInfoAvailable         deviceInfo = 0x1027
	deviceInfoCompilerAvailable deviceInfo = 0x1028
	deviceInfoBuiltInKernels    deviceInfo = 0x103F
	deviceInfoVendor            deviceInfo = 0x102C
	deviceInfoDriverVersion     deviceInfo = 0x102D

	deviceInfoExtensions deviceInfo = 0x1030
)

func (d Device) getInfo(name deviceInfo) (string, error) {
	size := clSize(0)
	st := getDeviceInfo(d, name, clSize(0), nil, &size)
	if st != CL_SUCCESS {
		return "", st.getError()
	}
	info := make([]byte, size)
	st = getDeviceInfo(d, name, size, info, nil)
	return string(info), st.getError()
}

func (d Device) GetExtensions() ([]Extension, error) {
	extensions, err := d.getInfo(deviceInfoExtensions)
	if err != nil {
		return nil, err
	}
	return strings.Split(extensions, " "), nil
}
