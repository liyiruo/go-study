package main

import "fmt"

func main() {

	var str1 string = "\\\""

	fmt.Println(str1)

	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers2[i]
	}

	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))

}
