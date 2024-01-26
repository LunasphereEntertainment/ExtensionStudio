package nodes

import "github.com/LunasphereEntertainment/ExtensionStudio/hacknet"

type AdminConfig struct {
	Type          string `xml:"type,attr"`
	ResetPassword bool   `xml:"resetPassword,attr"`
	SuperUser     bool   `xml:"isSuper,attr"`
}

type ComputerAttrs struct {
	ID                      string       `xml:"id,attr"`
	Name                    string       `xml:"name,attr"`
	IP                      string       `xml:"ip,attr"`
	Security                int          `xml:"security,attr"`
	AllowsDefaultBootModule bool         `xml:"allowsDefaultBootModule,attr"`
	Icon                    ComputerIcon `xml:"icon,attr"`
	Type                    int          `xml:"type,attr"`
}

type ComputerLink struct {
	Target string `xml:"target,attr"`
}

type Computer struct {
	*ComputerAttrs
	AdminPass      AdminPass         `xml:"adminPass"`
	Accounts       []ComputerAccount `xml:"account"`
	Ports          hacknet.CSVList   `xml:"ports"`
	PortsForCrack  PortsForCrack     `xml:"portsForCrack"`
	Proxy          *ProxyConfig      `xml:"proxy,omitempty"`
	Trace          *TraceConfig      `xml:"trace,omitempty"`
	AdminConfig    *AdminConfig      `xml:"admin,omitempty"`
	PortRemap      PortRemap         `xml:"portRemap"`
	HasTracker     Tracker           `xml:"tracker,omitempty"`
	PositionNear   *PositionNear     `xml:"positionNear,omitempty"`
	Files          []File            `xml:"file"`
	CustomThemes   []CustomThemeFile `xml:"customthemefile"`
	EncryptedFiles []EncryptedFile   `xml:"encryptedFile"`
	Links          []ComputerLink    `xml:"dlink"`
	EosLinks       []EosDevice       `xml:"eosDevice"`
}
