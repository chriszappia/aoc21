package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Main for part 2
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	defer f.Close()

	counts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	totalnums := 0
	scanner := bufio.NewScanner(f)

	inputData := []string{}

	for scanner.Scan() {
		totalnums++

		line := scanner.Text()
		inputData = append(inputData, line)
		// count the bits
		for pos, char := range line {
			if char == '1' {
				counts[pos]++
			}
		}
	}

	fmt.Printf("Counts: %d\n", counts)
	// calculate what the expected bit pattern is for the oxygen/c02 measurement
	expectedOxy := []string{}
	expectedC02 := []string{}
	for _, x := range counts {
		if x >= totalnums/2 {
			expectedOxy = append(expectedOxy, "1")
			expectedC02 = append(expectedC02, "0")
		} else {
			expectedOxy = append(expectedOxy, "0")
			expectedC02 = append(expectedC02, "1")
		}
	}
	fmt.Printf("ExpectedOx: %s\n", expectedOxy)
	fmt.Printf("ExpectedC2: %s\n", expectedC02)

	// TODO try without the slice copy
	// oxyCandidates := inputData[:]
	var oxyCandidates []string
	oxyCandidates = append(oxyCandidates, inputData...)
	for index, _ := range inputData[0] {
		// for each character in the input
		// count the values and determine the expected value for this bit
		onesInBit := CountOnes(oxyCandidates, index)
		expectedBit := "0"
		if onesInBit*2 >= len(oxyCandidates) {
			// more ones, expected is 1
			expectedBit = "1"
		}
		fmt.Printf("T: %d C1: %d, Ex: %s\n", len(oxyCandidates), onesInBit, expectedBit)

		oxyCandidates = Filter(oxyCandidates, index, string(expectedBit))
		if len(oxyCandidates) == 1 {
			break
		}
	}
	oxyInt, _ := strconv.ParseInt(oxyCandidates[0], 2, 32)
	fmt.Printf("Candidate Oxy: %s %d\n", oxyCandidates, oxyInt)

	var co2Candidates []string
	co2Candidates = append(co2Candidates, inputData...)
	for index, _ := range inputData[0] {
		// for each character in the input
		// count the values and determine the expected value for this bit
		onesInBit := CountOnes(co2Candidates, index)
		expectedBit := ""
		if onesInBit*2 >= len(co2Candidates) {
			expectedBit = "0"
		} else {
			expectedBit = "1"
		}
		fmt.Printf("T: %d C1: %d, Ex: %s\n", len(co2Candidates), onesInBit, expectedBit)

		co2Candidates = Filter(co2Candidates, index, expectedBit)
		if len(co2Candidates) == 1 {
			break
		}
	}
	co2Int, _ := strconv.ParseInt(co2Candidates[0], 2, 32)
	fmt.Printf("Candidate co2: %s %d\n", co2Candidates, co2Int)
	fmt.Printf("Result: %d", co2Int*oxyInt)
}

func Filter(vals []string, charIndex int, expectedValue string) []string {
	n := 0
	for _, x := range vals {
		if x[charIndex] == expectedValue[0] {
			vals[n] = x
			n++
		}
	}
	vals = vals[:n]
	return vals
}

func CountOnes(vals []string, bitToCount int) int {
	onesInBit := 0
	for _, val := range vals {
		if val[bitToCount] == '1' {
			onesInBit++
		}
	}
	fmt.Printf("vals: %d, bit %d, count %d\n", len(vals), bitToCount, onesInBit)
	return onesInBit
}

func main_1() {
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
