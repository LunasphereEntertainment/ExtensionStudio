package actions

type RunFunction struct {
	Name  string `xml:"FunctionName,attr"`
	Value string `xml:"FunctionValue,attr"`
}

func (rf RunFunction) TagName() string {
	return "RunFunction"
}

type CrashComputer struct {
	*DelayableAction
	*TargetableAction
	Source string `xml:"CrashSource,attr"`
}

func (c CrashComputer) TagName() string {
	return "CrashComputer"
}

type AddConditionalActions struct {
	*DelayableAction
	Path string `xml:"FilePath,attr"`
}

func (a AddConditionalActions) TagName() string {
	return "AddConditionalActions"
}

type SaveGame struct {
	*DelayableAction
}

func (a SaveGame) TagName() string {
	return "SaveGame"
}
