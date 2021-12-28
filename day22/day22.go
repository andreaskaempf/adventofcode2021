// Advent of Code 2021, Day 22
//
// Turn on/off points in space, defined by 3-d ranges (like rectangular
// cubes). Part 1 quite easy, but solution did not scale to Part 2, so
// that is still pending.
//
// AK, 22 Dec 2021 (Part 1)

package main

import (
	"fmt"
)

func main() {

	// Read instructions from a file
	fname := "sample1.txt"
	//fname := "sample2.txt"
	//fname := "sample3.txt" // for part 2
	fmt.Println("Reading", fname)
	instructions := readData(fname)

	// Part 1: Do the simulation
	n := part1(instructions)
	fmt.Printf("%d cubes are on (should be 39 for sample1, 590784 for sample2)\n", n)

	// Part 2: Do the simulation
	//n := part2(instructions)
	//fmt.Printf("%d cubes are on (should be  2758514936282235 for sample3)\n", n)
}
