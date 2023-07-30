package colors

import (
	"testing"
)

func Test_Get(t *testing.T) {
	label := "a test label"
	res := Get(label)

	expected_colors, _ := toRGB(hexes[1])

	for i, current := range res {
		expectation := expected_colors[i]
		if expectation != current {
			t.Error("HEX colors not matching", i, current, expectation)
			break
		}
	}
}
