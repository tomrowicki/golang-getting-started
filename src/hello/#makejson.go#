package main

import "fmt"
import "encoding/json"

func main() {
	var name string
	var address string
	var myMap map[string]string
	
	myMap = make(map[string]string)

	fmt.Printf("State your name:\n")
	fmt.Scan(&name)
	fmt.Printf("State your address:\n")
	fmt.Scan(&address)

	myMap["name"] = name
	myMap["address"] = address
	barr, _ := json.Marshal(myMap)
	fmt.Println(string(barr))
}
