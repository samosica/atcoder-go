package main

import "fmt"

func main() {
	var location string

	for i := 8; i >= 1; i -= 1 {
		var s string
		fmt.Scan(&s)

		for j := range s {
			if s[j] == '*' {
				location = fmt.Sprintf("%c%d", byte('a'+j), i)
			}
		}
	}

	fmt.Println(location)
}
