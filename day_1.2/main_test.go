package main

import "testing"

func TestFindBasement(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"FirstBasement",args{")"},1},
		{"FirstBasement",args{"()())"},5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBasement(tt.args.s); got != tt.want {
				t.Errorf("FindBasement() = %v, want %v", got, tt.want)
			}
		})
	}
}
