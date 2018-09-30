package main

import "fmt"
import "strconv"
import "sort"

func main() {
	var input string
	var sli = make([]int, 3)
	fmt.Printf("Provide an integer or x to quit:\n")
	for {
		fmt.Scan(&input)
		if input == "x" {
			break
		}
		parsedInt64, err := strconv.ParseInt(input, 10, 32)
		parsedInt32 := int(parsedInt64)
		if err != nil {
			fmt.Printf("Error has occured. %s", err)
		}
		sli = append(sli, parsedInt32)
		sort.Ints(sli)
		fmt.Printf("The sorted slice: %v", sli)
	}
}
