package main

import (
	"fmt"
	"os"
	"strings"
)

func GetInput() []string {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	m := strings.Split(string(file), "\n")
	return m[:len(m)-1]
}

func CheckSubStringAtCoordStep(s string, m []string, x, y, dx, dy int) bool {
	// this function does not check bounds of m

	for _, v := range s {
		if rune(m[y][x]) != v {
			return false
		}
		y += dy
		x += dx
	}

	return true
}

func CheckSubStringAtCoordStepInt(s string, m []string, x, y, dx, dy int) int {
	if CheckSubStringAtCoordStep(s, m, x, y, dx, dy) {
		return 1
	}
	return 0
}

func CountStringAtCoord(s string, m []string, x, y int) int {
	slen := len(s)
	mh := len(m)
	mw := len(m[0])
	count := 0

	// horizontal
	if mw-x >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, 1, 0)
	}

	// horizontal-backwards
	if x+1 >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, -1, 0)
	}

	// vertical
	if mh-y >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, 0, 1)
	}

	// vertical-backwards
	if y+1 >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, 0, -1)
	}

	// diagonal-right-down
	if mw-x >= slen && mh-y >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, 1, 1)
	}

	// diagonal-right-up
	if mw-x >= slen && y+1 >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, 1, -1)
	}

	// diagonal-left-down
	if x+1 >= slen && mh-y >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, -1, 1)
	}

	// diagonal-left-up
	if x+1 >= slen && y+1 >= slen {
		count += CheckSubStringAtCoordStepInt(s, m, x, y, -1, -1)
	}

	return count
}

func Star1() {
	matrix := GetInput()
	total := 0

	for y, row := range matrix {
		for x := range row {
			total += CountStringAtCoord("XMAS", matrix, x, y)
		}
	}

	fmt.Println(total)
}

func Star2() {
	matrix := GetInput()
	total := 0

	for y, row := range matrix[:len(matrix)-2] {
		for x := range row[:len(row)-2] {
			right_down := CheckSubStringAtCoordStep("MAS", matrix, x, y, 1, 1)
			left_up := CheckSubStringAtCoordStep("SAM", matrix, x, y, 1, 1)
			right_up := CheckSubStringAtCoordStep("MAS", matrix, x, y+2, 1, -1)
			left_down := CheckSubStringAtCoordStep("SAM", matrix, x, y+2, 1, -1)

			if (right_down || left_up) && (right_up || left_down) {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func main() {
	Star1()
	Star2()
}
