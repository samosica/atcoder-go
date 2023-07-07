package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var bIndex1 = strings.Index(s, "B")
	var s1 = s[bIndex1+1:]
	var bIndex2 = strings.Index(s1, "B") + bIndex1 + 1

	var rIndex1 = strings.Index(s, "R")
	s1 = s[rIndex1+1:]
	var rIndex2 = strings.Index(s1, "R") + rIndex1 + 1

	var kIndex = strings.Index(s, "K")

	if (bIndex1+bIndex2)%2 == 1 && rIndex1 < kIndex && kIndex < rIndex2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
