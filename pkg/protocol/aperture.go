package protocol

import "github.com/lucasl0st/nserial/pkg/model"

var byteToAperture = map[byte]model.Aperture{
	0x00: {F: 0, Decimal: 7},  // not tested
	0x01: {F: 0, Decimal: 85}, // not tested
	0x02: {F: 0, Decimal: 95}, // not tested
	0x03: {F: 1, Decimal: 0},  // not tested
	0x04: {F: 1, Decimal: 2},  // not tested, 0x05 could also be f1.2
	0x05: {F: 1, Decimal: 3},  // not tested
	0x06: {F: 1, Decimal: 4},
	0x07: {F: 1, Decimal: 5}, // not tested
	0x08: {F: 1, Decimal: 6},
	0x09: {F: 1, Decimal: 7}, // not tested
	0x0a: {F: 1, Decimal: 8},
	0x0b: {F: 1, Decimal: 9}, // not tested
	0x0c: {F: 2, Decimal: 0},
	0x0d: {F: 2, Decimal: 1}, // not tested
	0x0e: {F: 2, Decimal: 2}, // not tested
	0x0f: {F: 2, Decimal: 4}, // not tested
	0x10: {F: 2, Decimal: 5},
	0x11: {F: 2, Decimal: 6}, // not tested
	0x12: {F: 2, Decimal: 8},
	0x13: {F: 3, Decimal: 0}, // not tested
	0x14: {F: 3, Decimal: 2},
	0x15: {F: 3, Decimal: 4}, // not tested
	0x16: {F: 3, Decimal: 5},
	0x17: {F: 3, Decimal: 8}, // not tested
	0x18: {F: 4, Decimal: 0},
	0x19: {F: 4, Decimal: 2}, // not tested
	0x1a: {F: 4, Decimal: 5},
	0x1b: {F: 4, Decimal: 8}, // not tested
	0x1c: {F: 5, Decimal: 0},
	0x1d: {F: 5, Decimal: 3}, // not tested
	0x1e: {F: 5, Decimal: 6},
	0x1f: {F: 6, Decimal: 0}, // not tested
	0x20: {F: 6, Decimal: 3},
	0x21: {F: 6, Decimal: 7}, // not tested
	0x22: {F: 7, Decimal: 1},
	0x23: {F: 7, Decimal: 6}, // not tested
	0x24: {F: 8, Decimal: 0},
	0x25: {F: 8, Decimal: 5}, // not tested
	0x26: {F: 9, Decimal: 0},
	0x27: {F: 9, Decimal: 5}, // not tested
	0x28: {F: 10, Decimal: 0},
	0x29: {F: 10, Decimal: 5}, // not tested
	0x2a: {F: 11, Decimal: 0},
	0x2b: {F: 12, Decimal: 0}, // not tested
	0x2c: {F: 13, Decimal: 0},
	0x2d: {F: 13, Decimal: 5}, // not tested
	0x2e: {F: 14, Decimal: 0},
	0x2f: {F: 15, Decimal: 0}, // not tested
	0x30: {F: 16, Decimal: 0},
	0x31: {F: 17, Decimal: 0}, // not tested
	0x32: {F: 18, Decimal: 0},
	0x33: {F: 19, Decimal: 0}, // not tested
	0x34: {F: 20, Decimal: 0}, // not tested
	0x35: {F: 21, Decimal: 0}, // not tested
	0x36: {F: 22, Decimal: 0},
	0x37: {F: 23, Decimal: 0}, // not tested
	0x38: {F: 25, Decimal: 0},
	0x39: {F: 27, Decimal: 0}, // not tested
	0x3a: {F: 29, Decimal: 0},
	0x3b: {F: 31, Decimal: 0}, // not tested
	0x3c: {F: 32, Decimal: 0},
	0x3d: {F: 34, Decimal: 0}, // not tested
	0x3e: {F: 36, Decimal: 0},
	0x3f: {F: 38, Decimal: 0}, // not tested
	0x40: {F: 40, Decimal: 0},
	0x41: {F: 43, Decimal: 0}, // not tested
	0x42: {F: 45, Decimal: 0},
	0x43: {F: 48, Decimal: 0}, // not tested
	0x44: {F: 51, Decimal: 0},
	0x46: {F: 57, Decimal: 0},
	0x47: {F: 60, Decimal: 0}, // not tested
	0x48: {F: 64, Decimal: 0}, // not tested
}

func GetAperture(b byte) model.Aperture {
	a, ok := byteToAperture[b]
	if !ok {
		return model.Aperture{}
	}

	return a
}

func ApertureFromString(s string) model.Aperture {
	for _, aperture := range byteToAperture {
		if aperture.String() != s {
			continue
		}

		return aperture
	}

	return model.Aperture{}
}

