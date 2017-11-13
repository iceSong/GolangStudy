package main

import "fmt"

type MyReader struct {}

func (r *MyReader) Read(b []byte) (int, error) {
	b[0]='A'
	return 1, nil
}

func main() {
	reader := &MyReader{}
	a := make([]byte, 5, 5)
	fmt.Println(reader.Read(a))
	fmt.Println(a)
}

