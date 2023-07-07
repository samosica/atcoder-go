package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

func BytesToInt(bytes []byte) int {
	neg := bytes[0] == '-'
	start := 0
	if neg {
		start = 1
	}
	res := 0

	for i := start; i < len(bytes); i += 1 {
		res = res*10 + (int)(bytes[i]-'0')
	}

	if neg {
		res = -res
	}

	return res
}

func ReadInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	return BytesToInt(scanner.Bytes())
}

func Cross(p, q Point) int {
	return p.X*q.Y - p.Y*q.X
}

func PointSub(p, q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	N := ReadInt(scanner)

	nodes := make([]Point, N)
	for i := range nodes {
		nodes[i].X = ReadInt(scanner)
		nodes[i].Y = ReadInt(scanner)
	}

	Q := ReadInt(scanner)

	queries := make([]Point, Q)
	for i := range queries {
		queries[i].X = ReadInt(scanner)
		queries[i].Y = ReadInt(scanner)
	}

	indexes := make([]int, N)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		p := &nodes[indexes[i]]
		q := &nodes[indexes[j]]
		return p.X < q.X || p.X == q.X && p.Y < q.Y
	})

	leftmostBottomIndex := indexes[0]
	leftmostTopIndex := indexes[0]
	if nodes[indexes[1]].X == nodes[indexes[0]].X {
		leftmostTopIndex = indexes[1]
	}

	rightmostBottomIndex := indexes[N-1]
	rightmostTopIndex := indexes[N-1]
	if nodes[indexes[N-2]].X == nodes[indexes[N-1]].X {
		rightmostBottomIndex = indexes[N-2]
	}

	if leftmostBottomIndex > rightmostBottomIndex {
		rightmostBottomIndex += N
	}

	if rightmostTopIndex > leftmostTopIndex {
		leftmostTopIndex += N
	}

	// fmt.Println(leftmostBottomIndex, rightmostBottomIndex, leftmostTopIndex, rightmostTopIndex)

	getPosition := func(p Point) string {
		if p.X < nodes[leftmostBottomIndex].X ||
			nodes[rightmostTopIndex].X < p.X {
			return "OUT"
		}

		if p.X == nodes[leftmostBottomIndex].X &&
			nodes[leftmostBottomIndex].Y <= p.Y &&
			p.Y <= nodes[leftmostTopIndex%N].Y {
			return "ON"
		}

		if p.X == nodes[rightmostBottomIndex%N].X &&
			nodes[rightmostBottomIndex%N].Y <= p.Y &&
			p.Y <= nodes[rightmostTopIndex].Y {
			return "ON"
		}

		lb, ub := leftmostBottomIndex, rightmostBottomIndex

		for ub-lb > 1 {
			m := (lb + ub) / 2

			if nodes[m%N].X <= p.X {
				lb = m
			} else {
				ub = m
			}
		}

		bottomLineIndex := lb
		bottomLineIndex %= N

		lb, ub = rightmostTopIndex, leftmostTopIndex

		for ub-lb > 1 {
			m := (lb + ub) / 2

			if nodes[m%N].X >= p.X {
				lb = m
			} else {
				ub = m
			}
		}

		topLineIndex := lb
		topLineIndex %= N

		// fmt.Printf("bottom line: %d\n", bottomLineIndex)
		// fmt.Printf("top line: %d\n", topLineIndex)

		cb := Cross(
			PointSub(nodes[(bottomLineIndex+1)%N], nodes[bottomLineIndex]),
			PointSub(p, nodes[bottomLineIndex]),
		)

		ct := Cross(
			PointSub(nodes[(topLineIndex+1)%N], nodes[topLineIndex]),
			PointSub(p, nodes[topLineIndex]),
		)

		if cb == 0 || ct == 0 {
			return "ON"
		}

		if cb > 0 && ct > 0 {
			return "IN"
		}

		return "OUT"
	}

	for _, p := range queries {
		fmt.Println(getPosition(p))
	}
}
