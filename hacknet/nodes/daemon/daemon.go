package daemon

import "encoding/xml"

type Daemon struct {
	Value interface{}
}

func (dae *Daemon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	switch start.Name.Local {
	case "mailServer":
		var val MailServer
		err = d.DecodeElement(&val, &start)
		dae.Value = val
	}

	return err
}

type MailServer struct {
	Name         string  `xml:"name,attr"`
	Color        string  `xml:"color,attr"`
	GenerateJunk bool    `xml:"generateJunk,attr"`
	Emails       []Email `xml:"email"`
}

type Email struct {
	Recipient string `xml:"recipient,attr"`
	Sender    string `xml:"sender,attr"`
	Subject   string `xml:"subject,attr"`
	Body      string `xml:",chardata"`
}

type DeathRowDatabase bool

type AcademicDatabase bool

type IspSystem bool
