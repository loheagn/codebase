package leetcode

import (
	"reflect"
	"testing"
)

func Test_maximizeXor(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{[]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}}}, want: []int{3, 3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeXor(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
