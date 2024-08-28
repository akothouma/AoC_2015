package main

import "testing"

func TestFullWrapperDimension(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Finding wrapper dimension",args{"input.txt"},1586300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullWrapperDimension(tt.args.s); got != tt.want {
				t.Errorf("FullWrapperDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}
