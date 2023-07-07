package main

import "fmt"

type Sheet struct {
	Height int
	Width  int
	Grid   [][]bool
}

func CreatePlainSheet(height, width int) Sheet {
	grid := make([][]bool, height)
	for y := 0; y < height; y += 1 {
		grid[y] = make([]bool, width)
	}
	return Sheet{
		Height: height,
		Width:  width,
		Grid:   grid,
	}
}

func ReadSheet() Sheet {
	var sheet Sheet
	fmt.Scan(&sheet.Height)
	fmt.Scan(&sheet.Width)
	sheet.Grid = make([][]bool, sheet.Height)
	for y := 0; y < sheet.Height; y += 1 {
		var row string
		fmt.Scan(&row)

		sheet.Grid[y] = make([]bool, sheet.Width)
		for x := 0; x < sheet.Width; x += 1 {
			sheet.Grid[y][x] = row[x] == '#'
		}
	}
	return sheet
}

func PlaceSheet(s1 Sheet, s2 *Sheet, offsetY, offsetX int) bool {
	for y1 := 0; y1 < s1.Height; y1 += 1 {
		for x1 := 0; x1 < s1.Width; x1 += 1 {
			y2 := y1 + offsetY
			x2 := x1 + offsetX

			if 0 <= y2 && y2 < s2.Height && 0 <= x2 && x2 < s2.Width {
				s2.Grid[y2][x2] = s2.Grid[y2][x2] || s1.Grid[y1][x1]
			} else if s1.Grid[y1][x1] {
				return false
			}
		}
	}

	return true
}

func Matching(s1, s2 Sheet) bool {
	if s1.Height != s2.Height || s1.Width != s2.Width {
		return false
	}

	for y := 0; y < s1.Height; y += 1 {
		for x := 0; x < s1.Width; x += 1 {
			if s1.Grid[y][x] != s2.Grid[y][x] {
				return false
			}
		}
	}

	return true
}

func main() {
	A := ReadSheet()
	B := ReadSheet()
	X := ReadSheet()

	res := false

	for yA := -A.Height + 1; yA < X.Height; yA += 1 {
		for xA := -A.Width + 1; xA < X.Width; xA += 1 {
			for yB := -B.Height + 1; yB < X.Height; yB += 1 {
				for xB := -B.Width + 1; xB < X.Width; xB += 1 {
					C := CreatePlainSheet(X.Height, X.Width)
					ok := PlaceSheet(A, &C, yA, xA)
					if !ok {
						continue
					}

					ok = PlaceSheet(B, &C, yB, xB)
					if !ok {
						continue
					}

					if Matching(C, X) {
						res = true
					}
				}
			}
		}
	}

	// C := CreatePlainSheet(X.Height, X.Width)
	// ok := PlaceSheet(A, &C, 1, 0)
	// fmt.Println(ok)
	// ok = PlaceSheet(B, &C, 1, 0)
	// fmt.Println(ok)

	// for y := 0; y < C.Height; y += 1 {
	// 	for x := 0; x < C.Width; x += 1 {
	// 		if C.Grid[y][x] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(Matching(C, X))

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
