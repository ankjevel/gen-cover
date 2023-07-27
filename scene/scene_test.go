package scene

import (
	"testing"
)

func Test_New(t *testing.T) {
	width := 240
	height := 480
	res := New(width, height)

	if res.Height != height {
		t.Error("expected Height to match input", res.Height)
	}

	if res.Width != width {
		t.Error("expected Width to match input", res.Width)
	}
}
