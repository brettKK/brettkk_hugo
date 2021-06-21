package data

import "fmt"

// max heap
// usage:  1 heap sort ;   2 priority queue
type heap struct {
	Data []int
	N int
}

func NewHeap() heap {
	return heap{
		Data: make([]int, 0),
		N: 0,
	}
}

func (h *heap) Add(elem int) {
	h.N++
	if len(h.Data) == 0 {
		h.Data =  append(h.Data, 0, elem)
		return
	} else {
		h.Data = append(h.Data, elem)
		// Upper Change heap
		h.UpperChange(len(h.Data)-1)
	}
}

func (h *heap) DelMax() int{
	if len(h.Data) < 1 {
		return -1
	} else {
		max := h.Data[1]
		h.Data[1] = h.Data[len(h.Data) -1]
		h.N--
		h.Data = h.Data[:len(h.Data) - 1]
		//down change
		h.DownChange(1)
		return max
	}
	return -1
}


func (h *heap) DownChange(index int) {
	for h.getLeftChld(index) <= h.N &&h.getLeftChld(index) >= 1 && h.getRightChld(index) <= h.N && h.getRightChld(index) >= 1{
		cur := h.getLeftChld(index)
		if h.Data[h.getRightChld(index)] > h.Data[cur] {
			cur = h.getRightChld(index)
		}
		h.swap(index, cur)
		index = cur
	}
}

func (h *heap) UpperChange(index int) {
	for index > 1 {
		parent := h.getParent(index)
		if h.Data[parent] > h.Data[index] {
			break
		} else {
			h.swap(parent, index)
			index = parent
		}

	}
}

func (h *heap) swap(a, b int) {
	tmp := h.Data[a]
	h.Data[a] = h.Data[b]
	h.Data[b] = tmp
}

//func (h *heap) less(a, b int) bool {
//	if h.Data[a] < h.Data[b] {
//		return true
//	}
//	return false
//}

func (h *heap) getParent(chld int) int {
	return chld /2
}

func (h *heap) getLeftChld(parent int) int {
	if h.N < parent * 2 {
		return -1
	}
	return parent * 2
}

func (h *heap) getRightChld(parent int) int {
	if h.N < parent * 2 + 1 {
		return -1
	}
	return parent * 2 + 1
}

func (h *heap) PrintHeap() {
	fmt.Printf("===print heap===size=%d\n", h.N)
}