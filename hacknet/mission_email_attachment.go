package hacknet

type AttachmentType string

const (
	Note    AttachmentType = "note"
	Link    AttachmentType = "link"
	Account AttachmentType = "account"
)

type EmailAttachment interface {
	Type() AttachmentType
}

type NoteAttachment struct {
	Title   string `xml:"title,attr" json:"title"`
	Content string `xml:",chardata" json:"content"`
}

func (n *NoteAttachment) Type() AttachmentType {
	return Note
}

type LinkAttachment struct {
	Computer string `xml:"comp,attr" json:"target"`
}

func (l *LinkAttachment) Type() AttachmentType {
	return Link
}

type AccountAttachment struct {
	Computer string `xml:"comp,attr" json:"target"`
	Username string `xml:"user,attr" json:"username"`
	Password string `xml:"pass,attr" json:"password"`
}

func (a *AccountAttachment) Type() AttachmentType {
	return Account
}

//type EmailAttachments []interface{}

/*func (a *EmailAttachments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	partial := new([]EmailAttachment)
	err := d.DecodeElement(partial, &start)
	if err != nil {
		return err
	}

	for _, attach := range *partial {
		switch attach.Type {
		case Note:
			note := new(EmailNote)
			err = d.DecodeElement(note, &start)
			*a = append(*a, note)
		case Link:
			link := new(EmailLink)
			err = d.DecodeElement(link, &start)
			*a = append(*a, link)
		case Account:
			account := new(EmailAccount)
			err = d.DecodeElement(account, &start)
			*a = append(*a, account)
		}

		if err != nil {
			return err
		}
	}

	return err
}*/

/*type EmailAttachment struct {
	Type AttachmentType `xml:"-/"`
	//Title   string         `xml:"title,attr"`
	//Content string         `xml:",chardata"`
	//Target  string         `xml:"target,attr"`
	//User    string         `xml:"user,attr"`
	//Pass    string         `xml:"pass,attr"`
}*/

/*type EmailNote struct {
	xml.Name `xml:"note"`
	*EmailAttachment
	Title   string `xml:"title,attr" json:"title"`
	Content string `xml:",chardata" json:"content"`
}

type EmailLink struct {
	xml.Name `xml:"link"`
	*EmailAttachment
	Target string `xml:"comp,attr"`
}

type EmailAccount struct {
	xml.Name `xml:"account"`
	*EmailLink
	Username string `xml:"user,attr"`
	Password string `xml:"pass,attr"`
}*/
