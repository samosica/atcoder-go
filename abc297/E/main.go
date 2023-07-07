package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/binaryheap"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var a = make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	h := binaryheap.NewWithIntComparator()
	m := make(map[int]struct{})

	h.Push(0)
	for ; k > 0; k-- {
		minV_, _ := h.Pop()
		minV := minV_.(int)

		for i := range a {
			v := minV + a[i]
			_, found := m[v]
			if !found {
				h.Push(v)
				m[v] = struct{}{}
			}
		}
	}

	minV_, _ := h.Peek()
	minV := minV_.(int)
	fmt.Println(minV)
}
