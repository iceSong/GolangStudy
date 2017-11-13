package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967, // 注意逗号，go的map、slice不再一行时需要特别注意
	}
	fmt.Println(m["Bell Labs"])

	m2 := map[string]Vertex{
		"hello":Vertex{
			29.00,39,
		},
		"dog":Vertex{
			12,13,
		},
	}
	fmt.Println(m2)

	fmt.Println(m2["hello"])

	// 省略类型名称
	m3 := map[string]Vertex{
		"enha":{1,2},
		"gege":{3,4},
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	m4[12] = "age"
	m4[13] = "name"
	m4[15] = "space"
	fmt.Println(m4)
	m4[12] = "age++"
	delete(m4, 13)
	fmt.Println(m4)
	v,ok := m4[15]
	fmt.Println("The value:", v, "Present?", ok)
}