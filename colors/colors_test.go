package colors

import (
	"testing"
)

func Test_Get(t *testing.T) {
	label := "a test label"
	res := Get(label)

	expectations, _ := toRGB([5]HEX{"#949A6C", "#A7AD85", "#BBBF9E", "#CDD1B8", "#E0E2D2"})

	for i, current := range res {
		expectation := expectations[i]
		if expectation != current {
			t.Error("HEX colors not matching", i, current, expectation)
			break
		}
	}
}
