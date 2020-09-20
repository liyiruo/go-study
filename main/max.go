package main

import "fmt"

func main() {
	var a int = 200
	var b int = 200
	var res int
	res = mix(a, b)

	fmt.Println(res)

}

func mix(num1, num2 int) int {
	var result int
	if num2 < num1 {
		result = num1
	} else if num2 > num1 {
		result = num2
	}
	return result
}
