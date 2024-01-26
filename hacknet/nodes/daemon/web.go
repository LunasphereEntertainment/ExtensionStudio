package daemon

type WebServer struct {
	Name string `xml:"name,attr"`
	URL  string `xml:"url,attr"`
}
