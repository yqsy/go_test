package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {

	v := Vertex{3, 4}
	v.X = 4
	v.Y = 5

	p := &v
	p.X = 1e9

	fmt.Println(v)

	fmt.Println(v1, v2, v3, p)

}
