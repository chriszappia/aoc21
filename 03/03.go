package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	defer f.Close()

	counts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	totalnums := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		totalnums++

		line := scanner.Text()
		for pos, char := range line {
			if char == '1' {
				counts[pos]++
			}
		}
	}

	gammastr := ""
	epsilonstr := ""

	for _, value := range counts {
		if value > (totalnums / 2) {
			gammastr += "1"
			epsilonstr += "0"
		} else {
			gammastr += "0"
			epsilonstr += "1"
		}
	}

	fmt.Printf("G: %s, E: %s\n", gammastr, epsilonstr)
	fmt.Println(counts)

	gammaint, err := strconv.ParseInt(gammastr, 2, 32)
	epsilonint, err := strconv.ParseInt(epsilonstr, 2, 32)
	fmt.Printf("G: %d, E: %d\n", gammaint, epsilonint)
	fmt.Printf("result: %d", epsilonint*gammaint)
}
