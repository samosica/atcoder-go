package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i += 1 {
		total := 0

		for j := 0; j < 7; j += 1 {
			var a int
			fmt.Scan(&a)

			total += a
		}

		fmt.Print(total)
		if i+1 < n {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}
