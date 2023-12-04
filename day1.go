package main

import ("fmt"
"io/ioutil")

func main() {
	fileBytes, err := ioutil.ReadFile("day1_input")
	if err != nil {
		fmt.Print(err)
	}
	fileString := string(fileBytes)

	fmt.Print(fileString)
}