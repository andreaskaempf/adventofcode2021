// Part 1 solution

package main

import (
	"fmt"
)

// Do simulation only within -50..50 subspace
func part1(instructions []Instruction) int64 {

	// Set up a space with the centres of all 3x3 cubes, recording
	// the state of that cube
	space := map[Point]int64{}

	// Apply the instructions to the space
	for _, i := range instructions {
		fmt.Println("Processing", i)
		if i.x0 > i.x1 || i.y0 > i.y1 || i.z0 > i.z1 {
			panic("Malformed instruction")
		}
		for x := max(-50, i.x0); x <= min(50, i.x1); x++ {
			for y := max(-50, i.y0); y <= min(50, i.y1); y++ {
				for z := max(-50, i.z0); z <= min(50, i.z1); z++ {
					p := Point{x, y, z}
					space[p] = i.onOff
				}
			}
		}
	}

	// Now count up the number of on states
	var n int64 = 0
	for _, v := range space {
		n += v
	}
	return n
}
