package colors

import (
	"strconv"
	"strings"
)

var (
	hexes [][5]HEX
	rgbs  [][5]RGB
)

type RGB struct {
	R, G, B uint8
}

type HEX string

func (h HEX) String() string {
	return strings.Replace(string(h), "#", "", 1)
}

func (h HEX) UInt64() (u uint64, err error) {
	return strconv.ParseUint(h.String(), 16, 32)
}

func Get(label string) (res *[5]RGB) {
	var index int
	for _, r := range label {
		index += int(r)
	}
	res = &rgbs[index%len(rgbs)]
	return
}
