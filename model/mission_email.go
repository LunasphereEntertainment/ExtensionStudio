package model

import (
	"encoding/xml"
)

type Email struct {
	Sender      string           `xml:"sender" json:"sender"`
	Subject     string           `xml:"subject" json:"subject"`
	Body        string           `xml:"body" json:"body"`
	Attachments EmailAttachments `xml:"attachments" json:"attachments"`
}

type AttachmentType string

const (
	Note    AttachmentType = "note"
	Link    AttachmentType = "link"
	Account AttachmentType = "account"
)

type EmailAttachments []interface{}

func (a *EmailAttachments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
}

type EmailAttachment struct {
	Type AttachmentType `xml:"-/"`
	//Title   string         `xml:"title,attr"`
	//Content string         `xml:",chardata"`
	//Target  string         `xml:"target,attr"`
	//User    string         `xml:"user,attr"`
	//Pass    string         `xml:"pass,attr"`
}

type EmailNote struct {
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
}
