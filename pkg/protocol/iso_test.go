//go:build unit

package protocol_test

import (
	"testing"

	"github.com/lucasl0st/nserial/pkg/protocol"

	"github.com/stretchr/testify/assert"
)

func TestGetISO(t *testing.T) {
	tests := []struct {
		speed    byte
		expected uint
	}{
		{
			0x03,
			6,
		},
		{
			0x15,
			400,
		},
		{
			0x21,
			6400,
		},
	}

	for _, test := range tests {
		actual := protocol.GetISO(test.speed)
		assert.Equal(t, test.expected, actual)
	}
}
