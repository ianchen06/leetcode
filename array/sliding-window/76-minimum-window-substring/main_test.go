package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
		{name: "2", args: args{s: "a", t: "a"}, want: "a"},
		{name: "3", args: args{s: "a", t: "aa"}, want: ""},
		{name: "4", args: args{s: "a", t: "b"}, want: ""},
		{name: "5", args: args{s: "ab", t: "a"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
