package heap

import (
	"errors"
	"fmt"
)


type Heap struct {
	Data []int
	length int
	size int
}

func (h *Heap) Init(length int) {
	h.Data = make([]int, length + 1)
	h.length = length + 1
	h.size = 0
}
func (h Heap) Len() int {
	return h.size
}

func (h *Heap) Add(num int) error {
	if h.size >= h.length {
		return errors.New("out of range")
	}

	lastIndex := h.size + 1
	h.size = lastIndex
	h.Data[lastIndex] = num

	h.reBalanceUp(lastIndex)
	return nil
}

func (h *Heap) reBalanceUp(i int) {
	if i == 1 {
		return
	}
	p := h.getParent(i)
	if p > h.Data[i] {
		h.swap(i, i/2)
		h.reBalanceUp(i / 2)
	}
}

func (h *Heap) reBalanceDown(i int) {
	if i >= h.size {
		return
	}
	leftIndex := 2 * i
	rightIndex := leftIndex + 1
	if leftIndex <= h.size && rightIndex <= h.size {
		var min int
         if h.Data[i] < h.Data[leftIndex] {
         	min = i
		 } else {
		 	min = leftIndex
		 }

		if h.Data[min] > h.Data[rightIndex] {
			min = rightIndex
		}
		if min != i {
			h.swap(i, min)
			h.reBalanceDown(min)
		}
		return
	}
	if leftIndex <= h.size && rightIndex > h.size {
		if h.Data[i] > h.Data[leftIndex] {
			h.swap(i, leftIndex)
			h.reBalanceDown(leftIndex)
		}
	}
}

func (h Heap) getParent(i int) int{
	return h.Data[i/2]
}

func (h Heap) getBro(i int) (int, error) {
	if i <= 1 {
		return 0, nil
	}
	if i % 2 == 1 {
		if i + 1 >= h.size {
			return 0, errors.New("out of range")
		}
		return i + 1, nil
	} else {
		return i - 1, nil
	}
}


func (h *Heap) swap(i ,j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h *Heap) Poll() int {
	res := h.Data[1]
	h.delete()
	return res
}

func (h Heap) Peek() int {
	return h.Data[1]
}

func (h *Heap) delete() {
	h.Data[1] = h.Data[h.size]
	h.Data[h.size] = 0
	h.size--
	h.reBalanceDown(1)
}

func (h Heap) PrintHeap() {
	for _, n := range h.Data {
		fmt.Print(n, " ")
	}
	fmt.Println("")
}
