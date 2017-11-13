package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int{
	a := strings.Fields(s)
	fmt.Println(a)
	m := make(map[string]int)
	for _, v := range a {
		m[v] = m[v] + 1
	} 
	return m
}

func main() {
	m := WordCount("hi i love you i love you and me")
	fmt.Println(m)
}