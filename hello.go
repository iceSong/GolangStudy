package main

import (
	"fmt"
	"runtime"
)

var name, gender = "song", "men"

func main() {
	var age int
	age = 100
	fmt.Println(name, gender)
	fmt.Println(age)
	fmt.Println("hello world")
	fmt.Println(getOs())
}

func getOs() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		return os
	}
}
