package main

import (
	"fmt"
)

type Heap []int

func (heap *Heap) insert(v int) {

	*heap = append(*heap, v)

	// if is not root
	if len(*heap) > 1 {
		//recover heap property
		heap.percolate_up()
	}

}

//swap value

func (heap Heap) swap(i, j int) {
	tmp := heap[i]
	heap[i] = heap[j]
	heap[j] = tmp
}

//recover heap property
func (heap Heap) percolate_up() {
	lightIdx := len(heap)
	parentIdx := lightIdx / 2

	for parentIdx > 0 && (heap[lightIdx-1] < heap[parentIdx-1]) {
		heap.swap(lightIdx-1, parentIdx-1)
		lightIdx = parentIdx
		parentIdx = (lightIdx / 2)
	}
}

func (heap Heap) percolate_down() {
	minIdx := 0
	heavyIdx := 1
	sign := 1
	// sign == 0 loop, sign == 1 break
	for sign == 1 {
		child1Idx := heavyIdx * 2
		child2Idx := child1Idx + 1
		sign = 0
		if len(heap) < child1Idx { //child is none
			break
		} else if len(heap) < child2Idx { // child2Idx is none
			minIdx = child1Idx
		} else {
			if heap[child1Idx-1] < heap[child2Idx-1] {
				minIdx = child1Idx
			} else {
				minIdx = child2Idx
			}
		}

		if heap[heavyIdx-1] > heap[minIdx-1] {
			heap.swap(heavyIdx-1, minIdx-1)
			heavyIdx = minIdx
			sign = 1
		}
	}
}

func (heap *Heap) delete_min() int {
	if len(*heap) <= 0 {
		fmt.Println("Error: This heap is empty")
		return -1
	}

	min := (*heap)[0]
	heap.swap(0, len(*heap)-1)
	(*heap) = (*heap)[:len(*heap)-1]
	heap.percolate_down()
	return min
}

func main() {
	var heap Heap
	heap.insert(9)
	heap.insert(4)
	heap.insert(7)
	heap.insert(2)
	heap.insert(5)
	heap.insert(8)
	heap.insert(3)

	fmt.Println(heap)
	fmt.Println(heap.delete_min())
	fmt.Println(heap)
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())
	fmt.Println(heap.delete_min())

}
