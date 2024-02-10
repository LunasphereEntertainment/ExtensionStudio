package hacknet

import (
	"encoding/xml"
	"strings"
)

type ExtensionInfo struct {
	xml.Name           `xml:"HacknetExtension" json:"-"`
	Title              string `xml:"Name" json:"name"`
	AllowSaves         bool   `json:"allowSaves"`
	Language           `json:"language"`
	StartingNodes      CSVList                                 `xml:"StartingVisibleNodes" json:"startingNodes"`
	StartingMission    ExternalReference[Mission]              `xml:"StartingMission" json:"startingMission"`
	StartingActions    ExternalReference[ConditionalActionSet] `xml:"StartingActions" json:"startingActions"`
	Description        string                                  `xml:"Description" json:"description"`
	Factions           []ExternalReference[Faction]            `xml:"Faction" json:"factions"`
	StartsWithTutorial bool                                    `xml:"StartsWithTutorial" json:"startsWithTutorial"`
	HasIntroStartup    bool                                    `xml:"HasIntroStartup" json:"hasIntroStartup"`
	StartingTheme      ExternalReference[Theme]                `xml:"StartingTheme" json:"startingTheme"`
	IntroStartupSong   string                                  `xml:"IntroStartupSong" json:"introStartupSong"`
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
	sb.WriteString(ext.StartingMission.Path)

	return sb.String()
}
