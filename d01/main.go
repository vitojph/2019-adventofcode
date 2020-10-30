package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	fuel "github.com/vitojph/2019-adventofcode/d01/fuel"
)

// ReadInputMasses reads a text file containing one integer per line
// and returns a list of ints
func ReadInputMasses(fname string) []int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values []int

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		values = append(values, mass)
	}
	return values
}

func main() {
	fmt.Println("Day 1")

	totalFuel := 0
	masses := ReadInputMasses("input.txt")
	for _, mass := range masses {
		totalFuel += fuel.Fuel(mass)
	}

	fmt.Println("Total fuel:", totalFuel)

}
