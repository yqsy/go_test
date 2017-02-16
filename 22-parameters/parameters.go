package main

import "fmt"

func fuck(x *int) {
	*x = 1
}

type ff struct {
	X float64
}

func fuckff(f *ff) {
	f.X = 1.1
}

func (f *ff) fuckff() {
	f.X = 1.22
}

func main() {

	var a int

	fuck(&a)
	fmt.Println(a)

	f := ff{1.1}

	fuckff(&f)
	fmt.Println(f.X)

	f2 := ff{1.2}
	f2.fuckff()
	fmt.Println(f2.X)
}
