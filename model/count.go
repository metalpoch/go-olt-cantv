package model

type Count struct {
	ElementID int
	DateID    int
	BytesIn   int
	BytesOut  int
	Bandwidth int
}

type CountDiff struct {
	ElementID     int
	PrevDateID    int
	PrevBytesIn   int
	PrevBytesOut  int
	CurrDateID    int
	CurrBytesIn   int
	CurrBytesOut  int
	CurrBandwidth int
}
