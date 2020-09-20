package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hello")
	say("world")
}
func say(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
