package model

type Element struct {
	ID       uint
	Shell    uint
	Card     uint
	Port     uint
	DeviceID uint
}

type PartialElement struct {
	Shell uint
	Card  uint
	Port  uint
}
