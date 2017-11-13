package main

import "fmt"

type Vt struct {
	X,Y int
}

// 直接操作值
func (v *Vt) plus() int {
	v.X = 10
	return v.X + v.Y
}

// 拷贝值
func (v Vt) mo() int {
	v.Y = 100
	return v.X + v.Y
}

// 可以向本包内的任何对象赋值，基础对象除外
type MyFloat float64

func (f MyFloat) plus() float64 {
	f += 1
	return float64(f)
}

func (f *MyFloat) plusP() float64 {
	*f += 1
	return float64(*f)
}



func main() {
	v1 := Vt{
		1,2,
	}
	fmt.Println("plus:", v1, v1.plus(), v1)
	fmt.Println("plus:", v1, v1.mo(), v1)

	var mf MyFloat = 10.00
	fmt.Println(mf, mf.plus(), mf)
	fmt.Println(mf, mf.plusP(), mf)
}