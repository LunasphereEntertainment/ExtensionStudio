package hacknet

import (
	"encoding/xml"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/actions"
)

type Faction struct {
	xml.Name    `xml:"CustomFaction"`
	ID          string             `xml:"id,attr"`
	FactionName string             `xml:"name,attr"`
	StartValue  int                `xml:"playerVal,attr"`
	ActionSets  []FactionActionSet `xml:"Action"`
	*ProjectResource
}

type FactionActionSet struct {
	RequiresValue *int             `xml:"ValueRequired,attr,omitempty"`
	RequiredFlags CSVList          `xml:"Flags,attr"`
	Actions       []actions.Action `xml:",any"`
}
