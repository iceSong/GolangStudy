package main

import "fmt"

func fibonacci() func() int{

	b1, b2 := 0, 1

	return func() int {
		tmp := b1
		b1, b2 = b2, (b1 + b2)
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	f2 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f2())
	}
}