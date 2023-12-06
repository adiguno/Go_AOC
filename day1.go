package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// "io/ioutil"
	// "rsc.io/quote"
)

func main() {
	// mainStrings()

	// read entire file as a string
	// fileBytes, err := ioutil.ReadFile("day1_input")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fileString := string(fileBytes)

	// read file
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

	var total int
	for index, line := range lines {
		// fmt.Println(i, line)
		lineSum := getSum(getDigits(line))
		total += lineSum
		fmt.Printf("index %d: ", index)
		fmt.Printf("%s, ", line)
		fmt.Printf("sum = %d, ", lineSum)
		fmt.Printf("total = %d\n", total)

	}

	// testing
	// total += getSum(getDigits("1abc2"))
	// a := getSum(getDigits("1abc2"))
	// // a := getSum(getDigitsForRangeLoop("1labc2"))
	// if a == 12 {
	// 	fmt.Println("sum working: yes")
	// } else if a == 540 {
	// 	fmt.Println("bad")
	// }
	// // fmt.Println("First Integer:", 12)

	// fmt.Println(string(rune((total))))
	// fmt.Printf("Type of total: %T\n", total)
	fmt.Println("total", total)

}

func getSum(num1 int, num2 int) int {
	return num1*10 + num2
}

// cases
// no number
// single number => same number repeated
// 2+ numbers
func getDigits(text string) (int, int) {
	// fmt.Println(text)
	// for _, char := range text {
	// 	// fmt.Println(char) // outputs unicode e.g. 116, 108
	// 	fmt.Printf("%c\n", char)
	// }
	var first, last int
	var firstFound, lastFound bool
	i := 0

	for i < len(text) {
		char := text[i] //byte, rune?
		if unicode.IsDigit(rune(char)) {
			if !firstFound {
				// first = int(char)
				// firstRune := rune(char)
				// first = int(firstRune)
				first, _ = strconv.Atoi(string(char))
				last = first

				// fmt.Prsntln("char", char)
				// fmt.Println("firstRune", firstRune)
				// fmt.Println("firstRune int", int(firstRune))
				// fmt.Println("firstRune int string", string(int(firstRune))) // 1 expected
				// fmt.Println(string(char))
				// fmt.Printf("first: %d\n", first)
				// fmt.Println("firstRune int string", first) // 1 expected
				firstFound = true
				lastFound = true
			} else if lastFound {
				// else if firstFound && !lastFound {
				// 	// last = int(rune(char))
				// 	last, _ = strconv.Atoi(string(char))

				// 	lastFound = true
				// }
				// last = int(rune(char))
				last, _ = strconv.Atoi(string(char))
				lastFound = true
			}
		}
		i++
	}

	// fmt.Println("first: " + string(rune(first)))
	// fmt.Println("last: " + string(rune(last)))
	// first = 1
	// fmt.Println("first: ", first)
	// fmt.Printf("printf first: %d\n", first)
	// fmt.Println("last: ", last)
	// fmt.Printf("type of last: %T\n", last)
	// if last == 2 {
	// 	fmt.Println("okayyyy ")

	// }
	return first, last
}

func getDigitsForRangeLoop(text string) (int, int) {
	// fmt.Println(text)
	// for _, char := range text {
	// 	// fmt.Println(char) // outputs unicode e.g. 116, 108
	// 	fmt.Printf("%c\n", char)
	// }
	var first, last int
	var firstFound, lastFound bool
	for _, rune1 := range text {

		if unicode.IsDigit(rune1) {
			if !firstFound {
				// first, _ = strconv.Atoi(rune1)
				// fmt.Println("firstRune int string", string(int(firstRune))) // 1 expected
				// fmt.Println(string(char))
				// fmt.Printf("first: %d\n", first)
				firstFound = true
			} else if firstFound && !lastFound {
				last = int(rune1)
				lastFound = true
			} else if lastFound {
				last = int(rune1)
				lastFound = true
			}
		}

	}

	// fmt.Println("first: " + string(rune(first)))
	// fmt.Println("last: " + string(rune(last)))
	// first = 1
	fmt.Println("first: ", first)
	fmt.Println("last: ", last)
	fmt.Printf("printf first: %d\n", first)
	fmt.Printf("printf last: %d\n", last)
	fmt.Printf("type of last: %T\n", last)
	if last == 2 {
		fmt.Println("okayyyy ")
	}
	// result, err := strconv.Atoi(first)

	return first, last
}
