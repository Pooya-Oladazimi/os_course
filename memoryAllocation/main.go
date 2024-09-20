package main

import "fmt"

type FreeListNode struct {
	start  int
	length int
	next   *FreeListNode
}

func main() {
	freeList := &FreeListNode{start: 0, length: 1024, next: nil}
	newFreeList, err := firstFit(200, freeList)
	newFreeList, err = firstFit(700, newFreeList)
	newFreeList, err = firstFit(200, newFreeList)
	if err {
		fmt.Println("Cannot allocate memory. free space is ", getFreeMemory(newFreeList))
	} else {
		fmt.Println("Memory is allocated successfully")
	}

}

func firstFit(addressSpaceLength int, freeList *FreeListNode) (*FreeListNode, bool) {
	cNode := freeList
	var prevNode *FreeListNode = nil
	for cNode != nil {
		if cNode.length == addressSpaceLength {
			if prevNode == nil {
				cNode = nil
				return freeList, false
			}
			prevNode.next = cNode.next
			return freeList, false
		} else if cNode.length > addressSpaceLength {
			cNode.length -= addressSpaceLength
			cNode.start += addressSpaceLength
			return freeList, false
		} else {
			prevNode = cNode
			cNode = cNode.next
		}
	}
	return freeList, true
}

func getFreeMemory(freeList *FreeListNode) (memSize int) {
	cNode := freeList
	memSize = 0
	for cNode != nil {
		memSize += cNode.length
		cNode = cNode.next
	}
	return memSize
}
