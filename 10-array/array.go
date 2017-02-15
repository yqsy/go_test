package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[0:1]
	fmt.Println(s)

	names := []string{
		"John",
		"Paul",
		"Geroge",
		"Ringo",
	}

	// fmt.Println(names)

	x := names[0:2]
	y := names[2:4]

	x[0] = "123"

	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(names)
}
