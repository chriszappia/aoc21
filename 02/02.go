package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	horizontal := 0
	vertical := 0

	instructions := 0

	for scanner.Scan() {
		instructions++

		line := scanner.Text()
		parts := strings.Fields(line)
		instruction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		fmt.Println(parts)

		switch {
		case instruction == "forward":
			horizontal += amount
			break
		case instruction == "down":
			vertical += amount
			break
		case instruction == "up":
			vertical -= amount
			break
		default:
			os.Exit(1)
		}

	}

	fmt.Printf("Total: %d\n", instructions)
	fmt.Printf("H: %d, V: %d\n", horizontal, vertical)
	fmt.Printf("Multiplied: %d", horizontal*vertical)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
