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

	info, err := platforms[0].GetName()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("info:", info)

	devices, err := platforms[0].GetDevices(opencl.DeviceTypeAll)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("devices:", devices)

	ctx, err := devices[0].CreateContext()
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

	err = program.Build(devices[0])
	if err != nil {
		log.Fatal("err:", err)
	}

	kernel, err := program.CreateKernel("kern")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer kernel.Release()
	fmt.Println("kernel:", kernel)

	buffer, err := ctx.CreateBuffer([]opencl.MemFlag{opencl.MemFlagsWriteOnly}, dataSize*4)
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

	err = kernel.SetArg(0, &buffer)
	if err != nil {
		log.Fatal("err:", err)
	}

	err = queue.EnqueueNDRangeKernel(kernel, 1, []uint64{dataSize})
	if err != nil {
		log.Fatal("err:", err)
	}

	queue.Flush()
	queue.Finish()

	data := make([]float32, dataSize)
	err = queue.EnqueueReadBuffer(buffer, true, data)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("data:", data)
}
