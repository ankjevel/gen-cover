package colors

func fromHEX(hex HEX) (rgb RGB, err error) {
	var values uint64
	if values, err = hex.UInt64(); err != nil {
		return
	}

	rgb = RGB{
		R: uint8(values >> 16),
		G: uint8(values >> 8),
		B: uint8(values),
	}

	return
}

func toRGB(hexes [5]HEX) (result [5]RGB, err error) {
	gen := func() (gen [5]RGB) {
		for i, hex := range hexes {
			var rgb_result RGB
			if rgb_result, err = fromHEX(hex); err != nil {
				break
			}
			gen[i] = rgb_result
		}
		return gen
	}()
	if err == nil {
		result = [5]RGB{gen[4], gen[3], gen[2], gen[1], gen[0]}
	}
	return
}

func init() {
	hexes = [][5]HEX{
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
		if res, err := toRGB(hexes); err != nil {
			continue
		} else {
			rgbs = append(rgbs, res)
		}
	}
}
