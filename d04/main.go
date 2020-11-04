package main

import (
	"fmt"
	"strconv"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

/*

- 6-digit number
- value is within the 284639-748759 range
- Two adjacent digits are the same (like 22 in 122345)
- Going from left to right, the digits never decrease.

*/

func main() {
	fmt.Println("Day 4")

	begining := 284639
	end := 748759
	possiblePasswords := make([]int, 0)
	twoDigits := pcre.MustCompile(`(\d)\1`, 0)
	//digitsNeverDecrease := pcre.MustCompile(`(\d)(\d)(\d)(\d)(\d)(\d)`, 0)

	for i := begining; i <= end; i++ {
		candidate := strconv.Itoa(i)
		m := twoDigits.MatcherString(candidate, 0)

		if m.Matches() {
			if int(candidate[0]) <= int(candidate[1]) && int(candidate[1]) <= int(candidate[2]) && int(candidate[2]) <= int(candidate[3]) && int(candidate[3]) <= int(candidate[4]) && int(candidate[4]) <= int(candidate[5]) {
				possiblePasswords = append(possiblePasswords, i)
			}
		}

	}
	fmt.Printf("I found %v possible passwords\n", len(possiblePasswords))
	fmt.Printf("%v\n", possiblePasswords)
}
