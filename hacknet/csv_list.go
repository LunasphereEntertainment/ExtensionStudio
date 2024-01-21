package hacknet

import (
	"encoding/xml"
	"strings"
)

type CSVList []string

// UnmarshalXML custom unmarshaler for a CSV string to []string
func (c *CSVList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	parts := strings.Split(s, ",")
	for _, part := range parts {
		if len(part) > 0 {
			*c = append(*c, part)
		}
	}

	return nil
}

// UnmarshalXMLAttr enables same functionality as UnmarshalXML for attributes
func (c *CSVList) UnmarshalXMLAttr(attr xml.Attr) error {
	*c = strings.Split(attr.Value, ",")
	return nil
}

// MarshalXML custom marshaller to convert []string to CSV string
func (c CSVList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	raw := strings.Join(c, ",")
	return e.EncodeElement(raw, start)
}
