package main

import (
	"encoding/json"
	"fmt"
)

//import (
//	"fmt"
//	"strings"
//)

//func main() {
//	var num float64
//	fmt.Println("Enter the number : ")
//	fmt.Scanf("%f",&num)
//	result := strconv.FormatFloat(num, 'f', 0,64)
//	fmt.Println(result)
//}
//package main
//
//import (
//"fmt"
//"strings"
//)

func main() {
	var name string
	var address string

	personMap := make(map[string]string)

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter your address: ")
	fmt.Scan(&address)

	personMap[name] = address

	jsonObject, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(jsonObject))
	}
}