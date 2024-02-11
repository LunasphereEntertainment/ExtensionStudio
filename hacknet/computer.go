package hacknet

import "github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"

type ComputerAttrs struct {
	ID                      string             `xml:"id,attr" json:"id"`
	Name                    string             `xml:"name,attr" json:"name"`
	IP                      string             `xml:"ip,attr" json:"ip"`
	Security                int                `xml:"security,attr" json:"security"`
	AllowsDefaultBootModule bool               `xml:"allowsDefaultBootModule,attr" json:"allowsDefaultBootModule"`
	Icon                    nodes.ComputerIcon `xml:"icon,attr" json:"icon"`
	Type                    nodes.ComputerType `xml:"type,attr" json:"type"`
}

type ComputerLink struct {
	Target string `xml:"target,attr"`
}

type Computer struct {
	*ComputerAttrs
	AdminPass      *nodes.AdminPass        `xml:"adminPass,omitempty" json:"adminPass,omitempty"`
	Accounts       []nodes.ComputerAccount `xml:"account" json:"accounts"`
	Ports          CSVList                 `xml:"ports" json:"ports"`
	PortsForCrack  nodes.PortsForCrack     `xml:"portsForCrack" json:"portsForCrack"`
	Proxy          *nodes.ProxyConfig      `xml:"proxy,omitempty" json:"proxy"`
	Trace          *nodes.TraceConfig      `xml:"trace,omitempty" json:"trace"`
	AdminConfig    *nodes.AdminConfig      `xml:"admin,omitempty" json:"adminConfig"`
	PortRemap      nodes.PortRemap         `xml:"portRemap,omitempty" json:"portRemap"`
	HasTracker     nodes.Tracker           `xml:"tracker,omitempty" json:"hasTracker"`
	PositionNear   *nodes.PositionNear     `xml:"positionNear,omitempty" json:"positionNear"`
	Files          []nodes.File            `xml:"file" json:"files"`
	CustomThemes   []nodes.CustomThemeFile `xml:"customthemefile" json:"customThemes"`
	EncryptedFiles []nodes.EncryptedFile   `xml:"encryptedFile" json:"encryptedFiles"`
	Links          []ComputerLink          `xml:"dlink" json:"links"`
	EosLinks       []nodes.EosDevice       `xml:"eosDevice" json:"eosLinks"`
	*ProjectResource
}
