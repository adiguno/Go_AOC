package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func mainStrings() {

	const a = "1lab2"
	for index, rune1 := range a {
		fmt.Printf("%d: ", index)

		fmt.Printf("%s", string(rune1))
		fmt.Println()

		if unicode.IsDigit(rune1) {
			fmt.Println("is int (rune)", rune1)
			fmt.Println("is int (string)", string(rune1))
			temp, _ := strconv.Atoi(string(rune1))
			fmt.Println("is int (int)", temp)

		}
	}

	fmt.Println()

	const b = "1"
	c, _ := strconv.Atoi(b)
	fmt.Println("b: ", b)
	fmt.Printf("type of b: %T\n", b)
	fmt.Printf("type of c: %T\n", c)
	fmt.Println("c: ", c)

	fmt.Println()

	// plain string: ⌘
	const placeOfInterest = `⌘` // string literal
	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	// quoted string: "\u2318"
	fmt.Printf("ASCII-only quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	// hex bytes: e2 8c 98
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}
