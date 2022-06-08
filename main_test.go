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
