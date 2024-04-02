package model

type Measurements struct {
	IfName    map[uint]string
	ByteIn    map[uint]uint
	ByteOut   map[uint]uint
	Bandwidth map[uint]uint
}

type SaveMeasurement struct {
	ByteIn    uint
	ByteOut   uint
	Bandwidth uint16
	DateID    uint
	ElementID uint
}
