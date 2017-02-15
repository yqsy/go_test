package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}

	fmt.Println(s)

	printSlice(q)

	var x []int
	fmt.Println(x, len(x), cap(x))
	if x == nil {
		fmt.Println("nil!")
	}

	b := make([]int, 0, 5)
	fmt.Printf("%s len=%d cap=%d %v\n",
		"b", len(b), cap(b), b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
