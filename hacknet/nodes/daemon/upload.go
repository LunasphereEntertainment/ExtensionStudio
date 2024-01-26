package daemon

type UploadServer struct {
	Name      string `xml:"name,attr"`
	Folder    string `xml:"folder,attr"`
	NeedsAuth bool   `xml:"false,attr"`
	Color     string `xml:"color,attr"`
}
