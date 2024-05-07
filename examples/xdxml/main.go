package main

import (
	"fmt"
	"log"

	"github.com/Longpai/xdxml-lib/pkg/xdxlib/xdxml"
)

func main() {
	xl := xdxml.New()
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

}
