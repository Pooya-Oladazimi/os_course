/*
	Limited Direct Execution
*/

package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Average system call time is:", measueSysCall(), "μs")
}

func measueSysCall() float64 {
	total := float64(0)
	for i := 0; i < 10000; i++ {
		start := time.Now()
		syscall.Getpid()
		total += float64(time.Since(start).Microseconds())
	}
	return (total / 10000)
}
