package nodes

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ComputerType int

const (
	Corporate ComputerType = iota + 1
	Home
	Server
	Empty
)

func (ct *ComputerType) UnmarshalXMLAttr(attr xml.Attr) error {
	raw := attr.Value

	val, err := strconv.Atoi(raw)
	if err != nil {
		switch strings.ToLower(raw) {
		case "corporate":
			*ct = Corporate
		case "home":
			*ct = Home
		case "server":
			*ct = Server
		case "empty":
			*ct = Empty
		default:
			return fmt.Errorf("unrecognised computer type '%s'", raw)
		}
	} else {
		*ct = ComputerType(val)
	}

	return nil
}
