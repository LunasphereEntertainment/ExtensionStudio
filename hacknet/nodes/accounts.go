package nodes

type ComputerAccountType string

const (
	Admin       ComputerAccountType = "ADMIN"
	All         ComputerAccountType = "ALL"
	Mail        ComputerAccountType = "MAIL"
	MissionList ComputerAccountType = "MISSIONLIST"
)

type ComputerAccount struct {
	Username string              `xml:"username,attr"`
	Password string              `xml:"password,attr"`
	Type     ComputerAccountType `xml:"type,attr"`
}

type AdminPass struct {
	Password string `xml:"pass,attr"`
}
