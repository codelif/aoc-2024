package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func GetInput() string {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Star1() {
	ins := GetInput()
	total := 0

	for i := range ins {
		if CheckSubStringAtIndex(ins, "mul(", i) && unicode.IsDigit(rune(ins[i+4])) {
			j := i + 4
			for unicode.IsDigit(rune(ins[j])) {
				j++
			}

			if ins[j] != ',' {
				continue
			}

			k := j + 1
			if !unicode.IsDigit(rune(ins[k])) {
				continue
			}

			for unicode.IsDigit(rune(ins[k])) {
				k++
			}

			if ins[k] != ')' {
				continue
			}
			num1, _ := strconv.Atoi(ins[i+4 : j])
			num2, _ := strconv.Atoi(ins[j+1 : k])
			total += num1 * num2
		}
	}

	fmt.Println(total)
}
func CheckSubStringAtIndex(s, ss string, i int) bool {
	a := 0
	for a < len(ss) {
		if s[a+i] != ss[a] {
			return false
		}
		a++
	}

	return true
}

func Star2() {
	ins := GetInput()
	enabled := true
	total := 0

	for i := range ins {
		if CheckSubStringAtIndex(ins, "do()", i) {
			enabled = true
		}
		if CheckSubStringAtIndex(ins, "don't()", i) {
			enabled = false
		}

		if enabled && CheckSubStringAtIndex(ins, "mul(", i) && unicode.IsDigit(rune(ins[i+4])) {
			j := i + 4
			for unicode.IsDigit(rune(ins[j])) {
				j++
			}

			if ins[j] != ',' {
				continue
			}

			k := j + 1
			if !unicode.IsDigit(rune(ins[k])) {
				continue
			}

			for unicode.IsDigit(rune(ins[k])) {
				k++
			}

			if ins[k] != ')' {
				continue
			}
			num1, _ := strconv.Atoi(ins[i+4 : j])
			num2, _ := strconv.Atoi(ins[j+1 : k])
			total += num1 * num2
		}
	}

	fmt.Println(total)
}

func main() {
	Star1()
	Star2()
}
