package utils

import (
	"chpid/func/utils"
	"testing"
)

type testDateCase struct {
	input int
	want  string
}

var testSetYearGroup = []testDateCase{
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
	for _, test := range testSetYearGroup {
		testDate := utils.Date{}
		output := testDate.SetYear(test.input)
		if output.Error() != test.want {
			t.Errorf("\nFailed:\"%v\"\ngot:%v\nwant:%v\n", test.input, output, test.want)
		}
	}
}

var testSetMonthGroup = []testDateCase{
	{
		input: 13,
		want:  "wrong month",
	},
	{
		input: 0,
		want:  "wrong month",
	},
	{
		input: -1,
		want:  "wrong month",
	},
}

func TestSetMonth(t *testing.T) {
	for _, test := range testSetMonthGroup {
		testDate := utils.Date{}
		output := testDate.SetMonth(test.input)
		if output.Error() != test.want {
			t.Errorf("\nFailed:\"%v\"\ngot:%v\nwant:%v\n", test.input, output, test.want)
		}
	}
}

var testSetDayGroup = []testDateCase{
	{
		input: 32,
		want:  "wrong day",
	},
	{
		input: 0,
		want:  "wrong day",
	},
	{
		input: -1,
		want:  "wrong day",
	},
}

func TestSetDay(t *testing.T) {
	for _, test := range testSetDayGroup {
		testDate := utils.Date{}
		output := testDate.SetDay(test.input)
		if output.Error() != test.want {
			t.Errorf("\nFailed:\"%v\"\ngot:%v\nwant:%v\n", test.input, output, test.want)
		}
	}
}

// func TestRandBirthday(t *testing.T) {

// }
