package model

type Count struct {
	ElementID int
	Date      int
	BytesIn   int
	BytesOut  int
	Bandwidth int
}

type CountDiff struct {
	ElementID     int
	PrevDate      int
	PrevBytesIn   int
	PrevBytesOut  int
	CurrDate      int
	CurrBytesIn   int
	CurrBytesOut  int
	CurrBandwidth int
}
