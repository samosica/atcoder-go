package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	// not M <= N^2 ?
	if !((M+N-1)/N <= N) {
		fmt.Println(-1)
		return
	}

	res := 1<<63 - 1

	for X := 1; ; X += 1 {
		Y := (M + X - 1) / X

		if Y > N {
			continue
		}

		if Y < X {
			break
		}

		if prod := X * Y; prod < res {
			res = prod
		}
	}

	fmt.Println(res)
}
