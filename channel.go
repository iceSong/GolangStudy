package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum  // 将结果送入管道
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)  // 关闭channel
}

func main() {
	a := []int{7, 2, 4, 6, -1}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从管道提取
	fmt.Println(x, y, x + y)

	// 建立带缓冲长度的channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch)
	ch <- 3
	v, ok := <- ch
	fmt.Println(v, ok)

	// 通道的range和close
	cha := make(chan int, 10)
	go fibonacci(cap(cha), cha)
	for i := range cha {
		fmt.Println(i)
	}

}

