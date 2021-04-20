package algorithm

import "testing"

func Test_kmp(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{"abcabcabd", "abcabd"}, want: 3},
		{args: args{"hello", "ll"}, want: 2},
		{args: args{"aaaaa", "bba"}, want: -1},
		{args: args{"", ""}, want: 0},
		{args: args{"", "a"}, want: -1},
		{args: args{"a", ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
