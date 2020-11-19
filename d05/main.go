package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Program is a list of integers
type Program []int

// readIntProgram reads an IntCode file and
// returns the list of instructions
func readIntProgram(fname string) Program {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	var p Program
	for _, value := range strings.Split(string(data), ",") {
		i, _ := strconv.Atoi(value)
		p = append(p, i)
	}
	return p
}

func getMode(value string, pos int) int {
	pos++
	if len(value) >= pos {
		opmode, err := strconv.Atoi(value[len(value)-pos : len(value)-pos+1])
		if err == nil {
			if opmode == 0 {
				return 0
			}
			return 1
		}
	}
	return 0
}

// runProgram parses the instructions of a program
func runIntProgram(program Program, input int) error {
	output := 0
	for head := 0; head < len(program); head++ {
		v := program[head]
		s := strconv.Itoa(v)

		// the last character is the opcode
		opcode, _ := strconv.Atoi(s[len(s)-1:])
		param1, mode1, param2, mode2, mode3 := 0, 0, 0, 0, 0

		var step int
		// identify the type of instruction and assign the next move
		switch opcode {
		case 1, 2:
			step = 3
			// identify the modes
			mode1 = getMode(s, 2)
			mode2 = getMode(s, 3)
			mode3 = getMode(s, 4)
			// fetch the params
			if mode1 == 0 {
				param1 = program[program[head+1]]
			} else {
				param1 = program[head+1]
			}
			if mode2 == 0 {
				param2 = program[program[head+2]]
			} else {
				param2 = program[head+2]
			}
			// sum
			if opcode == 1 {
				if mode3 == 0 {
					program[program[head+3]] = param1 + param2
				} else {
					program[head+3] = param1 + param2
				}
				//log.Printf("%v sum %v", param1, param2)
			}
			// multiply
			if opcode == 2 {
				if mode3 == 0 {
					program[program[head+3]] = param1 * param2
				} else {
					program[head+3] = param1 * param2
				}
				//log.Printf("%v multiply %v", param1, param2)
			}
		case 3:
			step = 1
			// identify the modes
			mode1 = getMode(s, 2)
			// fetch the params
			if mode1 == 0 {
				input = program[program[head+1]]
			} else {
				input = program[head+1]
			}
			//log.Printf("read %v", input)
		case 4:
			step = 1
			// identify the modes
			mode1 = getMode(s, 2)
			// fetch the params
			if mode1 == 0 {
				output = program[program[head+1]]
			} else {
				output = program[head+1]
			}
			log.Printf("opcode:%v head:%v mode:%v output:%v", v, head, mode1, output)
			log.Printf("write %v in position %v", input, output)
		case 9:
			log.Println("Halt!")
			return nil
		}
		head += step
	}
	return nil
}

func main() {
	fmt.Println("Day 5")

	var input int
	var filename string

	flag.IntVar(&input, "i", 1, "Input instruction")
	flag.StringVar(&filename, "f", "input.txt", "Filename containing the program")
	flag.Parse()

	log.Println("Input instruction:", input)
	log.Println("Program filename:", filename)

	// the program contains a list of instructions
	program := readIntProgram(filename)
	runIntProgram(program, input)
}
