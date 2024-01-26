package nodes

type FileHeaders struct {
	Path string `xml:"path,attr"`
	Name string `xml:"name,attr"`
}

type File struct {
	*FileHeaders
	Content string `xml:",chardata"`
}

type CustomThemeFile struct {
	*FileHeaders
	ThemePath string `xml:"themePath,attr"`
}

type EncryptedFile struct {
	*FileHeaders
	IP        string  `xml:"ip,attr,omitempty"`
	Extension string  `xml:"extension,attr,omitempty"`
	Header    string  `xml:"header,attr,omitempty"`
	Password  *string `xml:"pass,attr,omitempty"`
	Content   string  `xml:",chardata"`
}
