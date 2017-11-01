package main

import (
	"fmt"
)

//define node store key and value
type Node struct {
	key   int
	value interface{}
	next  *Node
}

//define hashMap
type HashMap []Node

const cap int = 11

//init hashMap
func createHashMap() HashMap {
	hashMap := make(HashMap, cap)
	return hashMap
}

//define hash mothod
func (hashMap HashMap) hash(k int) int {
	if k < 0 {
		return -1
	}
	return k % cap
}

//put key value to hashmap
func (hashMap *HashMap) put(k int, v interface{}) {
	if hashMap.lookUp(k) != nil {
		fmt.Printf("Your key: %d already exist!\n", k)
		return
	}
	index := hashMap.hash(k)
	if index < 0 {
		fmt.Println("key must >= 0")
		return
	}
	if (*hashMap)[index].value == nil {
		(*hashMap)[index].key = k
		(*hashMap)[index].value = v
	} else {
		p := &(*hashMap)[index]
		for ; p != nil; p = p.next {
			if p.next == nil {
				p.next = new(Node)
				p.next.key = k
				p.next.value = v
				break
			}
		}
	}
}

//get node from key
func (hashMap HashMap) lookUp(k int) *Node {
	index := hashMap.hash(k)
	if index < 0 {
		fmt.Println("key must >= 0")
		return nil
	}
	for p := &hashMap[index]; p != nil; p = p.next {
		if p.key == k {
			return p
		}
	}
	return nil
}

//whether key is exist
func (hashMap HashMap) exist(k int) bool {
	return hashMap.lookUp(k) != nil
}

//from key get value
func (hashMap HashMap) get(k int) interface{} {
	if !hashMap.exist(k) {
		fmt.Printf("This key: %d is not exist!\n", k)
	}
	return hashMap.lookUp(k).value
}

//delete node
func (hashMap *HashMap) delete(k int) {
	if !hashMap.exist(k) {
		fmt.Printf("This key: %d is not exist!\n", k)
		return
	}
	p := hashMap.lookUp(k)
	p.key = 0
	p.value = nil
}

// print hashmap info
func (hashMap HashMap) prinfInfo() {
	for i := 0; i < len(hashMap)-1; i++ {
		fmt.Printf("--------------------------\n")
		fmt.Printf("Index is: %d\n", i)
		for p := &hashMap[i]; p != nil; p = p.next {
			fmt.Printf("Node key: %d  value: %v\n", p.key, p.value)
		}
		fmt.Printf("--------------------------\n")
	}
}

func main() {
	hashMap := createHashMap()
	hashMap.put(1, "a")
	hashMap.put(1, "b")
	hashMap.put(12, "c")
	hashMap.put(6, "gg")
	hashMap.prinfInfo()
	fmt.Println(hashMap.exist(12))
	hashMap.delete(12)
	fmt.Println(hashMap.exist(12))
	fmt.Printf("key is %d  value is: %v", 6, hashMap.get(6))

}
