//go:build unit

package protocol_test

import (
	"testing"

	"github.com/lucasl0st/nserial/pkg/model"
	"github.com/lucasl0st/nserial/pkg/protocol"

	"github.com/stretchr/testify/assert"
)

func TestGetAperture(t *testing.T) {
	tests := []struct {
		speed    byte
		expected model.Aperture
	}{
		{
			0x06,
			model.Aperture{F: 1, Decimal: 4},
		},
		{
			0x12,
			model.Aperture{F: 2, Decimal: 8},
		},
		{
			0x1e,
			model.Aperture{F: 5, Decimal: 6},
		},
	}

	for _, test := range tests {
		actual := protocol.GetAperture(test.speed)
		assert.Equal(t, test.expected.String(), actual.String())
	}
}

func TestGetMaxAperture(t *testing.T) {
	tests := []struct {
		speed    byte
		expected model.Aperture
	}{
		{
			0x0c,
			model.Aperture{F: 1, Decimal: 4},
		},
		{
			0x24,
			model.Aperture{F: 2, Decimal: 8},
		},
		{
			0x3c,
			model.Aperture{F: 5, Decimal: 6},
		},
	}

	for _, test := range tests {
		actual := protocol.GetMaxAperture(test.speed)
		assert.Equal(t, test.expected.String(), actual.String())
	}
}
