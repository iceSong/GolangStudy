package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个goroutine能访问c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v : make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}