var byteToMaxAperture = map[byte]model.Aperture{
	0x00: {F: 0, Decimal: 7}, // not tested
	0x01: {F: 0, Decimal: 8}, // not tested
	0x02: {F: 0, Decimal: 9}, // not tested
	0x03: {F: 1, Decimal: 0}, // not tested
	0x04: {F: 1, Decimal: 1}, // not tested
	0x05: {F: 1, Decimal: 2}, // not tested
	0x06: {F: 1, Decimal: 2}, // not tested
	0x07: {F: 1, Decimal: 3}, // not tested
	0x08: {F: 1, Decimal: 3}, // not tested
	0x09: {F: 1, Decimal: 3}, // not tested
	0x0a: {F: 1, Decimal: 4}, // not tested
	0x0b: {F: 1, Decimal: 4}, // not tested
	0x0c: {F: 1, Decimal: 4}, // tested with Sigma 50mm f1.4 Art
	0x0d: {F: 1, Decimal: 5}, // not tested
	0x0e: {F: 1, Decimal: 5}, // not tested
	0x0f: {F: 1, Decimal: 6}, // not tested
	0x10: {F: 1, Decimal: 6}, // not tested
	0x11: {F: 1, Decimal: 7}, // not tested
	0x12: {F: 1, Decimal: 7}, // not tested
	0x13: {F: 1, Decimal: 8}, // not tested
	0x14: {F: 1, Decimal: 8}, // tested with Nikon AF-D 50mm f1.8
	0x15: {F: 1, Decimal: 9}, // not tested
	0x16: {F: 2, Decimal: 0}, // not tested
	0x17: {F: 2, Decimal: 0}, // not tested
	0x18: {F: 2, Decimal: 1}, // not tested
	0x19: {F: 2, Decimal: 2}, // not tested
	0x1a: {F: 2, Decimal: 2}, // not tested
	0x1b: {F: 2, Decimal: 3}, // not tested
	0x1c: {F: 2, Decimal: 3}, // not tested
	0x1d: {F: 2, Decimal: 4}, // not tested
	0x1e: {F: 2, Decimal: 4}, // tested with Irix 15mm f2.4
	0x1f: {F: 2, Decimal: 5}, // not tested
	0x20: {F: 2, Decimal: 5}, // not tested
	0x21: {F: 2, Decimal: 6}, // not tested
	0x22: {F: 2, Decimal: 7}, // not tested
	0x23: {F: 2, Decimal: 8}, // not tested
	0x24: {F: 2, Decimal: 8}, // tested with Nikon AF-D 105mm f2.8 Micro Nikkor
	0x25: {F: 3, Decimal: 0}, // not tested
	0x26: {F: 3, Decimal: 0}, // not tested
	0x27: {F: 3, Decimal: 1}, // not tested
	0x28: {F: 3, Decimal: 2}, // not tested
	0x29: {F: 3, Decimal: 3}, // not tested
	0x2a: {F: 3, Decimal: 4}, // not tested
	0x2b: {F: 3, Decimal: 4}, // not tested
	0x2c: {F: 3, Decimal: 5}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x2d: {F: 3, Decimal: 6}, // not tested
	0x2e: {F: 3, Decimal: 8}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x2f: {F: 3, Decimal: 9}, // not tested
	0x30: {F: 4, Decimal: 0}, // not tested
	0x31: {F: 4, Decimal: 1}, // not tested
	0x32: {F: 4, Decimal: 2}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x33: {F: 4, Decimal: 3}, // not tested
	0x34: {F: 4, Decimal: 4}, // not tested
	0x35: {F: 4, Decimal: 5}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x36: {F: 4, Decimal: 7}, // not tested
	0x37: {F: 4, Decimal: 8}, // not tested
	0x38: {F: 5, Decimal: 0}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x39: {F: 5, Decimal: 1}, // not tested
	0x3a: {F: 5, Decimal: 2}, // not tested
	0x3b: {F: 5, Decimal: 3}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x3c: {F: 5, Decimal: 6}, // tested with Nikon AF-S 18-200mm f3.5-5.6 G VR DX
	0x3d: {F: 5, Decimal: 8}, // not tested
	0x3e: {F: 6, Decimal: 0}, // not tested
	0x3f: {F: 6, Decimal: 1}, // not tested
	0x40: {F: 6, Decimal: 3}, // not tested
	0x41: {F: 6, Decimal: 5}, // not tested
	0x42: {F: 6, Decimal: 7}, // not tested
	0x43: {F: 6, Decimal: 9}, // not tested
	0x44: {F: 7, Decimal: 1}, // not tested
	0x45: {F: 7, Decimal: 3}, // not tested
	0x46: {F: 7, Decimal: 5}, // not tested
	0x47: {F: 7, Decimal: 7}, // not tested
	0x48: {F: 8, Decimal: 0}, // not tested
}

func GetMaxAperture(b byte) model.Aperture {
	a, ok := byteToMaxAperture[b]
	if !ok {
		return model.Aperture{}
	}

	return a
}

func MaxApertureFromString(s string) model.Aperture {
	for _, aperture := range byteToMaxAperture {
		if aperture.String() != s {
			continue
		}

		return aperture
	}

	return model.Aperture{}
}
