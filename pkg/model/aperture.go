package model

import "fmt"

type Aperture struct {
	F       uint
	Decimal uint
}

func (a Aperture) String() string {
	if a.Decimal == 0 {
		return fmt.Sprintf("f/%d", a.F)
	}

	return fmt.Sprintf("f/%d.%d", a.F, a.Decimal)
}
