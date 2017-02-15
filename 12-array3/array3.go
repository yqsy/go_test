package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4}

	c := a[0:]
	c[0] = 100

	// 虽然 len(d) = 0, 但是它的空间为5
	d := make([]int, 0, 5)

	e := d[:2]
	f := d[2:5]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(len(d), cap(d))
	fmt.Println(e)
	fmt.Println(f)
}
