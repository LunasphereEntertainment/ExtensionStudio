package hacknet

import (
	"encoding/xml"
	"strings"
)

type ExtensionInfo struct {
	xml.Name   `xml:"HacknetExtension" json:"-"`
	Title      string `xml:"Name" json:"name"`
	AllowSaves bool   `json:"allowSaves"`
	Language
	StartingNodes      CSVList                                 `xml:"StartingVisibleNodes"`
	StartingMission    ExternalReference[Mission]              `xml:"StartingMission"`
	StartingActions    ExternalReference[ConditionalActionSet] `xml:"StartingActions"`
	Description        string                                  `xml:"Description"`
	Factions           []ExternalReference[Faction]            `xml:"Faction"`
	StartsWithTutorial bool                                    `xml:"StartsWithTutorial"`
	HasIntroStartup    bool                                    `xml:"HasIntroStartup"`
	StartingTheme      ExternalReference[Theme]                `xml:"StartingTheme"`
	IntroStartupSong   string                                  `xml:"IntroStartupSong"`
	// TODO: Sequencer Config
	// TODO: Workshop Info
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
