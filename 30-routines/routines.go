package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func xxxx() {
	for i := 0; i < 110; i++ {
		time.Sleep(time.Second * 5)
		fmt.Println("fuck u")
	}
}

func main() {
	go xxxx()
	go say("world")
	say("hello")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
