package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum  // 将结果送入管道
}

func main() {
	a := []int{7, 2, 4, 6, -1}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	 x, y := <-c, <-c // 从管道提取

	 fmt.Println(x, y, x + y)
}