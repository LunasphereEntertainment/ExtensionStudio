package nodes

import (
	"encoding/json"
)

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

func (ap *AdminPass) UnmarshalJSON(raw []byte) error {
	var pass string
	err := json.Unmarshal(raw, &pass)
	if err != nil {
		return err
	}
	ap.Password = pass
	return nil
}

func (ap *AdminPass) MarshalJSON() ([]byte, error) {
	return json.Marshal(ap.Password)
}
