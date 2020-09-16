package main

import (
	"fmt"
	"sync"
)

var once sync.Once = sync.Once{}
var instanse *singler

type singler struct {
	Name string
}

func NewSingler() *singler {
	once.Do(func() {
		instanse = new(singler)
		instanse.Name = "这是一个单例"
	})
	return instanse
}

func main() {
	singler := NewSingler()
	fmt.Println(singler.Name)
}
