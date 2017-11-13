package main

import "fmt"

type Phone interface {
	call() string
	sms() int
}

type Nokia struct {
	name string 
}

func (no *Nokia) call() string {
	return no.name + " is calling you"
}

func (no *Nokia) sms() int {
	return 0
}

func main() {
	var p Phone = &Nokia{"iphone"}
	fmt.Println(p.call())
}
