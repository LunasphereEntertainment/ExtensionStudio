package templates

import (
	"bytes"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/actions"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"testing"
)

func TestExecuteNodeTemplate(t *testing.T) {
	decryptionPass := "decryptionPassword"

	comp := nodes.Computer{
		ComputerAttrs: &nodes.ComputerAttrs{
			ID:                      "myTestPC",
			Name:                    "Simple Test PC",
			IP:                      "123.456.789.123",
			Security:                1,
			AllowsDefaultBootModule: false,
			Icon:                    nodes.DLCPC1,
			Type:                    0,
		},
		AdminPass: nodes.AdminPass{
			Password: "testy test",
		},
		Ports: hacknet.CSVList{"22", "25", "443"},
		PortsForCrack: nodes.PortsForCrack{
			Value: 3,
		},
		Trace: &nodes.TraceConfig{Time: 5678},
		AdminConfig: &nodes.AdminConfig{
			Type:          "fast",
			ResetPassword: false,
			SuperUser:     false,
		},
		Accounts: []nodes.ComputerAccount{
			{
				Username: "test",
				Password: "Password123!",
				Type:     nodes.Admin,
			},
		},
		PortRemap: nodes.PortRemap{"web": 1234, "22": 2},
		Links:     []nodes.ComputerLink{{Target: "advExamplePC2"}},
		Files: []nodes.File{
			{
				FileHeaders: &nodes.FileHeaders{
					Path: "home",
					Name: "Test_File.txt",
				},
				Content: `This is a test file in the home directory.
It has multiple lines
just
for
testing.

Yay!`,
			},
			{
				FileHeaders: &nodes.FileHeaders{
					Path: "bin",
					Name: "SSHCrack.exe",
				},
				Content: "#SSH_CRACK#",
			},
		},
		CustomThemes: []nodes.CustomThemeFile{
			{
				FileHeaders: &nodes.FileHeaders{
					Path: "sys",
					Name: "Custom_x-server.sys",
				},
				ThemePath: "Themes/SecondaryTheme.xml",
			},
		},
		EncryptedFiles: []nodes.EncryptedFile{
			{
				FileHeaders: &nodes.FileHeaders{
					Path: "home",
					Name: "encrypted_File.dec",
				},
				IP:        "192.168.1.1",
				Extension: ".txt",
				Header:    "This is the header",
				Password:  &decryptionPass,
				Content:   "This generates an encrypted file that can be decrypted using the password above. It decrypts to have the extension .txt",
			},
		},
		EosLinks: []nodes.EosDevice{
			{
				ID:           "eosIntroPhone",
				Name:         "Deliliah's ePhone 4S",
				Icon:         "ePhone2",
				Empty:        true,
				PassOverride: "notAlpine",
				Notes: []nodes.EosNote{
					{Content: `TestNote
More text`},
				},
				MailAccounts: []nodes.EosMailAccount{
					{
						Username: "test@jmail.com",
						Password: "thisIstheaccountpass",
					},
				},
				Files: []nodes.File{
					{FileHeaders: &nodes.FileHeaders{
						Path: "eos/test",
						Name: "crackedFile.txt",
					}, Content: `This is mostly useful for jailbroken phones`},
				},
			},
		},
	}

	out := bytes.Buffer{}
	err := ExecuteTemplate(comp, &out)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out.String())
}

func TestExecuteExtensionTemplate(t *testing.T) {
	ext := hacknet.ExtensionInfo{
		Title:      "Test Extension 1",
		AllowSaves: false,
		Language:   hacknet.English,
		StartingNodes: hacknet.CSVList{
			"advExamplePC",
			"playerComp",
		},
		StartingMission: "",
		StartingActions: "",
		//StartingMission:    ,
		//StartingActions:    "",
		Description: "This is an example, introductory extension.",
		Factions: []hacknet.ExternalReference[hacknet.Faction]{
			"Factions/ExampleFaction.xml",
		},
		StartsWithTutorial: false,
		HasIntroStartup:    false,
		StartingTheme:      "Themes/ExampleTheme.xml",
		IntroStartupSong:   "The_Quickening",
	}

	out := bytes.Buffer{}
	err := ExecuteTemplate(ext, &out)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out.String())
}

func TestExecuteMissionTemplate(t *testing.T) {
	missionStartSuppress := true
	missionStartValue := 7

	delayGoalTime := "10.0"
	flagTarget := "flagName"

	mission := hacknet.Mission{
		ID:           "testMission0",
		ActiveCheck:  true,
		VerifySender: false,
		Goals: []hacknet.Goal{
			{
				Type: hacknet.Delay,
				Time: &delayGoalTime,
			},
			{
				Type:   hacknet.HasFlag,
				Target: &flagTarget,
			},
		},
		Start: hacknet.MissionFunctions{
			Value:         "changeSong",
			Suppress:      &missionStartSuppress,
			FunctionValue: &missionStartValue,
		},
		End: hacknet.MissionFunctions{},
		NextMission: hacknet.NextMission{
			Path:   "NONE",
			Silent: false,
		},
		Branches: []hacknet.ExternalReference[hacknet.Mission]{
			"Missions/BranchExample/TestBranchMission.xml",
		},
		Posting: hacknet.BoardPosting{
			Title:         "Do the Extension Test Mission",
			RequiredFlags: hacknet.CSVList{"someCustomFlag"},
			RequiredRank:  3,
			Body: `This is the body text of the posting that will appear when the mission is clicked on. It should contain a basic outline, with any warnings the player needs.
Once accepted, the email should contain full details.`,
		},
		Email: hacknet.Email{
			Sender:  "Matt",
			Subject: "Test Mission Email",
			Body: `This is the body of the email.
Be careful with your formatting! The Hacknet parser does not account for auto-whitespace added to the left here.
Email contents are very important - small changes in wording can dramatically change how easy or hard a mission is.
Be very conscious about how much you are hinting at and guiding the player, and understand that it will be much harder
for everyone that's not you.

Good luck,
-Matt
    `,
			Notes: []hacknet.NoteAttachment{
				{Title: "An example note", Content: `Experiment with note formatting!
Remember that the note space is very small, so text overflow onto new lines happens quickly.

Guide the player! Hacknet missions are very frustrating when the player has too little direction and cant continue.
      `},
			},
			Links: []hacknet.LinkAttachment{
				{Computer: "missionTestNode"},
			},
			Accounts: []hacknet.AccountAttachment{
				{
					Computer: "missionTestNode",
					Username: "TestUser",
					Password: "testpass",
				},
			},
		},
	}

	out := bytes.Buffer{}
	err := ExecuteTemplate(mission, &out)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out.String())
}

func TestExecuteActionSetTemplate(t *testing.T) {
	actionSet := hacknet.ConditionalActionSet{
		OnConnect: []hacknet.ActionSequence{
			{Actions: []actions.Action{
				{Value: actions.AddIRCMessage{
					DelayableAction: &actions.DelayableAction{
						Delay: 5,
					},
					TargetableAction: &actions.TargetableAction{Target: "advExamplePC"},
					Author:           "DependableSkeleton",
					Content:          "Hey, you're back, having just completed your mission.",
				}},
			}},
		},
	}

	out := bytes.Buffer{}
	err := ExecuteTemplate(actionSet, &out)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out.String())
}
