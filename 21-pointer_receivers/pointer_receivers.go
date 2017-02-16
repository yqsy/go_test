package main

import "fmt"

type MyInt struct {
	a int
}

// * type, can modify value
func (m *MyInt) Modify(v int) {
	m.a = v
}

func main() {

	var a MyInt

	a.Modify(1)

	fmt.Println(a.a)
}
