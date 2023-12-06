package main_string

import (
	"fmt"
)

func main() {
	const placeOfInterest = `⌘` // string literal
	// plain string: ⌘
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
