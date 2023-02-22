package utils

import (
	"chpid/func/utils"
	"regexp"
	"testing"
)

type testGenCheckbitCase struct {
	input string
	want  string
}

var testGenCheckbitGroup = []testGenCheckbitCase{
	{
		input: "21050519200804600",
		want:  "1",
	},
	{
		input: "41102520010209611",
		want:  "3",
	},
	{
		input: "51182719200804611",
		want:  "2",
	},
}

func TestGenCheckbit(t *testing.T) {
	for _, test := range testGenCheckbitGroup {
		output := utils.GenCheckbit(test.input)
		if output != test.want {
			t.Errorf("\nFailed:\"%v\"\ngot:%v\nwant:%v\n", test.input, output, test.want)
		}
	}
}

func TestRandGen(t *testing.T) {
	regRuler := "(^\\d{18}$)|(^\\d{17}(\\d|X|)$)"
	reg := regexp.MustCompile(regRuler)
	for i := 0; i < 20; i++ {
		output := utils.RandGen()
		if !reg.MatchString(output) {
			t.Errorf("RandGen func output not pass regexp check")
		}
	}
}
