package model

type GoalType string

const (
	FileDeletion         GoalType = "filedeletion"
	ClearFolder          GoalType = "clearfolder"
	FileDownload         GoalType = "filedownload"
	FileChange           GoalType = "filechange"
	GetAdmin             GoalType = "getadmin"
	GetString            GoalType = "getstring"
	Delay                GoalType = "delay"
	HasFlag              GoalType = "hasflag"
	FileUpload           GoalType = "fileupload"
	AddDegree            GoalType = "AddDegree"
	WipeDegrees          GoalType = "wipedegrees"
	SendEmail            GoalType = "sendemail"
	RemoveDeathRowRecord GoalType = "removeDeathRowRecord"
	ModifyDeathRowRecord GoalType = "modifyDeathRowRecord"
	GetAdminPass         GoalType = "getadminpasswordstring"
)

type Goal struct {
	Type          GoalType `xml:"type,attr" json:"type"`
	Target        *string  `xml:"target,attr" json:"target,omitempty"`
	DestTarget    *string  `xml:"destTarget,attr" json:"destTarget,omitempty"`
	File          *string  `xml:"file,attr" json:"file,omitempty"`
	Path          *string  `xml:"path,attr" json:"path,omitempty"`
	DestPath      *string  `xml:"destPath,attr" json:"destPath,omitempty"`
	Keyword       *string  `xml:"keyword,attr" json:"keyword,omitempty"`
	Removal       *string  `xml:"removal,attr" json:"removal,omitempty"`
	CaseSensitive *string  `xml:"caseSensitive,attr" json:"caseSensitive,omitempty"`
	Time          *string  `xml:"time,attr" json:"time,omitempty"`
	Decrypt       *bool    `xml:"decrypt,attr" json:"decrypt,omitempty"`
	DecryptPass   *string  `xml:"decryptPass,attr" json:"decryptPass,omitempty"`
	MailServer    *string  `xml:"mailServer,attr" json:"mailServer,omitempty"`
	Recipient     *string  `xml:"recipient,attr" json:"recipient,omitempty"`
	Subject       *string  `xml:"subject,attr" json:"subject,omitempty"`
	Owner         *string  `xml:"owner,attr" json:"owner,omitempty"`
}
