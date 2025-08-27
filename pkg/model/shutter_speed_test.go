//go:build unit

package model_test

import (
	"testing"
	"time"

	"github.com/lucasl0st/nserial/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestShutterSpeed_String(t *testing.T) {
	tests := []struct {
		speed    model.ShutterSpeed
		expected string
	}{
		{
			model.ShutterSpeed(time.Hour),
			"60''",
		},
		{
			model.ShutterSpeed(time.Minute * 30),
			"30''",
		},
		{
			model.ShutterSpeed(time.Minute),
			"1''",
		},
		{
			model.ShutterSpeed(time.Second),
			"1'",
		},
		{
			model.ShutterSpeed(769 * time.Millisecond),
			"1/1.3",
		},
		{
			model.ShutterSpeed(625 * time.Millisecond),
			"1/1.6",
		},
		{
			model.ShutterSpeed(400 * time.Millisecond),
			"1/2.5",
		},
		{
			model.ShutterSpeed(time.Second / 1000),
			"1/1000",
		},
	}

	for _, test := range tests {
		actual := test.speed.String()
		assert.Equal(t, test.expected, actual)
	}
}
