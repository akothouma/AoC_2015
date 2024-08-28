package main

import "testing"

func TestFullRibbonDimension(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Full ribbon test",args{"input.txt"},3737498},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullRibbonDimension(tt.args.s); got != tt.want {
				t.Errorf("FullRibbonDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}
