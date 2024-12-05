package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func Star1() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	safe_count := 0
	is_safe := true
	asc := true
	var line []int

	for scanner.Scan() {
		line = nil

		line_input := strings.Split(scanner.Text(), " ")
		for _, v := range line_input {
			num, _ := strconv.Atoi(v)
			line = append(line, num)
		}

		for i, v := range line {
			if i == 0 {
				if v < line[1] {
					asc = true
				} else {
					asc = false
				}
				continue
			}

			if (v-line[i-1] > 0) != asc {
				is_safe = false
				break
			}

			if Abs(v, line[i-1]) > 3 || Abs(v, line[i-1]) < 1 {
				is_safe = false
				break
			}

			if i == len(line)-1 {
				is_safe = true
			}
		}

		if is_safe {
			safe_count++
		}
		is_safe = true
	}

	fmt.Println(safe_count)
}

func CheckSafe(line []int) bool {
	is_safe := true
	var asc bool
	for i, v := range line {
		if i == 0 {
			if v < line[1] {
				asc = true
			} else {
				asc = false
			}
			continue
		}

		if (v-line[i-1] >= 0) != asc {
			is_safe = false
			break
		}

		if Abs(v, line[i-1]) > 3 || Abs(v, line[i-1]) < 1 {
			is_safe = false
			break
		}

		if i == len(line)-1 {
			is_safe = true
		}
	}

	return is_safe
}

func Star2() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	safe_count := 0
	is_safe := true
	asc := true
	var line []int

	for scanner.Scan() {
		line = nil

		line_input := strings.Split(scanner.Text(), " ")
		for _, v := range line_input {
			num, _ := strconv.Atoi(v)
			line = append(line, num)
		}
		for i, v := range line {
			if i == 0 {
				if v < line[1] {
					asc = true
				} else {
					asc = false
				}
				continue
			}

			if (v-line[i-1] > 0) != asc {
				is_safe = false
			}

			if Abs(v, line[i-1]) > 3 || Abs(v, line[i-1]) < 1 {
				is_safe = false
			}

			if !is_safe {
				line2 := slices.Clone(line)
				line2 = slices.Delete(line2, 0, 1)
				deep_safe := CheckSafe(line2)
				if deep_safe {
					is_safe = true
					break
				}
				line2 = slices.Clone(line)
				line2 = slices.Delete(line2, i, i+1)
				deep_safe = CheckSafe(line2)
				if deep_safe {
					is_safe = true
					break
				}
				line2 = slices.Clone(line)
				line2 = slices.Delete(line2, i-1, i)
				deep_safe = CheckSafe(line2)
				if deep_safe {
					is_safe = true
				}
				break
			}

			if i == len(line)-1 {
				is_safe = true
			}
		}

		if is_safe {
			safe_count++
		}
		is_safe = true
	}

	fmt.Println(safe_count)
}

func main() {
	Star1()
	Star2()
}
