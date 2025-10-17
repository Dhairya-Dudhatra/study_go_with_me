package heap

import (
"fmt"
"log/slog"
)

type heapType string

const (
	MIN heapType = "min"
	MAX heapType = "max"
)

type Heap struct {
	heap []int
	heapType heapType
}

func NewHeap(heapType heapType) (*Heap, error) {
	if heapType != MIN && heapType != MAX {
		return nil, fmt.Errorf("Please provide valid heap type: min or max")
	}

	heapArray := make([]int, 0)
	return &Heap{
		heap: heapArray,
		heapType: heapType,
	}, nil
}

// Add elem into heap
func (h *Heap) Add (num int) {
	h.heap = append(h.heap, num)
	index := len(h.heap)-1
	switch h.heapType {
		case MIN:
			for index > 0 && h.heap[(index-1)/2] > h.heap[index] {
				// swap with parent
				h.heap[(index-1)/2], h.heap[index] = h.heap[index], h.heap[(index-1)/2]
				index = (index-1)/2
			}
		case MAX:
			for index > 0 && h.heap[(index-1)/2] < h.heap[index] {
				// swap with parent
				h.heap[(index-1)/2], h.heap[index] = h.heap[index], h.heap[(index-1)/2]
				index = (index-1)/2
			}
	}
	slog.Info("elem added into heap", slog.Int("element", num))
	return
}

// Pop elem from heap
func (h *Heap) Pop() (answer int) {
	answer = h.heap[0]
	lastInd := len(h.heap)-1
	h.heap[0] = h.heap[lastInd]
	h.heap = h.heap[:lastInd]
	
	index := 0
	childIndex1, childIndex2 := 1, 2
	switch h.heapType {
	case MIN:
		for childIndex1 < lastInd &&  (h.heap[index] > h.heap[childIndex1] || h.heap[index] > h.heap[childIndex2]) {
			if h.heap[index] > h.heap[childIndex1] {
				h.heap[childIndex1], h.heap[index] = h.heap[index], h.heap[childIndex1]
				index = childIndex1
			} else {
				h.heap[childIndex2], h.heap[index] = h.heap[index], h.heap[childIndex2]
				index = childIndex2
			}
			childIndex1, childIndex2 = 2*index+1, 2*index+2
		}
	case MAX:
		for childIndex1 < lastInd &&  (h.heap[index] < h.heap[childIndex1] || h.heap[index] < h.heap[childIndex2]) {
			if h.heap[index] < h.heap[childIndex1] {
				h.heap[childIndex1], h.heap[index] = h.heap[index], h.heap[childIndex1]
				index = childIndex1
			} else {
				h.heap[childIndex2], h.heap[index] = h.heap[index], h.heap[childIndex2]
				index = childIndex2
			}
		}
		childIndex1, childIndex2 = 2*index+1, 2*index+2
	}
	return answer
}

func (h Heap) GetHeap() []int {
	if len(h.heap) > 0 {
		return h.heap
	}
	return []int{}
}