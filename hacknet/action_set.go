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
	xml.Name         `xml:"ConditionalActions"`
	OnConnect        *ActionSequence `xml:"OnConnect,omitempty"`
	HasFlags         *ActionSequence `xml:"HasFlags,omitempty"`
	Instantly        *ActionSequence `xml:"Instantly,omitempty"`
	OnAdminGained    *ActionSequence `xml:"OnAdminGained,omitempty"`
	DoesNotHaveFlags *ActionSequence `xml:"DoesNotHaveFlags,omitempty"`
	OnDisconnect     *ActionSequence `xml:"OnDisconnect,omitempty"`
}
