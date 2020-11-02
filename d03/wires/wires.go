package wires

import (
	"errors"
	"strconv"
	"strings"
)

// Wires represents a circuit of wires
type Wires struct {
	Circuit [1000][1000]int
	Center  Point
}

// Point has a pair of coordinates
type Point struct {
	X int
	Y int
}

// ReadWirePath reads a wire path and updates a wires circuit
func ReadWirePath(path string, w *Wires) error {
	steps := strings.Split(path, ",")
	pointer := w.Center
	//log.Printf("%#v\n", pointer)
	for _, step := range steps {
		direction := string(step[0])
		movements, _ := strconv.Atoi(step[1:])
		switch direction {
		case "U":
			for i := pointer.Y - 1; i >= pointer.Y-movements; i-- {
				w.Circuit[i][pointer.X]++
			}
			pointer.Y -= movements
		case "D":
			for i := pointer.Y + 1; i < pointer.Y+movements+1; i++ {
				w.Circuit[i][pointer.X]++
			}
			pointer.Y += movements
		case "R":
			for i := pointer.X + 1; i < pointer.X+movements+1; i++ {
				w.Circuit[pointer.Y][i]++
			}
			pointer.X += movements
		case "L":
			for i := pointer.X - 1; i >= pointer.X-movements; i-- {
				w.Circuit[pointer.Y][i]++
			}
			pointer.X -= movements
		default:
			return errors.New("Can't parse wire path")
		}
	}
	return nil
}

// LocateCrosses searches for crosses in the wires
func LocateCrosses(w *Wires) []Point {
	crosses := make([]Point, 0)
	for r := range w.Circuit {
		for c := range w.Circuit[r] {
			if w.Circuit[r][c] > 1 {
				crosses = append(crosses, Point{c, r})
			}
		}
	}
	return crosses
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Manhattan computes Manhattan distance
func Manhattan(w *Wires, crosses []Point) int {
	nearest := 10000000
	for _, cross := range crosses {
		distance := absolute(w.Center.X-cross.X) + absolute(w.Center.Y-cross.Y)
		if distance < nearest {
			nearest = distance
		}
	}
	return nearest
}
