package main

import (
	"fmt"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	m = map[string]Vertex{
		"Bell Labs": Vertex{
			1.1, 1.1,
		},
		"Google": Vertex{
			2.2, 2.2,
		},
	}

	m["Microsoft"] = Vertex{3.3, 3.3}

	delete(m, "Microsoft")

	fmt.Println(m)

	v, ok := m["Google"]
	fmt.Println("The value:", v, "Present?", ok)

	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}
