package main

import (
	"fmt"
)

func main() {
	fmt.Println(sd())
}

func sd() int {
	i := 0
	defer func(){i++}()
	return i
}
