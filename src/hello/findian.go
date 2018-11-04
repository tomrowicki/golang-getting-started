package main

import "fmt"
import "strings"

func main() {
	//var providedString string
	first, second, third, fourth := "ian", "Ian", "iuiygaygn", "I d skd a efju N"
	const wrongFirst, wrongSecond, wrongThird = "ihhhhhn", "ina", "xian"
	fmt.Printf("Should return four trues.\n")
	fmt.Printf("%b\n", check(first))
	fmt.Printf("%b\n", check(second))
	fmt.Printf("%b\n", check(third))
	fmt.Printf("%b\n", check(fourth))

	fmt.Printf("Should return three falses.\n")
	fmt.Printf("%b\n", check(wrongFirst))
	fmt.Printf("%b\n", check(wrongSecond))
	fmt.Printf("%b\n", check(wrongThird))

	//fmt.Printf("Provide a string:")
	//fmt.Printf("Lowercased String: %s", strings.ToLower(providedString))
	//fmt.Scan(&providedString)
}

func check(str string) bool {
	s := strings.ToLower(str)
	return strings.Contains(s, "a") && strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n")
}
