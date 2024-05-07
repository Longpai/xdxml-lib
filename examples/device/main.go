package main

import (
	"fmt"
	"log"

	"github.com/Longpai/xdxml-lib/pkg/xdxlib/xdxml"
)

func main() {
	xl := xdxml.New()
	// dl := device.New(device.WithXdxml(xl))
	ret := xl.Init()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to initialize XDXML: %v", ret)
		return
	}
	defer func() {
		ret := xl.Shutdown()
		if ret != xdxml.SUCCESS {
			log.Fatalf("Unable to shutdown XDXML: %v", ret)
		}
	}()

	count, ret := xl.DeviceGetCount()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to get Conut: %v", ret)
		return
	}
	fmt.Printf("Count: %v\n", count)

	// var d xdxml.Device
	d, r := xl.DeviceGetHandleByIndex(0)
	if r != xdxml.SUCCESS {
		log.Fatalf("Unable to get handle: %v", ret)
		return
	}

	memory_st, ret := d.GetMemoryInfo()
	if ret != xdxml.SUCCESS {
		log.Fatalf("Unable to get mem of device at index %v: %v", 0, ret)
	}
	fmt.Printf("memory: %v MB\n", memory_st.Fb_total/(1024*1024))

}
