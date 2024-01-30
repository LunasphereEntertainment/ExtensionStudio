package nodes

import "github.com/LunasphereEntertainment/ExtensionStudio/hacknet"

type AdminConfig struct {
	Type          string `xml:"type,attr"`
	ResetPassword bool   `xml:"resetPassword,attr"`
	SuperUser     bool   `xml:"isSuper,attr"`
}

type ComputerAttrs struct {
	ID                      string       `xml:"id,attr" json:"id"`
	Name                    string       `xml:"name,attr" json:"name"`
	IP                      string       `xml:"ip,attr" json:"ip"`
	Security                int          `xml:"security,attr" json:"security"`
	AllowsDefaultBootModule bool         `xml:"allowsDefaultBootModule,attr" json:"allowsDefaultBootModule"`
	Icon                    ComputerIcon `xml:"icon,attr" json:"icon"`
	Type                    int          `xml:"type,attr" json:"type"`
}

type ComputerLink struct {
	Target string `xml:"target,attr"`
}

type Computer struct {
	*ComputerAttrs
	AdminPass      *AdminPass        `xml:"adminPass,omitempty" json:"adminPass,omitempty"`
	Accounts       []ComputerAccount `xml:"account" json:"accounts"`
	Ports          hacknet.CSVList   `xml:"ports" json:"ports"`
	PortsForCrack  PortsForCrack     `xml:"portsForCrack" json:"portsForCrack"`
	Proxy          *ProxyConfig      `xml:"proxy,omitempty" json:"proxy"`
	Trace          *TraceConfig      `xml:"trace,omitempty" json:"trace"`
	AdminConfig    *AdminConfig      `xml:"admin,omitempty" json:"adminConfig"`
	PortRemap      PortRemap         `xml:"portRemap,omitempty" json:"portRemap"`
	HasTracker     Tracker           `xml:"tracker,omitempty" json:"hasTracker"`
	PositionNear   *PositionNear     `xml:"positionNear,omitempty" json:"positionNear"`
	Files          []File            `xml:"file" json:"files"`
	CustomThemes   []CustomThemeFile `xml:"customthemefile" json:"customThemes"`
	EncryptedFiles []EncryptedFile   `xml:"encryptedFile" json:"encryptedFiles"`
	Links          []ComputerLink    `xml:"dlink" json:"links"`
	EosLinks       []EosDevice       `xml:"eosDevice" json:"eosLinks"`
}
