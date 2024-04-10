package model

type Snmp struct {
	IfName    map[int]string
	ByteIn    map[int]int
	ByteOut   map[int]int
	Bandwidth map[int]int
}
