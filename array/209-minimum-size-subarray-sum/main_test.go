package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok", args: args{target: 7, nums: []int{1, 2, 3, 4}}, want: 2,
		},
		{
			name: "ok", args: args{target: 4, nums: []int{1, 2, 3, 4}}, want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
