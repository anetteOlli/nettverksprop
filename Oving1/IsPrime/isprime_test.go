package IsPrime

import "testing"

func TestIsPrime(t *testing.T) {
	cases := []struct {
		in int
		want bool
	}{
		{10, false},
		{7, true},
		{149, true},
	}
	for _, c := range cases {
		got := IsPrime(c.in)
		if got != c.want {
			t.Errorf("IsPrime(%b) == %t, want %t", c.in, got, c.want)
		}
	}
}
