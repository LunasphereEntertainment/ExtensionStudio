package nodes

type PositionNear struct {
	Target        string  `xml:"target,attr"`
	Position      int     `xml:"position,attr"`
	Total         int     `xml:"total,attr"`
	ExtraDistance float64 `xml:"extraDistance,attr"`
}
