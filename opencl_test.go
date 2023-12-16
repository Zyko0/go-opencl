package opencl_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/Zyko0/go-opencl"
)

const (
	dataSize = 32
)

var (
	code = string(`
	kernel void kern(global float* out)
	{
		size_t i = get_global_id(0);
		out[i] = i;
	}
	`)
)

// https://github.com/PassKeyRa/go-opencl/blob/master/opencl/external/include/CL/cl.h

func Test_Compute(m *testing.T) {
	err := opencl.Initialize()
	if err != nil {
		log.Fatal("err:", err)
	}

	platforms, err := opencl.GetPlatforms()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("platforms", len(platforms))

	name, err := platforms[0].GetName()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("name:", name)

	version, err := platforms[0].GetVersion()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("version:", version)

	platformExtensions, err := platforms[0].GetExtensions()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("platform extensions:", platformExtensions)

	devices, err := platforms[0].GetDevices(opencl.DeviceTypeAll)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("devices:", len(devices))

	deviceExtensions, err := devices[0].GetExtensions()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("device extensions:", deviceExtensions)

	ctx, err := devices[0].CreateContext(nil)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer ctx.Release()
	fmt.Println("context:", ctx)

	queue, err := ctx.CreateCommandQueue(devices[0])
	if err != nil {
		log.Fatal("err:", err)
	}
	defer queue.Release()
	fmt.Println("queue:", queue)

	program, err := ctx.CreateProgram(code)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer program.Release()
	fmt.Println("program:", program)

	logs, err := program.Build(devices[0], nil)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("logs:", logs)

	kernel, err := program.CreateKernel("kern")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer kernel.Release()
	fmt.Println("kernel:", kernel)

	data := make([]float32, dataSize)
	bufferData := opencl.GetBufferData(data)
	fmt.Println("buffer data:", bufferData.DataSize, bufferData.Pointer)
	buffer, err := ctx.CreateBuffer(
		[]opencl.MemFlag{
			opencl.MemFlagsWriteOnly,
			opencl.MemFlagsAllocHostPtr,
		},
		uint(bufferData.DataSize),
	)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer buffer.Release()
	fmt.Println("buffer:", buffer)

	size, err := buffer.Size()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("buffer size:", size)

	err = kernel.SetArg(0, opencl.NewKernelArg(&buffer))
	if err != nil {
		log.Fatal("err:", err)
	}

	err = queue.EnqueueNDRangeKernel(kernel, 1, nil, []uint64{dataSize}, nil)
	if err != nil {
		log.Fatal("err:", err)
	}

	queue.Flush()
	queue.Finish()

	err = queue.EnqueueReadBuffer(buffer, true, opencl.GetBufferData(data))
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("data:", data)
}
