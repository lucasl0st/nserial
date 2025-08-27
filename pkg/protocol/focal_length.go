package protocol

type FocalLength struct {
	ByteStart byte
	ByteEnd   byte
	Name      string
}

var FocalLengths = []FocalLength{
	{
		ByteStart: 0x0,
		ByteEnd:   0x13,
		Name:      "6mm",
	},
	{
		ByteStart: 0x14,
		ByteEnd:   0x14,
		Name:      "7.5mm",
	},
	{
		ByteStart: 0x15,
		ByteEnd:   0x17,
		Name:      "8mm",
	},
	{ // Samyang 10mm 2.8 DX == 0x19
		ByteStart: 0x18,
		ByteEnd:   0x23,
		Name:      "10mm",
	},
	{ // Irix 15mm f2.4 == 0x26
		ByteStart: 0x23,
		ByteEnd:   0x28,
		Name:      "15mm",
	},
	{
		ByteStart: 0x29,
		ByteEnd:   0x2c,
		Name:      "16mm",
	},
	{ // Nikon AF-S 18-200mm f3.5-5.6 G VR DX >= 0x2d
		ByteStart: 0x2d,
		ByteEnd:   0x2f,
		Name:      "18mm",
	},
	{
		ByteStart: 0x30,
		ByteEnd:   0x34,
		Name:      "20mm",
	},
	{ // Sigma AF-D 24mm f2.8 == 0x37
		ByteStart: 0x35,
		ByteEnd:   0x3a,
		Name:      "24mm",
	},
	{ // Nikon AF-S 28-70mm f2.8 >= 0x3c
		ByteStart: 0x3b,
		ByteEnd:   0x41,
		Name:      "28mm",
	},
	{ // Nikon AF-S 35mm f1.8 G DX == 0x44
		ByteStart: 0x42,
		ByteEnd:   0x48,
		Name:      "35mm",
	},
	{ // Sigma 50mm f1.4 Art == 0x50
		ByteStart: 0x49,
		ByteEnd:   0x53,
		Name:      "50mm",
	},
	{
		ByteStart: 0x54,
		ByteEnd:   0x5a,
		Name:      "58mm",
	},
	{ // Tamron AF-D 70-210 f4-5.6 >= 0x5c
		ByteStart: 0x5b,
		ByteEnd:   0x5e,
		Name:      "70mm",
	},
	{ // Nikon AF 80-200 f2.8 >= 0x60
		ByteStart: 0x5f,
		ByteEnd:   0x61,
		Name:      "80mm",
	},
	{ // Nikon AF-D 85mm f1.4 == 0x62
		ByteStart: 0x62,
		ByteEnd:   0x65,
		Name:      "85mm",
	},
	{
		ByteStart: 0x66,
		ByteEnd:   0x68,
		Name:      "100mm",
	},
	{ // Nikon AF-D 105mm f2.8 Micro Nikkor == 0x6a
		ByteStart: 0x69,
		ByteEnd:   0x6f,
		Name:      "105mm",
	},
	{ // Nikon AF 80-200 f2.8 <= 0x80
		ByteStart: 0x7c,
		ByteEnd:   0x80,
		Name:      "200mm",
	},
	{ // Tamron AF-D 70-210 f4-5.6 <= 0x81
		ByteStart: 0x81,
		ByteEnd:   0x82,
		Name:      "210mm",
	},
}

func ByteToFocalLength(b byte) string {
	for _, length := range FocalLengths {
		if b < length.ByteStart {
			continue
		}

		if b > length.ByteEnd {
			continue
		}

		return length.Name
	}

	return "0mm"
}
