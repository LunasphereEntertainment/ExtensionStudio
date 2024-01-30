package nodes

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
)

type PortsForCrack struct {
	Value int `xml:"val,attr"`
}

func (pc *PortsForCrack) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &pc.Value)
}

func (pc *PortsForCrack) MarshalJSON() ([]byte, error) {
	return json.Marshal(pc.Value)
}

type PortRemap map[string]int

func (remap *PortRemap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		raw string
		out = make(map[string]int)
	)
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	parts := strings.Split(raw, ",")
	for _, p := range parts {
		subParts := strings.Split(strings.TrimSpace(p), "=")

		val, err := strconv.Atoi(subParts[1])
		if err != nil {
			return err
		}

		out[subParts[0]] = val
	}

	*remap = out

	return nil
}

func (remap *PortRemap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(remap.String(), start)
}

func (remap PortRemap) String() string {
	sb := strings.Builder{}
	i := 0
	for src, dest := range remap {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(src)
		sb.WriteRune('=')
		sb.WriteString(strconv.Itoa(dest))
		i++
	}

	return sb.String()
}
