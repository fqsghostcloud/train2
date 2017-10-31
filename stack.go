package main

import (
	"fmt"
)

//define stack struct
type Stack []interface{}

//get stack length
func (stack Stack) len() int {
	return len(stack)
}

//get stack capacity
func (stack Stack) cap() int {
	return cap(stack)
}

//push value
func (stack *Stack) push(v interface{}) {
	*stack = append(*stack, v)
}

//get top value not remove
func (stack Stack) top() interface{} {
	if len(stack) == 0 {
		fmt.Println("This stack cap is 0!")
		return nil
	}
	return stack[len(stack)-1]
}

//get pop value and remove
func (stack *Stack) pop() interface{} {
	tempStack := *stack
	if len(tempStack) == 0 {
		fmt.Println("This stack cap is 0!")
		return nil
	}

	value := tempStack[len(tempStack)-1]
	*stack = tempStack[:len(tempStack)-1]
	return value
}

//stack is empty
func (stack Stack) isEmpty() bool {
	return len(stack) == 0
}

//print stack info
func (stack Stack) printInfo() {
	if stack.isEmpty() {
		fmt.Println("This Stack is empty!")
		return
	}

	for i := 0; i < stack.len(); i++ {
		fmt.Printf("Stack value is: %v\n", stack[i])
	}
}

func main() {
	var testStack Stack
	testStack.push("a") //push data
	testStack.push("b")
	testStack.push("c")

	testStack.printInfo()
	fmt.Printf("\n-------------info-show-----------------------\n")
	fmt.Printf("Stack len is: %d\n", testStack.len())
	fmt.Printf("Stack cap is: %d\n", testStack.cap())
	fmt.Printf("Stack is empty: %v", testStack.isEmpty())

	fmt.Printf("\n\n------------------top-test-----------------\n")
	fmt.Printf("Stack top is: %v\n", testStack.top())
	fmt.Printf("Stack len is: %d\n", testStack.len())

	fmt.Printf("\n\n-----------------pop-test-------------------\n")
	fmt.Printf("Stack pop is: %v\n", testStack.pop())
	fmt.Printf("Stack len is: %d\n", testStack.len())
	fmt.Printf("Stack pop is: %v\n", testStack.pop())
	fmt.Printf("Stack len is: %d\n", testStack.len())

}
