package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)

	for i := range tick {
		fmt.Println(i)
	}
}
