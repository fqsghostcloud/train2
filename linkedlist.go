package main

import (
	"fmt"
)

// define node struct
type Node struct {
	prev, next *Node
	value      interface{}
}

// define linked list struct
type List struct {
	head *Node
	tail *Node
	len  int
}

// create list
func createList() List {
	l := List{}
	fmt.Println("Create linkedlist Success!")
	return l
}

// get lengh
func (l *List) lengh() int {
	return l.len
}

//insert  node by index
func (l *List) insertByIndex(index int, n *Node) {
	p := l.head.next
	for ; p != nil && index > 1; p = p.next {
		index--
	}

	n.prev = p.prev
	p.prev = n
	n.next = n.prev.next
	n.prev.next = n
	l.len++
	fmt.Println("Insert node success!")
}

// append node at tail
func (l *List) append(v interface{}) {
	if l.tail == nil && l.head == nil {
		l.head = new(Node)
		l.tail = l.head
	}
	l.tail.next = new(Node)
	l.tail.next.prev = l.tail
	l.tail = l.tail.next
	l.tail.value = v
	l.len++
	fmt.Printf("Push value: %v success!\n", v)
}

// pop node at tail
func (l *List) pop() {
	fmt.Printf("Pop value : %v success!\n", l.tail.value)
	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
}

//remove node by index
func (l *List) removeByIndex(index int) {
	p := l.head.next
	fmt.Printf("Remove index is: %d\n", index)
	if l.len < index || index <= 0 {
		fmt.Println("Index out of range!")
		return
	} else if index == 1 {
		l.head.next = p.next
		p.next.prev = l.head
		l.len--
		fmt.Println("Remove first node success!")
	} else if index == l.len {
		l.pop()
	} else {
		for ; p != nil && index >= 1; p = p.next {
			if index == 1 {
				p.prev.next = p.next
				p.next.prev = p.prev
				p.next = nil
				p.prev = nil
			}
			index--
		}
		l.len--
	}

}

//contains value
func (l *List) contains(v interface{}) int {
	count := 0
	if v == nil {
		fmt.Println("Your input is nil!")
		return 0
	}
	for p := l.head.next; p != nil; p = p.next {
		if p.value == v {
			count++
		}
	}

	if count > 0 {
		fmt.Printf("This linkedlist contains: %v %d items\n", v, count)
	} else {
		fmt.Printf("This linkedlist not contains: %v\n", v)
	}

	return count

}

//update linkedlist value by index
func (l *List) updateValueByIndex(index int, v interface{}) {
	p := l.head.next
	if l.len < index || index <= 0 {
		fmt.Println("Index out of range!")
		return
	}

	for ; p != nil && index >= 1; p = p.next {
		if index == 1 {
			fmt.Printf("Update value : %v to %v \n", p.value, v)
			p.value = v
		}
		index--
	}

}

//is empty
func (l *List) isEmpty() bool {
	return l.len == 0
}

//print node info
func (l *List) printNodeInfo(at *Node) {
	if at != nil {
		fmt.Printf("Node Value is: %v\n", at.value)
	} else {
		fmt.Println("Node is nil")
	}
}

// print linkedlist info
func (l *List) printListInfo() {
	if l == nil {
		fmt.Println("List is nil!")
	}
	for p := l.head.next; p != nil; p = p.next {
		l.printNodeInfo(p)
	}
	fmt.Printf("Node len is: %d", l.len)
}

func main() {
	// create List
	linkedList := createList()
	fmt.Printf("\n----------------------------------\n")
	//append value
	linkedList.append("a")
	linkedList.append("b")
	linkedList.append("c")
	linkedList.printListInfo()
	fmt.Printf("\n----------------------------------\n")
	// insert node by index
	newNode := new(Node)
	newNode.value = "k"
	linkedList.insertByIndex(2, newNode)
	linkedList.printListInfo()
	fmt.Printf("\n----------------------------------\n")
	//pop node
	linkedList.pop()
	linkedList.printListInfo()
	fmt.Printf("\n----------------------------------\n")

	//remove node by index
	linkedList.removeByIndex(2)
	linkedList.printListInfo()
	fmt.Printf("\n----------------------------------\n")

	//contains value
	linkedList.contains("b")
	fmt.Printf("\n----------------------------------\n")

	//update value
	linkedList.updateValueByIndex(2, "x")
	linkedList.printListInfo()
	fmt.Printf("\n----------------------------------\n")

	//is empty
	fmt.Printf("This linkedlist is empty: %v", linkedList.isEmpty())
	fmt.Printf("\n----------------------------------\n")

	//get linedlist length
	fmt.Printf("This linkedlist length is: %d", linkedList.lengh())
	fmt.Printf("\n----------------------------------\n")

}
