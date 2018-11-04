package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	var fileNameInput string
	var nameStructSlice = make([]Name, 0)

	fmt.Printf("Provide the name of the file with names:\n")
	fmt.Scan(&fileNameInput)
	file, _ := os.Open(fileNameInput)
	reader := bufio.NewReader(file)
	line, error := Readln(reader)
	for error == nil {
		// fmt.Println(line) // zamiast tego tworzyć structa i dodawać do slice'a
		nameSlice := strings.Split(line, " ")
		fullName := Name{fname: nameSlice[0], lname: nameSlice[1]}
		fmt.Println(fullName.fname)
		fmt.Println(fullName.lname)
		nameStructSlice = append(nameStructSlice, fullName)
		line, error = Readln(reader)
	}
	fmt.Println(nameStructSlice)
	for _, name := range nameStructSlice {
		fmt.Println(name.fname + " " + name.lname)
	}
}

type Name struct {
	fname string
	lname string
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
