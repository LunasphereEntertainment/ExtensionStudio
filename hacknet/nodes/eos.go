package nodes

type EosNote struct {
	Content string `xml:",chardata"`
}

type EosMailAccount struct {
	Username string `xml:"username,attr"`
	Password string `xml:"pass,attr"`
}

type EosDevice struct {
	ID           string           `xml:"id,attr" json:"id"`
	Name         string           `xml:"name,attr" json:"name"`
	Icon         ComputerIcon     `xml:"icon,attr" json:"icon"`
	Empty        bool             `xml:"empty,attr" json:"empty"`
	PassOverride string           `xml:"passOverride,attr,omitempty" json:"passOverride"`
	Notes        []EosNote        `xml:"node" json:"notes"`
	MailAccounts []EosMailAccount `xml:"mail" json:"mail"`
	Files        []File           `xml:"file" json:"files"`
}
