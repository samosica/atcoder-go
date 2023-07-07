package main

import (
	"fmt"
)

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)

	grundy := 0
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		grundy ^= a % (l + r) / l
	}

	if grundy != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
