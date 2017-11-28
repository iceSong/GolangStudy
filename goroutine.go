package main

import (
	"fmt"
	"time"
	"strconv"
)

var tmp = 0

func say(s string) {
	for i:=0;i<5;i++ {
		time.Sleep(100 * time.Millisecond)
		tmp++
		fmt.Println(s, strconv.Itoa(tmp))
	}
}

func main() {
	go say("go")
	say("ma")
}