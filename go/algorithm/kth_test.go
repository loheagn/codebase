package algorithm

import "testing"

func TestKth(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{[]int{6, 1, 5, 6, 7, 4, 3, 2}, 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kth(tt.args.arr, tt.args.k, func(i, j int) bool { return tt.args.arr[i] < tt.args.arr[j] }); got != tt.want {
				t.Errorf("Kth() = %v, want %v", got, tt.want)
			}
		})
	}
}
