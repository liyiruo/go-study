package main

import "fmt"

func main() {
	myMap := map[string]string{
		"1": "一",
		"2": "二",
		"3": "三",
	}
	fmt.Println("原始地图")

	for k := range myMap {
		fmt.Println(k, "首都是", myMap[k])
	}

	delete(myMap, "3")
	fmt.Println("=======删除3=======")

	for k := range myMap {
		fmt.Println(k, "首都是", myMap[k])
	}

}
