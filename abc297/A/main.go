package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	var t = make([]int, n)
	for i := range t {
		fmt.Scan(&t[i])
	}

	for i := 0; i+1 < n; i++ {
		if t[i+1]-t[i] <= d {
			fmt.Println(t[i+1])
			return
		}
	}

	fmt.Println(-1)
}
