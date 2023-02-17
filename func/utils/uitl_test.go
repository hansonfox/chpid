package utils

import (
	"testing"
)

type testDateCase struct {
	input int
	want  string
}

var testDateGroup = []testDateCase{
	{
		input: -12,
		want:  "year can't be negative or equal to 0",
	},
	{
		input: 0,
		want:  "year can't be negative or equal to 0",
	},
	{
		input: 9999,
		want:  "input year is too large",
	},
}

func TestSetYear(t *testing.T) {
	for _, test := range testDateGroup {
		testDate := Date{}
		output := testDate.SetYear(test.input)
		if output.Error() != test.want {
			t.Errorf("\nFailed:\"%v\"\ngot:%v\nwant:%v\n", test.input, output, test.want)
		}
	}
}
