package main

import "testing"

func TestCreateThreads(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		in3 int
		want int
	}{
		{-100, 100, 2,25},
		{0, 1000000, 5,78498},
		{167, 189, 500, 4},
	}
	for _, c := range cases {
		got := CreateThreads(c.in1, c.in2, c.in3)
		if got != c.want {
			t.Errorf("CreateThreads(%d, %d, %d) == %d, want %d", c.in1, c.in2, c.in3, got, c.want)
		}
	}
}
