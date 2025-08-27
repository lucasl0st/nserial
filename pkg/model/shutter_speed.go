package model

import (
	"fmt"
	"math"
	"time"
)

type ShutterSpeed time.Duration

func (s ShutterSpeed) String() string {
	d := time.Duration(s)

	// < 1s → fractions
	if d < time.Second {
		frac := float64(time.Second) / float64(d)

		// Common Nikon-style denominators with decimals
		// Round to 1 decimal if close enough
		if frac < 10 {
			rounded := math.Round(frac*10) / 10 // one decimal place
			if math.Abs(frac-rounded) < 0.05 {
				return fmt.Sprintf("1/%.1f", rounded)
			}
		}

		// Otherwise fall back to integer denominator
		den := int(math.Round(frac))
		return fmt.Sprintf("1/%d", den)
	}

	// >= 1 minute
	if d >= time.Minute {
		mins := int(d / time.Minute)
		return fmt.Sprintf("%d''", mins)
	}

	// whole seconds
	secs := int(d / time.Second)
	return fmt.Sprintf("%d'", secs)
}
