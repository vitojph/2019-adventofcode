package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func totalOrbits(o orbits) int {
	total := 0
	// TODO: compute all the indirect orbits too
	for _, v1 := range o {
		for _, v2 := range v1 {
			total += v2
		}
	}
	return total
}

type orbits map[string]map[string]int

func loadOrbits(fname string) orbits {
	orbits := make(map[string]map[string]int)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bodies := strings.Split(scanner.Text(), ")")
		_, ok1 := orbits[bodies[1]]
		if ok1 {
			_, ok2 := orbits[bodies[1]][bodies[0]]
			if ok2 {
				orbits[bodies[1]][bodies[0]]++
			} else {
				orbits[bodies[1]] = make(map[string]int)
				orbits[bodies[1]][bodies[0]] = 1
			}
		} else {
			orbits[bodies[1]] = make(map[string]int)
			orbits[bodies[1]][bodies[0]] = 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return orbits
}

func main() {
	fmt.Println("Day 6")

	orbits := loadOrbits("input.txt")
	//log.Printf("%v", orbits)

	n := totalOrbits(orbits)
	fmt.Println("Total orbits:", n)

}
