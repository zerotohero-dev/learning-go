package main

import "fmt"

// Set union.
func Union(a int, b int) int {
	return a | b
}

// Set intersection.
func Intersection(a int, b int) int {
	return a & b
}

// This is set subtraction, not regular subtraction.
func Subtraction(a int, b int) int {
	return a & ^b
}

// This is binary negation, not a simple sign inversion.
func Negate(a int) int {
	return ^a
}

func SetBit(a int, bit int) int {
	return a | 1 << uint(bit)
}

func ClearBit(a int, bit int) int {
	return a & ^(1 << uint(bit))
}

func TestBit(a int, bit int) bool {
	return a & (1 << uint(bit)) != 0
}

func ExtractLastBit(a int) int {
	return a & -a
}

func RemoveLastBit(a int) int {
	return a & (a-1)
}

func AllOnes()  int {
	return ^0
}

func CountOnes(num int) int {
	count := 0

	for num != 0 {
		num = num & (num-1)
		count++
	}

	return count
}

func PowerOfTwo(num int) bool {
	return num & (num-1) == 0
}

func PowerOfFour(num int) bool {

	// 0x…5555 ==== …10101
	return (num > 0) && (num & (num-1) == 0) && (num & 0x55555555 != 0)
}

func main() {
	fmt.Println(
		Union( 8, 16),
		Intersection( 8, 16),
		Subtraction(72, 56),
		// -8 is two’s complement; -9 is two’s complement minus one.
		// which is all the bits inverted.
		Negate(8),
		SetBit(8, 2),
		ClearBit(12, 2),
		TestBit(8, 3),
		ExtractLastBit(10),
		RemoveLastBit(10),
		AllOnes(),
		CountOnes(42),
		PowerOfTwo(64),
		PowerOfFour(16777216),
	)
}
