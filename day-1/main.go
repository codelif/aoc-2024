package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetInput() ([]int, []int) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line_input := strings.Split(scanner.Text(), "   ")
		temp, _ := strconv.Atoi(line_input[0])
		list1 = append(list1, temp)
		temp, _ = strconv.Atoi(line_input[1])
		list2 = append(list2, temp)
	}

	return list1, list2
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Star1() {
	list1, list2 := GetInput()

	sort.Ints(list1)
	sort.Ints(list2)

	total_distance := 0

	for i, v := range list1 {
		total_distance += Abs(v, list2[i])
	}

	fmt.Println(total_distance)
}

func GenerateFrequencyMap(list []int) map[int]int {
	freq := make(map[int]int)

	for _, v := range list {
		freq[v]++
	}

	return freq
}

func Star2() {
	list1, list2 := GetInput()

	freqs := GenerateFrequencyMap(list2)
	done := make(map[int]bool)
	similarity_score := 0

	for _, v := range list1 {
		if done[v] {
			continue
		}
		similarity_score += v * freqs[v]
		done[v] = true
	}

	fmt.Println(similarity_score)
}

func main() {
	Star1()
	Star2()
}
