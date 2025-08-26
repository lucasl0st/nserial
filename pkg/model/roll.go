package model

type Roll struct {
	Number uint    `json:"number"`
	ISO    uint    `json:"iso"`
	Frames []Frame `json:"frames"`
	Raw    string  `json:"raw"`
}
