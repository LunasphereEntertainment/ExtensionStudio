package actions

import (
	"encoding/xml"
	"reflect"
)

var (
	ActionTypes = map[string]IAction{
		RunFunction{}.TagName():                RunFunction{},
		LoadMission{}.TagName():                LoadMission{},
		AddAsset{}.TagName():                   AddAsset{},
		CopyAsset{}.TagName():                  CopyAsset{},
		AddMissionToHubServer{}.TagName():      AddMissionToHubServer{},
		RemoveMissionFromHubServer{}.TagName(): RemoveMissionFromHubServer{},
		AddThreadToMissionBoard{}.TagName():    AddThreadToMissionBoard{},
		AddIRCMessage{}.TagName():              AddIRCMessage{},
		CrashComputer{}.TagName():              CrashComputer{},
		DeleteFile{}.TagName():                 DeleteFile{},
		AddConditionalActions{}.TagName():      AddConditionalActions{},
		SaveGame{}.TagName():                   SaveGame{},
		LaunchHackScript{}.TagName():           LaunchHackScript{},
		SwitchToTheme{}.TagName():              SwitchToTheme{},
		StartScreenBleedEffect{}.TagName():     StartScreenBleedEffect{},
		CancelScreenBleedEffect{}.TagName():    CancelScreenBleedEffect{},
		AppendToFile{}.TagName():               AppendToFile{},
		KillExe{}.TagName():                    KillExe{},
		HideNode{}.TagName():                   HideNode{},
		GivePlayerUserAccount{}.TagName():      GivePlayerUserAccount{},
		ChangeIP{}.TagName():                   ChangeIP{},
		HideAllNodes{}.TagName():               HideAllNodes{},
		ShowNode{}.TagName():                   ShowNode{},
		SetLock{}.TagName():                    SetLock{},
	}
)

type IAction interface {
	TagName() string
}

type Action struct {
	Value IAction
}

func (a *Action) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for tagName, elem := range ActionTypes {
		if start.Name.Local == tagName {
			val := reflect.New(reflect.TypeOf(elem)).Elem().Interface()
			if err := d.DecodeElement(&val, &start); err != nil {
				return err
			}
			a.Value = val.(IAction)
			break
		}
	}

	return nil
}

type DelayableAction struct {
	DelayHost *string `xml:"DelayHost,attr,omitempty"`
	Delay     float64 `xml:"Delay,attr"`
}

type TargetableAction struct {
	Target string `xml:"TargetComp,attr"`
}
