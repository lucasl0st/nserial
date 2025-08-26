package protocol

import (
	"fmt"

	"github.com/lucasl0st/nserial/pkg/model"
)

func ParseRoll(data []byte) model.Roll {
	header := data[:4]
	frame := data[4:]

	number := header[1]
	iso := GetISO(header[2])

	roll := model.Roll{
		Number: uint(number),
		ISO:    iso,
		Raw:    fmt.Sprintf("%x", header),
	}

	for i := 0; i < len(frame); i += 5 {
		frame := ParseFrame(frame[i : i+5])
		roll.Frames = append(roll.Frames, frame)
	}

	return roll
}

func ParseFrame(data []byte) model.Frame {
	shutterSpeed := data[0]
	aperture := data[1]
	focalLength := data[2]
	maxAperture := data[3]
	number := data[4]

	return model.Frame{
		Number: uint(number),

		ShutterSpeed:    GetShutterSpeed(shutterSpeed),
		ShutterSpeedHex: fmt.Sprintf("%x", shutterSpeed),

		Aperture:    GetAperture(aperture),
		ApertureHex: fmt.Sprintf("%x", aperture),

		FocalLength:    ByteToFocalLength(focalLength),
		FocalLengthHex: fmt.Sprintf("%x", focalLength),

		MaxAperture:    GetMaxAperture(maxAperture),
		MaxApertureHex: fmt.Sprintf("%x", maxAperture),

		Raw: fmt.Sprintf("%x", data),
	}
}
