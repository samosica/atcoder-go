package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)

	for i := range a {
		fmt.Scan(&a[i])
	}

	m := make(map[int]struct{})
	ok := false

	for i := range a {
		m[a[i]] = struct{}{}
		_, ex1 := m[a[i]-x]
		_, ex2 := m[a[i]+x]
		if ex1 || ex2 {
			ok = true
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
