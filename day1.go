package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// mainStrings()

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
		lineSum := getSum(getDigits(line))
		total += lineSum
		if index < 10 {

			fmt.Printf("index %d: ", index)
			fmt.Printf("%s, ", line)
			fmt.Printf("sum = %d, ", lineSum)
			fmt.Printf("total = %d\n", total)
		}

	}

	fmt.Println("total", total) // correct = 55208

}

func getSum(num1 int, num2 int) int {
	return num1*10 + num2
}

// cases
// no number
// single number => same number repeated
// 2+ numbers
// above three + numbers spelled out as letters
func getDigits(text string) (int, int) {
	var first, last int
	var firstFound, lastFound bool
	// validNumbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, char := range text {
		if unicode.IsDigit(char) {
			if !firstFound {
				first, _ = strconv.Atoi(string(char))
				last = first

				firstFound = true
				lastFound = true
			} else if lastFound {

				last, _ = strconv.Atoi(string(char))
				lastFound = true
			}
		} else if unicode.IsLetter(char) {
			// todo implement a trie
			fmt.Println()
		}

	}

	return first, last
}
