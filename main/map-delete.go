package main

import "fmt"

func main() {

	/*创建map*/
	countryCapitalMap := map[string]string{
		"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	//打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	//删除意大利
	delete(countryCapitalMap, "Italy")

	fmt.Println("删除意大利")

	fmt.Println("删除意大利")
	//再次打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
