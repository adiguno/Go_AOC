package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	// "io/ioutil"
	// "rsc.io/quote"
)

func main() {
	// read entire file as a string
	// fileBytes, err := ioutil.ReadFile("day1_input")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fileString := string(fileBytes)

	file, err := os.Open("day1_input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner((file))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Print(scanner.Err())
	}

	// for i, line := range lines {
	// 	// fmt.Println(i, line)
	// 	fmt.Println(i)
	// 	getDigits(line)
	// }
	var total int
	total += getSum(getDigits("1abc2"))
	a := getSum(getDigits("1abc2"))
	if a == 12 {
		fmt.Println("yes")
	} else if a == 540 {
		fmt.Println("bad")
	}
	// fmt.Println("First Integer:", 12)

	// fmt.Println(string(rune((total))))
	fmt.Println("total", total)
	fmt.Printf("Type of total: %T\n", total)

}

func getSum(num1 int, num2 int) int {
	return num1*10 + num2
}

func getDigits(text string) (int, int) {
	// fmt.Println(text)
	// for _, char := range text {
	// 	// fmt.Println(char) // outputs unicode e.g. 116, 108
	// 	fmt.Printf("%c\n", char)
	// }
	var first, last int
	var firstFound, lastFound bool
	i := 0

	// cases
	// no number
	// single number
	// 2+ numbers
	for i < len(text) {
		char := text[i] //byte
		if unicode.IsDigit(rune(char)) {
			if !firstFound {
				first = int(char)
				firstRune := rune(char)
				fmt.Println("char", char)
				fmt.Println("firstRune", firstRune)
				fmt.Println("firstRune int", int(firstRune))
				fmt.Println("firstRune int string", string(int(firstRune)))
				// fmt.Println(string(char))
				// fmt.Printf("first: %d\n", first)
				firstFound = true
			} else if firstFound && !lastFound {
				last = int(char)
				lastFound = true
			} else if lastFound {
				last = int(char)
				lastFound = true
			}
		}
		i++
	}

	// fmt.Println("first: " + string(rune(first)))
	// fmt.Println("last: " + string(rune(last)))
	fmt.Println("first: ", first)
	fmt.Println("last: ", last)
	fmt.Printf("type of last: %T\n", last)
	if last == 2 {
		fmt.Println("okayyyy ")

	}
	return first, last
}
