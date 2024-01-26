package templates

import (
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
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

	data, err := ExecuteTemplate(comp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
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

	data, err := ExecuteTemplate(ext)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
