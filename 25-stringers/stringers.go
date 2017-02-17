package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"A", 42}
	b := Person{"B", 43}

	fmt.Println(a.String())

	fmt.Println(a, b)
}
