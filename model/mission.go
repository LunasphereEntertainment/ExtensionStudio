package model

import "encoding/xml"

type MissionFunctions struct {
	Value         string `xml:",chardata" json:"value"`
	Suppress      *bool  `xml:"suppress,attr,omitempty" json:"suppress,omitempty"`
	FunctionValue *int   `xml:"val,attr,omitempty" json:"functionValue,omitempty"`
}

type NextMission struct {
	Path   ExternalReference[Mission] `xml:",chardata" json:"path"`
	Silent bool                       `xml:"IsSilent" json:"silent"`
}

type Mission struct {
	xml.Name     `xml:"mission" json:"-"`
	ID           string                       `xml:"id,attr" json:"id"`
	ActiveCheck  bool                         `xml:"activeCheck,attr" json:"activeCheck"`
	VerifySender bool                         `xml:"shouldIgnoreSenderVerification,attr" json:"verifySender"`
	Goals        []Goal                       `xml:"goals>goal" json:"goals"`
	Start        MissionFunctions             `xml:"missionStart"`
	End          MissionFunctions             `xml:"missionEnd"`
	NextMission  NextMission                  `xml:"nextMission" json:"nextMission"`
	Branches     []ExternalReference[Mission] `xml:"branchMissions>branch"`
	Posting      BoardPosting                 `xml:"posting"`
	Email        `xml:"email" json:"email"`
}

type BoardPosting struct {
	Title         string  `xml:"title,attr" json:"title"`
	RequiredFlags CSVList `xml:"reqs,attr" json:"requiredFlags"`
	RequiredRank  int     `xml:"requiredRank,attr" json:"requiredRank"`
	Body          string  `xml:",chardata"`
}
