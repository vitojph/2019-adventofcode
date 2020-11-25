package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type orbitMap map[string]map[string]int

type path []string

func (p path) position(planet string) int {
	for i, v := range p {
		if v == planet {
			return i
		}
	}
	return 10000
}

// shortestPath computes the shortest path between two points
func (o orbitMap) shortestPath(p1, p2 string) int {
	path1 := make(path, 0)
	path2 := make(path, 0)
	o.orbitsToCOM(p1, &path1)
	o.orbitsToCOM(p2, &path2)

	intersection1 := len(path1)
	intersection2 := len(path2)
	for i, planet := range path1 {
		if path2.position(planet) < intersection2 {
			intersection1 = i
			intersection2 = path2.position(planet)
		}
	}
	return len(path1[:intersection1]) + len(path2[:intersection2])
}

// orbitsToCOM computes the number of indirect orbits
// from a given planet to COM
func (o orbitMap) orbitsToCOM(planet string, p *path) {
	if planet != "COM" {
		for k := range o[planet] {
			*p = append(*p, k)
			o.orbitsToCOM(k, p)
		}
	}
}

// nOrbitsToCOM computes the number of indirect orbits
// from a given planet to COM
func (o orbitMap) nOrbitsToCOM(planet string) int {
	total := 0
	if planet != "COM" {
		for k := range o[planet] {
			total++
			total += o.nOrbitsToCOM(k)
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
	//fmt.Printf("%v\n", orbits)
	direct := orbits.directOrbits()
	indirect := 0
	for k := range orbits {
		indirect += orbits.nOrbitsToCOM(k)
	}
	fmt.Println("Direct orbits:", direct)
	fmt.Println("Total orbits:", indirect)

	transfers := orbits.shortestPath("SAN", "YOU")
	fmt.Println("Number of orbital transfers:", transfers)
}
