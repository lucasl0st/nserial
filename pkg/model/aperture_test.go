//go:build unit

package model_test

import (
	"testing"

	"github.com/lucasl0st/nserial/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestAperture_String(t *testing.T) {
	tests := []struct {
		speed    model.Aperture
		expected string
	}{
		{
			model.Aperture{F: 2, Decimal: 0},
			"f/2",
		},
		{
			model.Aperture{F: 2, Decimal: 8},
			"f/2.8",
		},
	}

	for _, test := range tests {
		actual := test.speed.String()
		assert.Equal(t, test.expected, actual)
	}
}
