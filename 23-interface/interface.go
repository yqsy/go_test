package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	var c Abser = &Vertex{4, 3}

	fmt.Println(c.Abs())
	//a = v

	fmt.Println(a.Abs())

	var human Human = Man{}
	human.Speak()

	human = &Woman{}
	human.Speak()

	var u interface{}

	u = 1
	describe(u)

	u = "1"
	describe(u)
}

type Human interface {
	Speak()
}

type Man struct {
}

type Woman struct {
}

func (man Man) Speak() {
	fmt.Println("I'm a man")
}
func (woman *Woman) Speak() {
	fmt.Println("I'm a woman")
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
