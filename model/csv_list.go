package model

import (
	"encoding/xml"
	"strings"
)

type CSVList []string

func (n *CSVList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var listRaw string
	err := d.DecodeElement(&listRaw, &start)
	if err != nil {
		return err
	}

	vals := strings.Split(listRaw, ",")
	// Strip the empty vals off.
	if len(vals) != 1 || len(vals[0]) != 0 {
		*n = vals
	}

	return nil
}

func (n *CSVList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	listRaw := strings.Join(*n, ",")
	return e.EncodeElement(listRaw, start)
}
