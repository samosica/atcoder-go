package main

import "fmt"

func dfs(v, s int, A []int, visited []int, count *int) {
	visited[v] = s

	w := A[v]
	if visited[w] == s {
		for ; w != v; w = A[w] {
			*count += 1
		}
		*count += 1
		return
	}

	if visited[w] != -1 {
		return
	}

	dfs(w, s, A, visited, count)
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
		A[i] -= 1
	}

	visited := make([]int, N)
	for i := range A {
		visited[i] = -1
	}

	count := 0
	for i := range A {
		if visited[i] == -1 {
			dfs(i, i, A, visited, &count)
		}
	}

	fmt.Println(count)
}
