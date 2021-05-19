package algorithm

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	tests := []struct {
		src []int
		re  []int
	}{
		{[]int{6, 1, 5, 6, 7, 4, 3, 2}, []int{1, 2, 3, 4, 5, 6, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			QuickSort(tt.src, func(i, j int) bool { return tt.src[i] < tt.src[j] })
			if !reflect.DeepEqual(tt.src, tt.re) {
				t.Errorf("want = %#v,\nbut get = %#v\n", tt.re, tt.src)
			}
		})
	}
}
