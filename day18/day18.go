package main

import (
	"fmt"
	"math"
)

// Parse a string into a flat list of tokens. Assumes numbers
// are one or two digit (only single digit for input data, but double
// digits needed for tests), and commas are ignored.
func parse(s string) []byte {

	res := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '[' || c == ']' {
			res = append(res, c)
		} else if isdigit(c) {
			n := s[i] - '0'      // convert this digit to a number
			if isdigit(s[i+1]) { // two digits?
				n = n*10 + (s[i+1] - '0')
				i++ // skip the second digit
			}
			res = append(res, n)
		}
	}
	return res
}

// Print a tokenized expression
func printExpr(expr []byte) {
	for _, b := range expr {
		if b == '[' || b == ']' {
			fmt.Print(string(b), " ")
		} else {
			fmt.Print(int(b), " ")
		}
	}
	fmt.Print("\n")
}

// "Add" two tokenized expressions, but putting both inside a list
func add(a []byte, b []byte) []byte {
	res := []byte{}
	res = append(res, '[')
	for _, c := range a {
		res = append(res, c)
	}
	for _, c := range b {
		res = append(res, c)
	}
	res = append(res, ']')
	return res
}

// "Explode" a pair:
// 1. Find the first pair of numbers that is nested 4 (or more?) deep
// 2. Add the left value of that pair the the next number to its left (if any)
// 3. Add the right value of that pair to the next number on its right (if any)
// 4. Replace the pair (including brackets) with the number zero
// Returns the (possibly changed) expression, and true/false indicating
// whether a change was made.
func explode(expr []byte) ([]byte, bool) {

	// Find the first pair of numbers that is within 4 outer pairs
	level := 0
	pair := -1 // location of explodable pair if found
	for i, b := range expr {
		if level == 4 && b == '[' && isnumber(expr[i+1]) && isnumber(expr[i+2]) && expr[i+3] == ']' {
			pair = i
			break
		} else if b == '[' {
			level += 1
		} else if b == ']' {
			level -= 1
		}
	}

	// If no pair found, stop here
	if pair == -1 {
		return expr, false
	}

	// Find the first number to the left and add left side of pair to it
	for j := pair - 1; j >= 0; j++ {
		if isnumber(expr[j]) {
			expr[j] += expr[pair+1]
			break
		}

	}

	// Find the first number to the right of the pair, and add the right
	// number of the pair to it
	for j := pair + 4; j < len(expr); j++ { // first token after the pair found
		if isnumber(expr[j]) {
			expr[j] += expr[pair+2] // add the second number of the pair
			break
		}
	}

	// Now replace the four characters of the original pair with a zero
	expr[pair] = 0
	res := removeBytes(expr, pair+1, 3)
	return res, true
}

// Go through expression, find any number >= 10, and split it. Returns
// possibly change expression, and true/false if changed or not.
func splitFirst(expr []byte) ([]byte, bool) {

	// Find first number >= 10
	ten := -1
	for i := 0; i < len(expr); i++ {
		if isnumber(expr[i]) && expr[i] >= 10 {
			ten = i
			break
		}
	}

	// Return unchanged if none found
	if ten == -1 {
		fmt.Println("No >= 10 found")
		return expr, false
	}

	// Split the number into [a b]
	sexpr := splitNum(expr[ten])

	// Insert the expression in place of the original number
	res := []byte{}
	res = append(res, expr[:ten]...)
	res = append(res, sexpr...)
	res = append(res, expr[ten+1:]...)

	// Return the changed expression
	return res, true
}

// "Split" a number: replace it with a pair; the left element of the pair
// should be the regular number divided by two and rounded down, while the right
// element of the pair should be the regular number divided by two and rounded up.
// For example, 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6], and so on.
func splitNum(n byte) []byte {
	left := byte(math.Floor(float64(n) / 2))
	right := byte(math.Ceil(float64(n) / 2))
	return []byte{'[', left, right, ']'}
}

// Remove elements from a slice
// res := append(expr[:pair+1], expr[pair+4]...)
func removeBytes(s []byte, from int, nremove int) []byte {
	res := s[:from]
	for i := from + nremove; i < len(s); i++ {
		res = append(res, s[i])
	}
	return res
}

// If a token is not a bracket, it's a number
func isnumber(b byte) bool {
	return b != '[' && b != ']'
}

// Determine character is a digit
func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {

}
