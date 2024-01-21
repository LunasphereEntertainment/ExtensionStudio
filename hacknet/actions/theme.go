package actions

type SwitchToTheme struct {
	*DelayableAction
	ThemePath       string  `xml:"ThemePath,attr"`
	FlickerDuration float64 `xml:"FlickerInDuration,attr"`
}

func (a SwitchToTheme) TagName() string {
	return "SwitchToTheme"
}
