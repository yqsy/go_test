package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	foo()
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	x, y := split(17)
	fmt.Println(x, y)

	var i int

	i = i + 1

	fmt.Println(i, c, python, java)
}
