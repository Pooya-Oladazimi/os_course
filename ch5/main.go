package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println("I am the process id ", os.Getpid(), "and about to fork a child")
	rc, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if rc < 0 {
		fmt.Println("Fork failed!")
		syscall.Exit(1)
	} else if rc == 0 {
		fmt.Println("I am the child: ", os.Getpid())
	} else {
		wait(int(rc))
		fmt.Println("I am the parent: ", os.Getpid(), " and my child is: ", int(rc))
	}
}

func wait(pid int) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}
	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}
}
