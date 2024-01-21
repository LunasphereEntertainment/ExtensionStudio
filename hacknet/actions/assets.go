package actions

type AddAsset struct {
	Filename  string `xml:"FileName,attr"`
	Content   string `xml:"FileContents,attr"`
	Target    string `xml:"TargetComp,attr"`
	TargetDir string `xml:"TargetFolderpath,attr"`
}

func (a AddAsset) TagName() string {
	return "AddAsset"
}

type CopyAsset struct {
	SourceComputer      string `xml:"SourceComp,attr"`
	SourcePath          string `xml:"SourceFilePath,attr"`
	SourceName          string `xml:"SourceFileName,attr"`
	DestinationComputer string `xml:"DestComp,attr"`
	DestinationPath     string `xml:"DestFilePath,attr"`
}

func (a CopyAsset) TagName() string {
	return "CopyAsset"
}

type DeleteFile struct {
	*DelayableAction
	*TargetableAction
	FilePath string `xml:"FilePath,attr"`
	FileName string `xml:"FileName,attr"`
}

func (a DeleteFile) TagName() string {
	return "DeleteFile"
}

type AppendToFile struct {
	*DelayableAction
	*TargetableAction
	TargetDir  string `xml:"TargetFolderpath,attr"`
	TargetName string `xml:"TargetFilename,attr"`
	Content    string `xml:",chardata"`
}

func (a AppendToFile) TagName() string {
	return "AppendToFile"
}
