package actions

type LaunchHackScript struct {
	*DelayableAction
	*TargetableAction
	Script       string `xml:"Filepath,attr"`
	Source       string `xml:"SourceComp,attr"`
	RequireLogs  bool   `xml:"RequireLogsOnSource,attr"`
	SourceIntact bool   `xml:"RequireSourceIntact,attr"`
}

func (a LaunchHackScript) TagName() string {
	return "LaunchHackScript"
}

type KillExe struct {
	*DelayableAction
	Name string `xml:"ExeName,attr"`
}

func (a KillExe) TagName() string {
	return "KillExe"
}

type HideNode struct {
	*DelayableAction
	*TargetableAction
}

func (a HideNode) TagName() string {
	return "HideNode"
}

type GivePlayerUserAccount struct {
	*DelayableAction
	*TargetableAction
	Username string `xml:"Username,attr"`
}

func (a GivePlayerUserAccount) TagName() string {
	return "GivePlayerUserAccount"
}

type ChangeIP struct {
	*DelayableAction
	*TargetableAction
	NewIP string `xml:"NewIP,attr"`
}

func (a ChangeIP) TagName() string {
	return "ChangeIP"
}

type HideAllNodes struct {
	*DelayableAction
}

func (a HideAllNodes) TagName() string {
	return "HideAllNodes"
}

type ShowNode struct {
	*HideNode
}

func (a ShowNode) TagName() string {
	return "ShowNode"
}

type TerminalModule string

const (
	Terminal TerminalModule = "terminal"
	Ram      TerminalModule = "ram"
	Netmap   TerminalModule = "netmap"
	Display  TerminalModule = "display"
)

type SetLock struct {
	*DelayableAction
	Module TerminalModule `xml:"Module,attr"`
	Locked bool           `xml:"IsLocked,attr"`
	Hidden bool           `xml:"IsHidden,attr"`
}

func (a SetLock) TagName() string {
	return "SetLock"
}

type AlertIconType string

const (
	Mail AlertIconType = "mail"
	IRC  AlertIconType = "irc"
)

type ChangeAlertIcon struct {
	*DelayableAction
	Target string        `xml:"Target,attr"`
	Type   AlertIconType `xml:"Type,attr"`
}

func (a ChangeAlertIcon) TagName() string {
	return "ChangeAlertIcon"
}

type NetmapSortMethod string

const (
	Scatter      NetmapSortMethod = "scatter"
	Grid         NetmapSortMethod = "grid"
	SequenceGrid NetmapSortMethod = "seqgrid"
	Chaos        NetmapSortMethod = "CHAOS"
)

type ChangeNetmapSortMethod struct {
	*DelayableAction
	Method NetmapSortMethod `xml:"Method,attr"`
}
