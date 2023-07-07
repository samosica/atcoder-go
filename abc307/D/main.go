package main

import (
	"fmt"
)

type Stack []int

func CreateStack() Stack {
	return nil
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) Peek() int {
	return s[len(s)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

type Range struct {
	Left  int
	Right int
}

func main() {
	var (
		N int
		S string
	)
	fmt.Scan(&N, &S)

	lpStack := CreateStack()
	var rangeList []Range

	for i, j := 0, 0; i < N; {
		for j < N && S[j] != '(' && S[j] != ')' {
			j += 1
		}

		if j == N {
			break
		}

		if S[j] == '(' {
			lpStack.Push(j)
		} else if !lpStack.IsEmpty() {
			newRange := Range{
				Left:  lpStack.Pop(),
				Right: j,
			}

			for len(rangeList) > 0 {
				r := rangeList[len(rangeList)-1]
				if newRange.Left < r.Left {
					rangeList = rangeList[:len(rangeList)-1]
				} else {
					break
				}
			}
			rangeList = append(rangeList, newRange)
		}

		j += 1
		i = j
	}

	for i, j := 0, 0; i < N; i += 1 {
		for j < len(rangeList) && rangeList[j].Right < i {
			j += 1
		}

		if j == len(rangeList) || i < rangeList[j].Left {
			fmt.Printf("%c", S[i])
		}
	}
	fmt.Println()
}
