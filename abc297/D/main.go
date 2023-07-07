package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	t := 0
	for a%b > 0 {
		t += a / b
		a, b = b, a%b
	}

	t += a/b - 1

	fmt.Println(t)
}
