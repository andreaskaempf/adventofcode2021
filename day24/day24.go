// Advent of Code 2021, Day 24
//
// Check whether a 14-digit code is valid based on assembly-language like
// series of computations, and find the highest (Part 1) or lowest (Part 2)
// codes that are valid, i.e, leave a "z" register at zero after the
// computations.
//
// Not a good solution, used brute force with Go parallelization and some
// pruning of solutions, came up with the right answers but not efficient
// or clever.
//
// AK, 29 December 2021

package main

import (
	"fmt"
	"os"
	"strconv"
)

// The magic numbers
var A = [14]int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}          // 7 pairs!
var B = [14]int{12, 11, 13, 11, 14, -10, 11, -9, -3, 13, -5, -10, -4, -5} // 7 negs!
var C = [14]int{4, 11, 5, 11, 14, 7, 11, 4, 6, 5, 9, 12, 14, 14}

// Simplified implementation of the original algorithm
func monad3(inpString string, verbose bool) int {

	// The registers (all start at zero)
	var w, z, prevZ, increases int

	// Convert the input into an array of digits
	// TODO: Can we avoid doing this each time, e.g., with list of integers as
	// argument instead of string?
	input := parseInput(inpString)
	if len(input) == 0 {
		return -1
	}

	// Simplified transcription of the problem, entirely in terms of w and z
	// (x,y no longer needed). Quick fail if any digit causes Branch 2 to
	// be chosen when A == 26, or vice versa, or if z keeps increasing.
	for i := 0; i < 14; i++ {
		w = input[i]          // the next digit
		if w == (z%26)+B[i] { // x = 0, y = 0
			if A[i] == 1 {
				return -1 // can't be valid, fail here
			}
			z = div(z, A[i]) // Branch 1
		} else { // x = 1, y = w + C[i]
			if A[i] == 26 {
				return -1 // can't be valid, fail here
			}
			z = div(z, A[i])*26 + w + C[i] // Branch 2
		}

		// z needs to start decreasing at some point, stop if this is not
		// happening by the midway point
		if z > prevZ {
			increases++
		}
		prevZ = z
		if increases > 7 {
			return -1
		}

		// Verbose debugging info
		if verbose {
			fmt.Printf("%2d: w = %d, z = %8d, A = %3d, B = %3d, C = %3d\n",
				i, w, z, A[i], B[i], C[i])
		}
	}

	// Input is considered valid if z is zero at the end
	return z
}

// Parse string of 14 digits into an array of integers,
// return empty slice if not valid (i.e., not 14 long,
// not all digits, or contains zeros)
func parseInput(input string) []int {

	// Convert the input into an array of digits
	result := []int{}
	for i := 0; i < len(input); i++ {
		n := int(input[i] - '0')
		if n < 1 || n > 9 {
			fmt.Println("Invalid digit:", input)
			return []int{}
		}
		result = append(result, n)
	}

	// Check length (can't do this on string because of runes)
	if len(result) != 14 {
		fmt.Println("Invalid length:", input)
		return []int{}
	}

	// Okay
	return result
}

// Check if a string has any zeros in it
func hasZeros(s string) bool {
	for _, c := range s {
		if c == '0' {
			return true
		}
	}
	return false
}

// A range of numbers, used to create a batch for background execution
type Range struct {
	x0, x1 int64
}

// Split all numbers into batches and run them concurrently
func doAll() {

	// Create a list of number ranges to process
	// TODO: Parallelization was added for Part 2, so need to
	// adapt the calculation and use of ranges to increase for Part 2
	//var n int64 = 99999999999999  // For Part 1
	var n int64 = 11111111111111 // For Part 2
	var batchSize int64 = 100000000
	batches := []Range{}
	//for n >= 11111111111111 {  // Part 1
	for n <= 99999999999999 { // Part 2
		b := Range{n, n + batchSize - 1}
		batches = append(batches, b)
		n += batchSize // Decrease for Part 1
	}

	// Input and output queues
	// Concurrent worker pattern borrowed from: https://gobyexample.com/worker-pools
	numJobs := len(batches)
	jobs := make(chan Range, numJobs)
	results := make(chan int, numJobs)

	// Start workers, one per CPU
	for i := 0; i < 8; i++ {
		go worker(jobs, results)
	}

	// Add each demand group to the jobs queue
	for _, b := range batches {
		jobs <- b
	}
	close(jobs)

	// Aggregate the results from the ouput queue
	total := 0
	for i := 0; i < numJobs; i++ {
		res := <-results
		total += res
	}
}

// One worker for concurrent processing, processes any available
// items groups and sends results to output channel
func worker(jobs <-chan Range, results chan<- int) {
	for j := range jobs {
		results <- processItem(j)
	}
}

// Process one range of numbers
// TODO: Before starting the loop, check if this range is feasible, based on
// the leading digits
func processItem(r Range) int {

	fmt.Printf("Processing %d .. %d\n", r.x0, r.x1)
	solutions := 0
	for n := r.x0; n <= r.x1; n++ { // TODO: Need to decrease for Part 1

		// Convert to string
		s := strconv.FormatInt(n, 10)
		if hasZeros(s) { // can we skip over entire block of zeros?
			continue
		}

		// Calculate monad, and report solution if found
		// TODO: Deal with -1 if calculation failed?
		z := monad3(s, false)
		if z == 0 {
			fmt.Println("*** Solution found:", s)
			os.WriteFile(s, []byte("z == 0\n"), 0644)
			solutions += 1
			os.Exit(1)
		}
	}

	fmt.Printf("  finished %d .. %d: %d solutions found\n", r.x0, r.x1, solutions)
	return solutions
}

// Main execution
func main() {

	// Try just one value if command line argument
	if len(os.Args) == 2 {
		s := os.Args[1]
		z := monad3(s, true)
		if z > 0 {
			fmt.Println(s, ": z =", z)
		}
	} else {

		// Otherwise do everything
		doAll()
	}
}
