package main

import "fmt"

func main() {
	code := []Address{
		{addr: 100, dest: 0},
		{addr: 101, dest: 0},
		{addr: 200, dest: 0},
	}
	heap := []Address{
		{addr: 300, dest: 1},
		{addr: 301, dest: 1},
		{addr: 500, dest: 1},
	}
	stack := []Address{
		{addr: 900, dest: 2},
		{addr: 899, dest: 2},
		{addr: 1600, dest: 2},
	}

	job := &Proc{id: 1, length: 1000, code: code, heap: heap, stack: stack}
	mmu := &MMU{
		codeReg:  Register{base: 99, bound: 202, direction: 1},
		heapReg:  Register{base: 299, bound: 610, direction: 1},
		stackReg: Register{base: 1901, bound: 800, direction: -1},
	}
	runWithSegmentation(job, mmu)

}

type Register struct {
	base      int
	bound     int
	direction int
}

type MMU struct {
	codeReg  Register
	stackReg Register
	heapReg  Register
}

type Address struct {
	addr int
	dest int
}

type Proc struct {
	id     int
	length int
	code   []Address
	heap   []Address
	stack  []Address
}

func runWithSegmentation(job *Proc, mmu *MMU) {
	for _, adr := range job.code {
		if adr.dest == 0 {
			destAddr := mmu.codeReg.base + (mmu.codeReg.direction * adr.addr)
			if destAddr > mmu.codeReg.bound {
				fmt.Println("Segmentation Fault on address", destAddr, "for code")
			} else {
				fmt.Println("Code Address", destAddr, "accessed successfully")
			}
		}
	}
	for _, adr := range job.heap {
		if adr.dest == 1 {
			destAddr := mmu.heapReg.base + (mmu.heapReg.direction * adr.addr)
			if destAddr > mmu.heapReg.bound {
				fmt.Println("Segmentation Fault on address", destAddr, "for heap")
			} else {
				fmt.Println("Heap Address", destAddr, "accessed successfully")
			}
		}
	}
	for _, adr := range job.stack {
		if adr.dest == 2 {
			destAddr := mmu.stackReg.base + (mmu.stackReg.direction * adr.addr)
			if destAddr < mmu.stackReg.bound {
				fmt.Println("Segmentation Fault on address", destAddr, "for stack")
			} else {
				fmt.Println("Stack Address", destAddr, "accessed successfully")
			}
		}
	}
}
