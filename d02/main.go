package main

import (
	"fmt"
)

func resetIntcode() []int {
	return []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 9, 23, 1, 23, 9, 27, 1, 10, 27, 31, 1, 13, 31, 35, 1, 35, 10, 39, 2, 39, 9, 43, 1, 43, 13, 47, 1, 5, 47, 51, 1, 6, 51, 55, 1, 13, 55, 59, 1, 59, 6, 63, 1, 63, 10, 67, 2, 67, 6, 71, 1, 71, 5, 75, 2, 75, 10, 79, 1, 79, 6, 83, 1, 83, 5, 87, 1, 87, 6, 91, 1, 91, 13, 95, 1, 95, 6, 99, 2, 99, 10, 103, 1, 103, 6, 107, 2, 6, 107, 111, 1, 13, 111, 115, 2, 115, 10, 119, 1, 119, 5, 123, 2, 10, 123, 127, 2, 127, 9, 131, 1, 5, 131, 135, 2, 10, 135, 139, 2, 139, 9, 143, 1, 143, 2, 147, 1, 5, 147, 0, 99, 2, 0, 14, 0}
}

func completeGravityAssist(noun, verb int, intcode []int) bool {
	intcode[1] = noun
	intcode[2] = verb

	for pointer := 0; pointer < len(intcode); pointer += 4 {
		opcode := intcode[pointer]
		param1 := intcode[pointer+1]
		param2 := intcode[pointer+2]
		target := intcode[pointer+3]

		switch opcode {
		case 1:
			intcode[target] = intcode[param1] + intcode[param2]
		case 2:
			intcode[target] = intcode[param1] * intcode[param2]
		case 99:
			if intcode[0] == 19690720 {
				return true
			}
			return false
		default:
			return false
		}
	}
	return false
}

func main() {
	fmt.Println("Day 2")

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			res := completeGravityAssist(noun, verb, resetIntcode())
			if res {
				fmt.Println("Noun:", noun)
				fmt.Println("Verb:", verb)
				fmt.Println("Result:", (100*noun)+verb)
			}
			continue
		}
		continue
	}
}
