package main

import "testing"

func Test_try(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Ok string", args{"e1NTSEF9QWkrK1JNMHN1UUxLMDYwaTRMZVlvL0lua3V5NWpuVzQ="}},
		{"Corrupted 1", args{"e1NTSEF9NEE4QkNoQVdGUmxwanZ1ZUJ3UUhRWWd1UTJzaHJLUU5GeFhUeHc9PQ="}},
		{"Corrupted 2", args{"e1NTSEF9NEE4UeHc9PQ="}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			try(tt.args.i)
		})
	}
}

func Test_isASCII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Ascii", args{"e1NTSEF9QWkrK1JNMHN1UUxLMDYwaTRMZVlvL0lua3V5NWpuVzQ="}, true},
		{"Non ASCII", args{"🧡💛💚💙💜"}, false},
		{"Corrupted 2", args{"e1NTSEF9NEE4UeHc9PQ="}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isASCII([]byte(tt.args.s)); got != tt.want {
				t.Errorf("isASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
