package main

import "fmt"

type T1 struct {
}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// x := i.(float32)
	// fmt.Println(x)

	var str1 interface{} = T1{}

	switch str1.(type) {
	case T1:
		fmt.Println("T1")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
}
