package main
import "fmt"

func main() {
	p := []int{2,3,5,7,11,13}
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	// 切片, 省略下下标从0开始	
	fmt.Println("p[:3] == ", p[:3])
	// 省略上标到结束
	fmt.Println("p[4:]", p[4:])

	// 使用make函数创建slice
	a := make([]int, 5)
	showSlice("a", a)
	b := make([]int, 0, 5)
	showSlice("b", b)
	c := b[:2]
	showSlice("c", c)
	d := c[1:5]
	showSlice("d", d)

	// 空slice nil
	var z []int
	showSlice("z", z)

	// 向slice添加元素
	z = append(z, 1)
	showSlice("z", z) 
	z = append(z, 2,3,4,5)
	showSlice("z", z)
	z = append(z, 0)

	// for range
	for i,v := range z {
		fmt.Printf("z[%d] = %d\n", i,v)
	}

}


func showSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	if nil == x {
		fmt.Printf("%s == nil\n", s)
	}
}