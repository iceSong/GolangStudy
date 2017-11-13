package main

import (
 "fmt"
 "math"
)

func main() {
	/**
	* 函数也是值
	*/
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(12,34))

	pos, neg := adder(), adder()
	for i:=0; i<10; i++{
		fmt.Println(
			pos(i),
			neg(-2*2),
		)
	}


	f1, f2 := adder(),adder()
	fmt.Println(f1(1), f2(5))
}

/**
* 函数闭包
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}