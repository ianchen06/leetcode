package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		t []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{t: []int{73, 74, 75, 71, 69, 72, 76, 73}}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{name: "", args: args{t: []int{30, 40, 50, 60}}, want: []int{1, 1, 1, 0}},
		{name: "", args: args{t: []int{30, 60, 90}}, want: []int{1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
