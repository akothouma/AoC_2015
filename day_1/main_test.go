package main

import (
	"testing"
)

func TestFindFloor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Incrementing bracket",args{"("},1},
		{"Decrementing bracket",args{")"},-1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFloor(tt.args.s); got != tt.want {
				t.Errorf("FindFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
