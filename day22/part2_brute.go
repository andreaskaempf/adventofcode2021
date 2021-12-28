// Naive first attempt at part 2, runs out of memory
package main

import "fmt"

// Part 2: Do simulation on the whole space (runs out of memory)
func part2_brute_force(instructions []Instruction) int64 {

	// Set up a space with the centres of all 3x3 cubes, recording
	// the state of that cube
	space := map[Point]int64{}

	// Apply the instructions to the space
	for _, i := range instructions {
		volume := (i.x1 - i.x0 + 1) * (i.y1 - i.y0 + 1) * (i.z1 - i.z0 + 1)
		fmt.Println("Processing", i, "-", volume, "points")
		if i.x0 > i.x1 || i.y0 > i.y1 || i.z0 > i.z1 {
			panic("Malformed instruction")
		}
		for x := i.x0; x <= i.x1; x++ {
			for y := i.y0; y <= i.y1; y++ {
				for z := i.z0; z <= i.z1; z++ {
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
