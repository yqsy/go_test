package main

import "fmt"

// 0 1 2 3 5 8 13 21 34 55

func calcfib(n int) int {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	n1, n2, n3 := 0, 1, 0

	for i := 0; i < n-2; i++ {
		n3 = n1 + n2
		n1 = n2
		n2 = n3
	}

	return n3
}

func fibonacci() func() int {
	x, y := 1, 0
	return func() int {
		z := x + y
		x = y
		y = z
		return x
	}
}

func main() {

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println(calcfib(10))

}
