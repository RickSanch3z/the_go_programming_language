package main

import "fmt"

// pc[i] is the popution count of i
var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i / 2] + byte(i&1)
	}
}



func main() {
	fmt.Printf("%d\n",PopCount(2199928))
	fmt.Printf("%d\n",PopCountLoopTable(2199928))
	fmt.Printf("%d\n",PopCountLoopShift(2199928))
	fmt.Printf("%d\n",PopCountClearRightMost(2199928))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0 * 8))] +
		pc[byte(x>>(1 * 8))] +
		pc[byte(x>>(2 * 8))] +
		pc[byte(x>>(3 * 8))] +
		pc[byte(x>>(4 * 8))] +
		pc[byte(x>>(5 * 8))] +
		pc[byte(x>>(6 * 8))] +
		pc[byte(x>>(7 * 8))])
}


func PopCountLoopTable(x uint64) int {
	v := 0
	
	for i := 0; i < int(x / 8); i++ {
		v += int(pc[byte(x>>(i * 8))])
	}

	return v
}

func PopCountLoopShift(x uint64) int {
	v := 0

	for x != 0 {
		if x&1 == 1 {
			v++
		}
		x >>= 1
	}

	return v
}

func PopCountClearRightMost(x uint64) int {
	var v int = 0

	for ;x != 0; v++ {
		x &= (x - 1)
	}

	return v
}
