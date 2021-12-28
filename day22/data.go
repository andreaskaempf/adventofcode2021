// Data structures and read input files

package main

import (
	"bufio"
	"os"
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
