package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewVent(line string) Vent {
	components := strings.Fields(line)

	point1Components := strings.Split(components[0], ",")
	x1, _ := strconv.Atoi(point1Components[0])
	y1, _ := strconv.Atoi(point1Components[1])

	point2Components := strings.Split(components[2], ",")
	x2, _ := strconv.Atoi(point2Components[0])
	y2, _ := strconv.Atoi(point2Components[1])

	return Vent{x1: x1, y1: y1, x2: x2, y2: y2}
}

func (v Vent) IsStraightLine() bool {
	return v.x1 == v.x2 || v.y1 == v.y2
}

func (v Vent) GetPoints() []Point {
	points := []Point{}
	if v.IsStraightLine() {
		// which axis is the line across
		if v.x1 == v.x2 {
			start := v.y1
			end := v.y2
			if v.y1 > v.y2 {
				start = v.y2
				end = v.y1
			}
			for i := start; i <= end; i++ {
				points = append(points, Point{v.x1, i})
			}
		} else {
			start := v.x1
			end := v.x2
			if v.x1 > v.x2 {
				start = v.x2
				end = v.x1
			}
			for i := start; i <= end; i++ {
				points = append(points, Point{i, v.y1})
			}
		}
	} else {
		// i guess part 2 will say how to deal with not-straight lines
	}
	return points
}

func main() {
	// Main for part 2
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []string{}
	vents := []Vent{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		vents = append(vents, NewVent(line))
	}

	pointCounts := make(map[Point]int)

	for _, vent := range vents {
		if vent.IsStraightLine() {
			for _, point := range vent.GetPoints() {
				if val, ok := pointCounts[point]; ok {
					// it was in the map
					pointCounts[point] = val + 1
				} else {
					pointCounts[point] = 1
				}
			}
		}
	}

	// fmt.Println(vents)
	// fmt.Println(pointCounts)
	count := 0
	for k, v := range pointCounts {
		if v > 1 {
			fmt.Println(k)
			fmt.Println(v)
			count += 1
		}
	}
	fmt.Printf("Count %d\n", count)
}
