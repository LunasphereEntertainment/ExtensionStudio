package nodes

type AdminConfig struct {
	Type          string `xml:"type,attr"`
	ResetPassword bool   `xml:"resetPassword,attr"`
	SuperUser     bool   `xml:"isSuper,attr"`
}
