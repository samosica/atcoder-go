package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)
	ok := true
	for i := 0; i+1 < n; i += 1 {
		if s[i] == s[i+1] {
			ok = false
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
