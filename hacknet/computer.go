package hacknet

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ComputerIcon string

const (
	Laptop     ComputerIcon = "laptop"
	Chip       ComputerIcon = "chip"
	Kellis     ComputerIcon = "kellis"
	Tablet     ComputerIcon = "tablet"
	EPhone     ComputerIcon = "ePhone"
	EPhone2    ComputerIcon = "ePhone2"
	Psylance   ComputerIcon = "Psylance"
	PacificAir ComputerIcon = "PacificAir"
	Alchemist  ComputerIcon = "Alchemist"
	DLCLaptop  ComputerIcon = "DLCLaptop"
	DLCPC1     ComputerIcon = "DLCPC1"
	DLCPC2     ComputerIcon = "DLCPC2"
	DLCServer  ComputerIcon = "DLCServer"
)

type ComputerAttrs struct {
	ID                      string       `xml:"id,attr"`
	Name                    string       `xml:"name,attr"`
	IP                      string       `xml:"ip,attr"`
	Security                int          `xml:"security,attr"`
	AllowsDefaultBootModule bool         `xml:"allowsDefaultBootModule,attr"`
	Icon                    ComputerIcon `xml:"icon,attr"`
	Type                    int          `xml:"type,attr"`
}

type ComputerAccountType string

const (
	Admin       ComputerAccountType = "ADMIN"
	All         ComputerAccountType = "ALL"
	Mail        ComputerAccountType = "MAIL"
	MissionList ComputerAccountType = "MISSIONLIST"
)

type ComputerAccount struct {
	Username string              `xml:"username,attr"`
	Password string              `xml:"password,attr"`
	Type     ComputerAccountType `xml:"type,attr"`
}

type AdminPass struct {
	Password string `xml:"pass,attr"`
}

type PortsForCrack struct {
	Value int `xml:"val,attr"`
}

type ProxyConfig struct {
	Duration float64 `xml:"time,attr"`
}

type TraceConfig struct {
	Time int `xml:"time,attr"`
}

type AdminConfig struct {
	Type          string `xml:"type,attr"`
	ResetPassword bool   `xml:"resetPassword,attr"`
	SuperUser     bool   `xml:"isSuper,attr"`
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

type Tracker bool

func (t *Tracker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}

	*t = true
	return nil
}

func (t Tracker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t {
		var i struct{}
		return e.EncodeElement(&i, start)
	}
	return nil
}

type Computer struct {
	*ComputerAttrs
	AdminPass     AdminPass         `xml:"adminPass"`
	Accounts      []ComputerAccount `xml:"account"`
	Ports         CSVList           `xml:"ports"`
	PortsForCrack PortsForCrack     `xml:"portsForCrack"`
	Proxy         *ProxyConfig      `xml:"proxy,omitempty"`
	Trace         *TraceConfig      `xml:"trace,omitempty"`
	AdminConfig   *AdminConfig      `xml:"admin,omitempty"`
	PortRemap     PortRemap         `xml:"portRemap"`
	HasTracker    Tracker           `xml:"tracker,omitempty"`
}
