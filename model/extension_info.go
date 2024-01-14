package model

import (
	"encoding/xml"
	"strings"
)

type ExtensionInfo struct {
	xml.Name   `xml:"HacknetExtension" json:"-"`
	Title      string `xml:"Name" json:"name"`
	AllowSaves bool   `json:"allowSaves"`
	Language
	StartingNodes   CSVList                    `xml:"StartingVisibleNodes"`
	StartingMission ExternalReference[Mission] `xml:"StartingMission"`
}

func (ext *ExtensionInfo) String() string {
	sb := strings.Builder{}

	sb.WriteString("Extension Name: ")
	sb.WriteString(ext.Title)
	sb.WriteRune('\n')
	sb.WriteString("Allows Saves? ")
	if ext.AllowSaves {
		sb.WriteString("Yes")
	} else {
		sb.WriteString("No")
	}
	sb.WriteRune('\n')
	sb.WriteString("Starting Nodes: ")
	if len(ext.StartingNodes) > 0 {
		sb.WriteString(strings.Join(ext.StartingNodes, ", "))
	} else {
		sb.WriteString("<None>")
	}
	sb.WriteRune('\n')

	sb.WriteString("Starting Mission: ")
	sb.WriteString(string(ext.StartingMission))

	return sb.String()
}
