package colors

import (
	"strconv"
	"strings"
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

func fromHEX(hex HEX) (rgb RGB, err error) {
	var values uint64
	if values, err = hex.UInt64(); err != nil {
		return
	}

	rgb = RGB{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
	}

	return
}

var rgbs [][5]RGB

func Get(label string) (res *[5]RGB) {
	var s int
	for _, r := range label {
		s += int(r)
	}
	l := len(rgbs)
	res = &rgbs[s%l]

	return
}

func toRGB(hexes [5]HEX) (result [5]RGB, err error) {
	var h0, h1, h2, h3, h4 RGB
	if h0, err = fromHEX(hexes[0]); err != nil {
		return
	}
	if h1, err = fromHEX(hexes[1]); err != nil {
		return
	}
	if h2, err = fromHEX(hexes[2]); err != nil {
		return
	}
	if h3, err = fromHEX(hexes[3]); err != nil {
		return
	}
	if h4, err = fromHEX(hexes[4]); err != nil {
		return
	}
	result = [5]RGB{h4, h3, h2, h1, h0}
	return
}

func init() {
	hexes := [][5]HEX{
		{"#21AA78", "#42B986", "#64C796", "#85D4A7", "#A6E1BB"},
		{"#949A6C", "#A7AD85", "#BBBF9E", "#CDD1B8", "#E0E2D2"},
		{"#FF9A6C", "#FFAB83", "#FFBC9B", "#FFCDB4", "#FFDFCF"},
		{"#FF006C", "#FF2D85", "#FF599E", "#FF86B8", "#FFB3D2"},
		{"#FFE36C", "#FFEA83", "#FFF09B", "#FFF5B4", "#FFF9CF"},
		{"#FF7F00", "#FF952D", "#FFAC59", "#FFC286", "#FFD9B3"},
		{"#FF7FDA", "#FF93E2", "#FFA8E9", "#FFBEEF", "#FFD5F6"},
		{"#97A0C8", "#A8B0D3", "#BAC0DD", "#CCD1E7", "#DEE2F0"},
	}

	for _, hexes := range hexes {
		res, err := toRGB(hexes)
		if err != nil {
			continue
		}
		rgbs = append(rgbs, res)
	}
}
