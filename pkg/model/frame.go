package model

type Frame struct {
	Number uint `json:"number"`

	ShutterSpeed    ShutterSpeed `json:"shutter_speed"`
	ShutterSpeedHex string       `json:"shutter_speed_hex"`

	Aperture    Aperture `json:"aperture"`
	ApertureHex string   `json:"aperture_hex"`

	FocalLength    string `json:"focal_length"`
	FocalLengthHex string `json:"focal_length_hex"`

	MaxAperture    Aperture `json:"max_aperture"`
	MaxApertureHex string   `json:"max_aperture_hex"`

	Raw string `json:"raw"`
}
