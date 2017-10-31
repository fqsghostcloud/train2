package main

import (
	"fmt"
	"sync"
)

var instance *singleton
var once sync.Once

//define singleton struct
type singleton struct {
	name string
}

//get singleton instance
func getInstance() *singleton {
	if instance != nil {
		return instance
	}
	once.Do(
		func() {
			instance = &singleton{"hello"}
		})
	return instance
}

func main() {
	instance = getInstance()
	fmt.Printf("singleton mothod name is: %s\n", instance.name)
}
