// program to benchmark different "PopCount" implementations

package popcount_bench

import "testing"

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i / 2] + byte(i&1)
	}
}

func BenchmarkPopCountTable(b * testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(2199928)
	}
}

func BenchmarkPopCountLoopTable(b * testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoopTable(2199928)
	}
}

func BenchmarkPopCountShift(b * testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoopShift(2199928)
	}
}

func BenchmarkPopCountClearRightMost(b * testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRightMost(2199928)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
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
