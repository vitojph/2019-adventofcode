package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// orbitMap contains an orbit map
type orbitMap map[string]map[string]int

// orbitsToCOM computes the number of indirect orbits
// from a given planet to COM
func (o orbitMap) orbitsToCOM(planet string) int {
	total := 0
	if planet != "COM" {
		for k := range o[planet] {
			total++
			total += o.orbitsToCOM(k)
		}
	}
	return total
}

// directOrbits returns the number of direct orbits
// contained in an orbit map
func (o orbitMap) directOrbits() int {
	total := 0
	// TODO: compute all the indirect orbits too
	for _, v1 := range o {
		for _, v2 := range v1 {
			total += v2
		}
	}
	return total
}

// loadOrbits reads a file defining orbits
// and returns an orbit map
func loadOrbits(fname string) orbitMap {
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
	direct := orbits.directOrbits()
	indirect := 0
	for k := range orbits {
		indirect += orbits.orbitsToCOM(k)
	}
	fmt.Println("Direct orbits:", direct)
	fmt.Println("Total orbits:", indirect)
}
