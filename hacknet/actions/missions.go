package actions

type LoadMission struct {
	Path string `xml:"MissionName,attr"`
}

func (a LoadMission) TagName() string {
	return "LoadMission"
}

type AddMissionToHubServer struct {
	*TargetableAction
	Path       string  `xml:"MissionFilepath,attr"`
	Assignment *string `xml:"AssignmentTag,attr,omitempty"`
}

func (am AddMissionToHubServer) TagName() string {
	return "AddMissionToHubServer"
}

type RemoveMissionFromHubServer struct {
	*TargetableAction
	Path string `xml:"MissionFilepath,attr"`
}

func (rm RemoveMissionFromHubServer) TagName() string {
	return "RemoveMissionFromHubServer"
}

type AddThreadToMissionBoard struct {
	*TargetableAction
	Path string `xml:"ThreadFilepath,attr"`
}

func (ad AddThreadToMissionBoard) TagName() string {
	return "AddThreadToMissionBoard"
}
