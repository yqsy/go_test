package main

import (
	"fmt"
	"time"
)

func GetValue(i chan int) {
	time.Sleep(time.Second * 1)
	i <- 1
}

func main() {
	ch := make(chan int, 2)
	go GetValue(ch)
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
