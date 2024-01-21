package actions

type StartScreenBleedEffect struct {
	*DelayableAction
	AlertTitle     string  `xml:"AlertTitle,attr"`
	CompleteAction string  `xml:"CompleteAction,attr"`
	Duration       float64 `xml:"TotalDurationSeconds,attr"`
}

func (a StartScreenBleedEffect) TagName() string {
	return "StartScreenBleedEffect"
}

type CancelScreenBleedEffect struct {
	*DelayableAction
}

func (a CancelScreenBleedEffect) TagName() string {
	return "CancelScreenBleedEffect"
}
