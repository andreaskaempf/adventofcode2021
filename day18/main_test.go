// Unit tests for Day 18

package main

import (
	//"fmt"
	"testing"
)

// Test basic expression operations
func TestParseAdd(t *testing.T) {

	// Test parsing a simple expression
	shouldBe1 := []byte{'[', 1, 2, ']'}
	if !same(parse("[1,2]"), shouldBe1) {
		t.Error("Unable to parse [1,2]")
	}

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
