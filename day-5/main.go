package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rule struct {
	before, after int
}

func GetInput() ([]rule, [][]int) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	parser_in_rules := true

	var rules []rule
	var orderings [][]int

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			parser_in_rules = false
			continue
		}

		if parser_in_rules {
			rules_arr := strings.Split(t, "|")
			before, _ := strconv.Atoi(rules_arr[0])
			after, _ := strconv.Atoi(rules_arr[1])
			rules = append(rules, rule{before, after})
		} else {
			ordering_arr := strings.Split(t, ",")
			ordering := make([]int, len(ordering_arr))
			for i, v := range ordering_arr {
				e, _ := strconv.Atoi(v)
				ordering[i] = e
			}

			orderings = append(orderings, ordering)
		}

	}

	return rules, orderings
}

func CheckCorrectOrder(ordering []int, order map[int]int) bool {
	for i := range ordering[:len(ordering)-1] {
		if order[ordering[i]] < order[ordering[i+1]] {
			return false
		}
	}
	return true
}

func GetOrderForOrdering(ordering []int, rules []rule) map[int]int {
	order := make(map[int]int)

	ordering_set := make(map[int]bool)
	for _, v := range ordering {
		ordering_set[v] = true
	}

	for _, rule := range rules {
		if ordering_set[rule.before] && ordering_set[rule.after] {
			order[rule.before]++
			order[rule.after]--
		}
	}
	return order
}

func GetOrderingFromOrder(ordering []int, order map[int]int) []int {
	sort.Slice(ordering, func(i, j int) bool {
		return order[ordering[i]] > order[ordering[j]]
	})

	return ordering
}

func Star1() {
	rules, orderings := GetInput()

	sum_middle := 0
	for _, v := range orderings {
		order := GetOrderForOrdering(v, rules)
		if CheckCorrectOrder(v, order) {
			sum_middle += v[len(v)/2]
		}
	}

	fmt.Println(sum_middle)
}
func Star2() {
	rules, orderings := GetInput()

	sum_middle := 0
	for _, v := range orderings {
		order := GetOrderForOrdering(v, rules)
		if !CheckCorrectOrder(v, order) {
			new_v := GetOrderingFromOrder(v, order)
			sum_middle += new_v[len(new_v)/2]
		}
	}

	fmt.Println(sum_middle)
}

func main() {
	Star1()
	Star2()
}
