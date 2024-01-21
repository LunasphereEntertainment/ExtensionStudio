package actions

type AddIRCMessage struct {
	*DelayableAction
	*TargetableAction
	Author string `xml:"Author,attr"`
	// TODO: un/marshal attachments
	Content string `xml:",chardata"`
}

func (irc AddIRCMessage) TagName() string {
	return "AddIRCMessage"
}
