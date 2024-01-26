package nodes

import "encoding/xml"

type ProxyConfig struct {
	Duration float64 `xml:"time,attr"`
}

type TraceConfig struct {
	Time int `xml:"time,attr"`
}

type Tracker bool

func (t *Tracker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	err := d.DecodeElement(&raw, &start)
	*t = true
	return err
}

func (t *Tracker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if *t {
		return e.EncodeElement(struct{}{}, start)
	}

	return nil
}
