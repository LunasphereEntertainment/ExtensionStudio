package hacknet

import (
	"encoding/xml"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/actions"
)

type ActionSequence struct {
	Target               *string          `xml:"target,attr,omitempty"`
	NeedsMissionComplete *bool            `xml:"needsMissionComplete,attr,omitempty"`
	RequiredFlags        CSVList          `xml:"requiredFlags,attr,omitempty"`
	Flags                CSVList          `xml:"Flags,attr,omitempty"`
	Actions              []actions.Action `xml:",any"`
}

type ConditionalActionSet struct {
	xml.Name         `xml:"ConditionalActions" json:"-"`
	OnConnect        []ActionSequence `xml:"OnConnect,omitempty" json:"onConnect"`
	HasFlags         []ActionSequence `xml:"HasFlags,omitempty" json:"hasFlags"`
	Instantly        []ActionSequence `xml:"Instantly,omitempty" json:"instantly"`
	OnAdminGained    []ActionSequence `xml:"OnAdminGained,omitempty" json:"onAdminGained"`
	DoesNotHaveFlags []ActionSequence `xml:"DoesNotHaveFlags,omitempty" json:"doesNotHaveFlags"`
	OnDisconnect     []ActionSequence `xml:"OnDisconnect,omitempty" json:"onDisconnect"`
	*ProjectResource
}
