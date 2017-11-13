package main

import "fmt"

type FirstStruct struct {
	X int
	Y int
}

func main() {
	st := FirstStruct{1,2}
	fmt.Println(st)

	fmt.Println(st.X)

	// 结构体指针
	sp := &st
	fmt.Println(sp.Y)

	var (
		v1 = FirstStruct{2,3}
		v2 = FirstStruct{X:4}
		v3 = FirstStruct{Y:5}
		p = &FirstStruct{6,7}
	)

	fmt.Println(v1, v2, v3, p)

	i := 43
	p2 := &i
	fmt.Println(p2)
}