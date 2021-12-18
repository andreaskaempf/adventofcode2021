// Unit tests for Day 18

package main

import (
	//"fmt"
	"testing"
)

// Test basic expression operations
func TestParse(t *testing.T) {

	// Test parsing a simple expression
	shouldBe1 := []byte{'[', 1, 2, ']'}
	if !same(parse("[1,2]"), shouldBe1) {
		t.Error("Unable to parse [1,2]")
	}

	// For test cases, we need to be able to parse two-digit numbers
	if !same(parse("[13]"), []byte{'[', 13, ']'}) {
		t.Error("Unable to parse [13]")
	}
}

// Test addition
func TestAdd(t *testing.T) {

	// Test adding [1,2] + [[3,4],5]
	res2 := add(parse("[1,2]"), parse("[[3,4],5]"))
	shouldBe2 := parse("[[1,2],[[3,4],5]]")
	if !same(res2, shouldBe2) {
		t.Error("Unable to add expressions")
	}
}

// Test "exploding" of pairs in expressions
func TestExplode(t *testing.T) {

	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4]
	res, yes := explode(parse("[[[[[9,8],1],2],3],4]"))
	sb := parse("[[[[0,9],2],3],4]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (1)")
	}

	// [7,[6,[5,[4,[3,2]]]]] becomes [7,[6,[5,[7,0]]]]
	res, yes = explode(parse("[7,[6,[5,[4,[3,2]]]]]"))
	sb = parse("[7,[6,[5,[7,0]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (2)")
	}

	// [[6,[5,[4,[3,2]]]],1] becomes [[6,[5,[7,0]]],3]
	res, yes = explode(parse("[[6,[5,[4,[3,2]]]],1]"))
	sb = parse("[[6,[5,[7,0]]],3]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (3)")
	}
	// [[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
	res, yes = explode(parse("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"))
	sb = parse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (4)")
	}
	// [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[7,0]]]]
	res, yes = explode(parse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"))
	sb = parse("[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (5)")
	}
}

func TestSplit(t *testing.T) {

	// 10 becomes [5,5]
	if !same(splitNum(10), parse("[5,5]")) {
		t.Error("Unable to split (1)")
	}

	// 11 becomes [5,6]
	if !same(splitNum(11), parse("[5,6]")) {
		t.Error("Unable to split (2)")
	}

	// 12 becomes [6,6]
	if !same(splitNum(12), parse("[6,6]")) {
		t.Error("Unable to split (3)")
	}

	// [13] becomes [[5,6]]
	res0, yes0 := splitFirst(parse("[13]"))
	if !(yes0 && same(res0, parse("[[6,7]]"))) {
		t.Error("Unable to split first [13]")
		printExpr(parse("[13]"))
		printExpr(res0)
	}

	// [[[[0,7],4],[15,   [0,13]]],[1,1]] becomes
	// [[[[0,7],4],[[7,8],[0,13]]],[1,1]]
	x1 := parse("[[[[0,7],4],[15,[0,13]]],[1,1]]")
	sb1 := parse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")
	res1, yes1 := splitFirst(x1)
	if !(yes1 && same(res1, sb1)) {
		t.Error("Unable to split first (1)")
	}
	// [[[[0,7],4],[[7,8],[0,13]]],   [1,1]] (result from above) becomes
	// [[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]
	x2 := parse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")
	sb2 := parse("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
	res2, yes2 := splitFirst(x2)
	if !(yes2 && same(res2, sb2)) {
		t.Error("Unable to split first (2)")
	}
}

// Compare two byte arrays
func same(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
