package hghstring

import "testing"

func TestLogPrint(t *testing.T) {
	cases := []int{
		1, 2, 3, 11, 12, 13,
	}

	for _, c := range cases {
		ret := LogPrint(c)

		if !ret {
			t.Errorf("Num is %d\n", c)
		}
	}
}
