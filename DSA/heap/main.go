package main

import (
	"fmt"
	"heap/heap"
)

func main() {

	heap, err := heap.NewHeap(heap.MAX)
	if err != nil {
		fmt.Println("Failed to create heap: ", err)
		return
	}

	heap.Add(-6)
	heap.Add(3)
	heap.Add(2)
	heap.Add(-7)
	heap.Add(-5)
	heap.Add(10)

	fmt.Println("Heap Current Status: ", heap.GetHeap())
	fmt.Println("Popping the element: ", heap.Pop())
	fmt.Println("Heap Current Status: ", heap.GetHeap())
}