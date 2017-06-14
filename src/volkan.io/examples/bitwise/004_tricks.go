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

// This implements a chain of adders (https://en.wikipedia.org/wiki/Adder_(electronics))
// a xor b -> bitwise addition, a and b -> bitwise carryover
// shift the carryover to right to add to the next digit.
func Sum(a int, b int) int {
	if b == 0 {
		return a
	}

	return Sum(a ^ b, (a & b) << 1)
}

func LargestPowerOfTwo(a int) int {
	a |= a>>1
	a |= a>>2
	a |= a>>4
	a |= a>>8
	a |= a>>16

	return (a+1)>>1
}

func ReverseBits(a uint32) uint32 {
	var mask uint32 = 1<<31
	var res uint32 = 0

	for i := 0; i<32; i++ {
		if a & 1 == 1 {
			res |= mask
		}

		mask >>= 1
		a >>= 1
	}

	return res
}

func BitwiseAndAllBetween(m int, n int) int {
	a := 0

	for m != n {
		m >>= 1
		n >>= 1
		a++
	}

	return m<<uint(a)
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
		Sum(42, 42),
		LargestPowerOfTwo(46),
		ReverseBits(268435456),
		BitwiseAndAllBetween(5, 7),
	)
}
