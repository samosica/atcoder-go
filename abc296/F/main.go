package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type FenwickTree struct {
	Nodes []int
}

func CreateFenwickTree(size int) FenwickTree {
	return FenwickTree{
		Nodes: make([]int, size+1),
	}
}

func (t FenwickTree) Size() int {
	return len(t.Nodes) - 1
}

func (t FenwickTree) Add(k, v int) {
	k += 1
	for ; k <= t.Size(); k += k & -k {
		t.Nodes[k] += v
	}
}

func (t FenwickTree) PrefixSum(k int) int {
	k += 1
	sum := 0
	for ; k > 0; k -= k & -k {
		sum += t.Nodes[k]
	}
	return sum
}

func BytesToInt(bytes []byte) int {
	res := 0
	for _, b := range bytes {
		res = res*10 + (int)(b-'0')
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N := BytesToInt(scanner.Bytes())

	seqs := make([][]int, 2)
	counts := make([]map[int]int, 2)

	for i := 0; i < 2; i += 1 {
		seqs[i] = make([]int, N)
		counts[i] = make(map[int]int)

		for j := 0; j < N; j += 1 {
			scanner.Scan()
			seqs[i][j] = BytesToInt(scanner.Bytes())
			counts[i][seqs[i][j]] += 1
		}
	}

	if !reflect.DeepEqual(counts[0], counts[1]) {
		fmt.Println("No")
		return
	}

	noDuplicates := true

	for _, v := range counts[0] {
		if v > 1 {
			noDuplicates = false
			break
		}
	}

	if !noDuplicates {
		fmt.Println("Yes")
		return
	}

	invs := []int{0, 0}

	for i := 0; i < 2; i += 1 {
		tree := CreateFenwickTree(N + 1)

		for j := N - 1; j >= 0; j -= 1 {
			v := seqs[i][j]
			invs[i] += tree.PrefixSum(v)
			tree.Add(v, 1)
		}
	}

	if invs[0]%2 == invs[1]%2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
