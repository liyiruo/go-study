package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)
}
