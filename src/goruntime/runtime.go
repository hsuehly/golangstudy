package goruntime

import (
	"fmt"
	"runtime"
)

func Runtime() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu数量", cpuNum)
}
