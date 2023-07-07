package main

import (
	"fmt"
)

func IsPalindrome(text string) bool {
	for i := 0; i < len(text)/2; i += 1 {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var N int
	fmt.Scan(&N)

	S := make([]string, N)
	for i := 0; i < N; i += 1 {
		fmt.Scan(&S[i])
	}

	res := false

	for i := 0; i < N; i += 1 {
		for j := 0; j < N; j += 1 {
			if i == j {
				continue
			}

			if IsPalindrome(S[i] + S[j]) {
				res = true
			}
		}
	}

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
