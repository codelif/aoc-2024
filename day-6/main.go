package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Guard struct {
	x, y, dx, dy int
}

func GetInput() ([][]rune, Guard) {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	mn := strings.Split(string(file), "\n")
	if mn[len(mn)-1] == "" {
		mn = mn[:len(mn)-1]
	}
	m := make([][]rune, len(mn))

	for i, v := range mn {
		t := make([]rune, len(mn[0]))
		for j, r := range v {
			t[j] = r
		}
		m[i] = t
	}

	for y, row := range m {
		for x, c := range row {
			if strings.ContainsRune("^v><", c) {
				return m, GetGuard(c, x, y)
			}
		}
	}

	panic("no guard found")
}

func Sin(deg int) int {
	return int(math.Sin(DegToRad(deg)))
}

func Cos(deg int) int {
	return int(math.Cos(DegToRad(deg)))
}
func DegToRad(deg int) float64 {

	return math.Pi * float64(deg) / 180
}
func GetGuard(c rune, x, y int) Guard {
	dir_map := map[rune]int{'>': 0, '^': 90, '<': 180, 'v': 270}
	deg := dir_map[c]

	return Guard{x, y, Cos(deg), -Sin(deg)}
}

func RotateGuard(g *Guard, deg int) {
	// use rotation matrix
	// dy is inverted because left-handed cartesaian coordinates
	new_dx := g.dx*Cos(deg) + g.dy*Sin(deg)
	new_dy := -g.dx*Sin(deg) + g.dy*Cos(deg)

	g.dx = new_dx
	g.dy = new_dy
}

func Next(m [][]rune, g Guard) (bool, rune, int, int) {
	nx := g.x + g.dx
	ny := g.y + g.dy

	mh := len(m)
	mw := len(m[0])

	if nx >= mw || nx < 0 || ny >= mh || ny < 0 {
		return false, rune(' '), nx, ny
	}

	return true, m[ny][nx], nx, ny
}

func Star1() []Coord {
	m, g := GetInput()

	var path []Coord
	path_d := 0
	n, c, _, _ := Next(m, g)
	for n {
		if c == '#' {
			RotateGuard(&g, -90)
			n, c, _, _ = Next(m, g)
			continue
		}
		if m[g.y][g.x] != 'X' {
			path_d += 1
			m[g.y][g.x] = 'X'
		}

		path = append(path, Coord{g.x, g.y})
		g.x += g.dx
		g.y += g.dy
		n, c, _, _ = Next(m, g)
	}

	path = append(path, Coord{g.x, g.y})
	path_d += 1

	fmt.Println(path_d)

	return path
}

func CheckLoop(m [][]rune, ga Guard) bool {
	snapshots := make(map[Guard]bool)
	g := Guard{ga.x, ga.y, ga.dx, ga.dy}
	n, c, _, _ := Next(m, g)
	for n {
		if snapshots[g] {
			return true
		}
		if c == '#' {
			snapshots[Guard{g.x, g.y, g.dx, g.dy}] = true
			RotateGuard(&g, -90)
			n, c, _, _ = Next(m, g)
			continue
		}

		g.x += g.dx
		g.y += g.dy
		n, c, _, _ = Next(m, g)
	}
	return false
}

type Coord struct {
	x, y int
}

func Star2(path []Coord) {
	m, g := GetInput()
	blocked := make(map[Coord]bool)
	loops := 0


  for _, c := range path {
    if (blocked[c]){
      continue
    }
    if (g.x == c.x && g.y == c.y){
      continue
    }
		m[c.y][c.x] = '#'
		if CheckLoop(m, g) {
			loops++
			blocked[c] = true
		}
		m[c.y][c.x] = '.'

	}

	fmt.Println(loops)
}
func main() {
	Star2(Star1())
}
