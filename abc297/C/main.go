package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		s = strings.ReplaceAll(s, "TT", "PC")
		fmt.Println(s)
	}
}
