// Advent of Code 2021, Day 22
//
// Turn on/off points in space, defined by 3-d ranges (like rectangular
// cubes). Part 1 quite easy, but solution did not scale to Part 2, so
// that is still pending.
//
// AK, 22 Dec 2021 (Part 1)

package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// An instruction
type Instruction struct {
	onOff                  int64
	x0, x1, y0, y1, z0, z1 int64
}

// A point in space
type Point struct {
	x, y, z int64
}

// Read "instructions" from an input file
func readData(fname string) []Instruction {

	insts := []Instruction{}

	// Read each line of input file, each looking like:
	// on x=11..13,y=11..13,z=11..13
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		t := scanner.Text()

		// on/off beginning of line
		instr := Instruction{}
		parts := strings.Split(t, " ")
		if parts[0] == "on" {
			instr.onOff = 1
		}

		// x/y/z coords (x=11..13)
		for _, xyz := range strings.Split(parts[1], ",") {
			lr := strings.Split(xyz, "=")
			lohi := strings.Split(lr[1], "..")
			if lr[0] == "x" {
				instr.x0 = atoi(lohi[0])
				instr.x1 = atoi(lohi[1])
			} else if lr[0] == "y" {
				instr.y0 = atoi(lohi[0])
				instr.y1 = atoi(lohi[1])
			} else if lr[0] == "z" {
				instr.z0 = atoi(lohi[0])
				instr.z1 = atoi(lohi[1])
			} else {
				panic("Bad letter")
			}

		}

		insts = append(insts, instr)
	}
	return insts
}

// Parse an integer
func atoi(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic("Bad number")
	}
	return i
}

// Simple integer min/max
func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

//  Min/max of a list of numbers
func minmax(l []int64) (int64, int64) {
	var minVal, maxVal int64
	for i, n := range l {
		if i == 0 || n < minVal {
			minVal = n
		}
		if i == 0 || n > maxVal {
			maxVal = n
		}
	}
	return minVal, maxVal
}

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

// Part 2: Do simulation on the whole space (runs out of memory)
func part2(instructions []Instruction) int64 {

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

// Visualize instructions by drawing overlapping squares
// (not complete)
func visualize(instructions []Instruction) {

	// Dimensions
	var w int64 = 1000
	var h int64 = 1000
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(w), int(h)}

	// Create image and define colors
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	red := color.RGBA{200, 10, 10, 0xff}
	green := color.RGBA{10, 200, 10, 0xff}

	// Get min/max coords
	xvals := []int64{}
	yvals := []int64{}
	zvals := []int64{}
	for _, i := range instructions {
		xvals = append(xvals, i.x0, i.x1)
		yvals = append(yvals, i.y0, i.y1)
		zvals = append(zvals, i.z0, i.z1)
	}

	// Calculate scaling factor
	minx, maxx := minmax(xvals)
	xrange := maxx - minx + 1
	miny, maxy := minmax(yvals)
	yrange := maxy - miny + 1
	minz, maxz := minmax(zvals)
	zrange := maxz - minz + 1
	scale := float64(min(w, h)) / float64(max(xrange, max(yrange, zrange)))

	// Draw each instruction as a rectangle
	for _, i := range instructions {
		c := red
		if i.onOff == 1 {
			c = green
		}
		line(img, scale, i.x0, i.y0, i.x1, i.y0, c) // top horiz
		line(img, scale, i.x0, i.y1, i.x1, i.y1, c) // bottom horiz
		line(img, scale, i.x0, i.y0, i.x0, i.y1, c) // left
		line(img, scale, i.x1, i.y0, i.x1, i.y1, c) // right
	}

	//img.Set(x, y, red)
	//img.Set(x, y, color.White)

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func line(img *image.RGBA, scale float64, x0, y0, x1, y1 int64, c color.RGBA) {

}

func main() {

	// Read instructions from a file
	//fname := "sample1.txt"
	//fname := "sample2.txt"
	fname := "sample3.txt" // for part 2
	fmt.Println("Reading", fname)
	instructions := readData(fname)

	// Part 1: Do the simulation
	n := part1(instructions)
	fmt.Printf("%d cubes are on (should be 39 for sample1, 590784 for sample2)\n", n)

	// Visualize the instructions
	//visualize(instructions)

	// Part 2: Do the simulation
	//n := part2(instructions)
	//fmt.Printf("%d cubes are on (should be  2758514936282235 for sample3)\n", n)
}
