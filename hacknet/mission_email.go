package hacknet

type Email struct {
	Sender   string              `xml:"sender" json:"sender"`
	Subject  string              `xml:"subject" json:"subject"`
	Body     string              `xml:"body" json:"body"`
	Notes    []NoteAttachment    `xml:"attachments>note"`
	Links    []LinkAttachment    `xml:"attachments>link"`
	Accounts []AccountAttachment `xml:"attachments>account"`
}
