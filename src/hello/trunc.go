package main
import "fmt"
import "strconv"

func main() {
     var floatInput float32
     fmt.Printf("Specify a decimal number:\n")
     fmt.Scan(&floatInput)
     formattedNumber := strconv.FormatFloat(float64(floatInput), 'f', 0, 32)
     fmt.Printf("Number specified: %s", formattedNumber)
}