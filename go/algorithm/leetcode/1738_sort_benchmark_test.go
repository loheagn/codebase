package leetcode

import (
	"math/rand"
	"testing"
)

func randomArr(n, l, r int) []int {
	rand.Seed(619)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 8
	}
	return arr
}

func randomArrs(n, l, r int) [][]int {
	arrs := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		arrs[i] = randomArr(n, l, r)
	}
	return arrs
}

func BenchmarkQuickSort_no_equal(b *testing.B) {
	arrs := randomArrs(1000, 0, 100000)
	for i := 0; i < b.N; i++ {
		for _, arr := range arrs {
			QuickSort_no_equal(arr)
		}
	}
}

func BenchmarkQuickSort_equal(b *testing.B) {
	arrs := randomArrs(1000, 0, 100000)
	for i := 0; i < b.N; i++ {
		for _, arr := range arrs {
			QuickSort_equal(arr)
		}
	}
}
