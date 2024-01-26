package nodes

type EosNote struct {
	Content string `xml:",chardata"`
}

type EosMailAccount struct {
	Username string `xml:"username,attr"`
	Password string `xml:"pass,attr"`
}

type EosDevice struct {
	ID           string           `xml:"id,attr"`
	Name         string           `xml:"name,attr"`
	Icon         ComputerIcon     `xml:"icon,attr"`
	Empty        bool             `xml:"empty,attr"`
	PassOverride string           `xml:"passOverride,attr,omitempty"`
	Notes        []EosNote        `xml:"node"`
	MailAccounts []EosMailAccount `xml:"mail"`
	Files        []File           `xml:"file"`
}
