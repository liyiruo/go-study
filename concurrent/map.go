package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["1"] = "yi"
	countryCapitalMap["2"] = "er"
	countryCapitalMap["3"] = "san"
	countryCapitalMap["4"] = "si"

	for k, v := range countryCapitalMap {
		fmt.Println(k, "==", v)
	}

	//判断某个key是否存在
	s := "3"
	capital, ok := countryCapitalMap[s]

	if ok {
		fmt.Println(s, "==", capital)
	} else {
		fmt.Println(s, "==不存在")
	}

}
