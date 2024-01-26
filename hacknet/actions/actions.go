package actions

import (
	"bytes"
	"encoding/xml"
)

type IAction interface {
	TagName() string
}

type Action struct {
	Value IAction
}

func (a *Action) XML() string {
	buff := bytes.Buffer{}
	err := xml.NewEncoder(&buff).Encode(a.Value)
	if err != nil {
		panic(err)
	}

	return buff.String()
}

func (a *Action) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "RunFunction":
		val := RunFunction{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "LoadMission":
		val := LoadMission{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AddAsset":
		val := AddAsset{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "CopyAsset":
		val := CopyAsset{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AddMissionToHubServer":
		val := AddMissionToHubServer{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "RemoveMissionFromHubServer":
		val := RemoveMissionFromHubServer{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AddThreadToMissionBoard":
		val := AddThreadToMissionBoard{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AddIRCMessage":
		val := AddIRCMessage{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "CrashComputer":
		val := CrashComputer{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "DeleteFile":
		val := DeleteFile{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AddConditionalActions":
		val := AddConditionalActions{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "SaveGame":
		val := SaveGame{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "LaunchHackScript":
		val := LaunchHackScript{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "SwitchToTheme":
		val := SwitchToTheme{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "StartScreenBleedEffect":
		val := StartScreenBleedEffect{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "CancelScreenBleedEffect":
		val := CancelScreenBleedEffect{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "AppendToFile":
		val := AppendToFile{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "KillExe":
		val := KillExe{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "HideNode":
		val := HideNode{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "GivePlayerUserAccount":
		val := GivePlayerUserAccount{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "ChangeIP":
		val := ChangeIP{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "HideAllNodes":
		val := HideAllNodes{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "ShowNode":
		val := ShowNode{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "SetLock":
		val := SetLock{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
		a.Value = val
	case "ChangeAlertIcon":
		val := ChangeAlertIcon{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
		}
	case "ChangeNetmapSortMethod":
		val := ChangeNetmapSortMethod{}
		if err := d.DecodeElement(&val, &start); err != nil {
			return err
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
