package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasl0st/nserial/pkg/model"
	"github.com/lucasl0st/nserial/pkg/protocol"
)

func ParseSDF(data []byte) (*model.Roll, error) {
	sdf := string(data)

	number, err := parseSDFRollNumber(sdf)
	if err != nil {
		return nil, err
	}

	iso, err := parseSDFIso(sdf)
	if err != nil {
		return nil, err
	}

	frames, err := parseSDFFrames(sdf)
	if err != nil {
		return nil, err
	}

	return &model.Roll{
		Number: number,
		ISO:    iso,
		Frames: frames,
	}, nil
}

func parseSDFRollNumber(sdf string) (uint, error) {
	value, err := parseSDFKeyValue(sdf, "filmno")
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseUint(strings.TrimSpace(value), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(number), nil
}

func parseSDFIso(sdf string) (uint, error) {
	value, err := parseSDFKeyValue(sdf, "ISO")
	if err != nil {
		return 0, err
	}

	iso, err := strconv.ParseUint(strings.TrimSpace(value), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(iso), nil
}

func parseSDFFrames(sdf string) ([]model.Frame, error) {
	separator, err := parseSDFKeyValue(sdf, "separator")
	if err != nil {
		return nil, err
	}

	key := "data"
	index := strings.Index(sdf, fmt.Sprintf("%s=", key))
	if index == -1 {
		return nil, errors.New("key not found")
	}

	text := sdf[index+len(key)+1:]
	text = collapseTabs(text)
	text = strings.TrimSpace(text)

	var frames []model.Frame

	for _, frameText := range strings.Split(text, "\r\n") {
		parts := strings.Split(frameText, separator)
		if len(parts) < 5 {
			continue
		}

		numberStr := strings.TrimSpace(parts[0])
		number, err := strconv.ParseUint(numberStr, 10, 32)
		if err != nil {
			return nil, err
		}

		shutterSpeedStr := strings.TrimSpace(parts[1])
		shutterSpeed := protocol.ShutterSpeedFromString(shutterSpeedStr)

		apertureStr := strings.TrimSpace(parts[2])
		aperture := protocol.ApertureFromString(apertureStr)
		if aperture.F == 0 && aperture.Decimal == 0 {
			if apertureStr != "F--" && apertureStr != "Lo" && apertureStr != "Hi" {
				return nil, errors.New("failed to parse aperture")
			}
		}

		focalLengthStr := strings.TrimSpace(parts[3])
		maxApertureStr := strings.TrimSpace(parts[4])
		maxAperture := protocol.MaxApertureFromString(maxApertureStr)
		if maxAperture.F == 0 && maxAperture.Decimal == 0 {
			return nil, errors.New("failed to parse max aperture")
		}

		frames = append(frames, model.Frame{
			Number:       uint(number),
			ShutterSpeed: shutterSpeed,
			Aperture:     aperture,
			FocalLength:  focalLengthStr,
			MaxAperture:  maxAperture,
		})
	}

	return frames, nil
}

func parseSDFKeyValue(sdf, key string) (string, error) {
	index := strings.Index(sdf, fmt.Sprintf("%s=", key))
	if index == -1 {
		return "", errors.New("key not found")
	}

	text := sdf[index+len(key)+1:]
	end := strings.Index(text, "\r\n")
	return text[:end], nil
}

var tabCollapse = regexp.MustCompile(`\t+`)

func collapseTabs(s string) string {
	return tabCollapse.ReplaceAllString(s, "\t")
}
