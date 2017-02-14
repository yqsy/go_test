package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false //
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) //
)

func main() {
	var str1 string
	str1 = "hello world"
	str2 := "hello world2"
	fmt.Println(str1, str2)

	fmt.Printf("Type: %T value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v \n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
