package main

import (
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	actions2 "github.com/LunasphereEntertainment/ExtensionStudio/hacknet/actions"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	basePath = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension"
)

func TestLoadExtension(t *testing.T) {
	ext, err := LoadXML[hacknet.ExtensionInfo](basePath + "\\ExtensionInfo.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, hacknet.English, ext.Language)
	assert.Equal(t, "Intro Extension", ext.Title)
	assert.True(t, ext.AllowSaves)

	assert.Len(t, ext.StartingNodes, 1)
	assert.Equal(t, "advExamplePC", ext.StartingNodes[0])

	assert.Equal(t, hacknet.ExternalReference[hacknet.Mission]("Missions/Intro/IntroMission1.xml"), ext.StartingMission)
	assert.Equal(t, hacknet.ExternalReference[hacknet.ConditionalActionSet]("Actions/StartingActions.xml"), ext.StartingActions)
}

func TestLoadMission(t *testing.T) {
	mission, err := LoadXML[hacknet.Mission](basePath + "\\Missions\\ExampleMission.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, "testMission0", mission.ID)
	assert.Equal(t, true, mission.ActiveCheck)
	assert.Equal(t, false, mission.VerifySender)

	// Goals
	assert.Len(t, mission.Goals, 15)

	// StartMission
	assert.Equal(t, "changeSong", mission.Start.Value)
	assert.Equal(t, 7, *mission.Start.FunctionValue)
	assert.Equal(t, true, *mission.Start.Suppress)

	// EndMission
	assert.Equal(t, "addRank", mission.End.Value)
	assert.Equal(t, 1, *mission.End.FunctionValue)
	assert.Nil(t, mission.End.Suppress)

	// NextMission
	assert.Equal(t, false, mission.NextMission.Silent)
	assert.Equal(t, hacknet.ExternalReference[hacknet.Mission]("NONE"), mission.NextMission.Path)

	// Branches
	assert.Len(t, mission.Branches, 1)
	assert.Equal(t, hacknet.ExternalReference[hacknet.Mission]("Missions/BranchExample/TestBranchMission.xml"), mission.Branches[0])

	// Email
	assert.Equal(t, "Matt", mission.Email.Sender)
	assert.Equal(t, "Test Mission Email", mission.Email.Subject)
	// Email - Note Attachments
	assert.Len(t, mission.Email.Notes, 1)
	assert.Equal(t, "An example note", mission.Email.Notes[0].Title)

	assert.Len(t, mission.Email.Links, 1)
	assert.Len(t, mission.Email.Accounts, 1)
}

func TestLoadActionSet(t *testing.T) {
	actions, err := LoadXML[hacknet.ConditionalActionSet](basePath + "\\Actions\\ExampleConditionalActionSet.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.NotNil(t, actions.OnConnect)
	assert.Len(t, actions.OnConnect, 1)
	assert.True(t, *(actions.OnConnect[0].NeedsMissionComplete))
	assert.Len(t, actions.OnConnect[0].RequiredFlags, 1)
	assert.Equal(t, "decypher", actions.OnConnect[0].RequiredFlags[0])
	assert.IsType(t, actions2.AddIRCMessage{}, actions.OnConnect[0].Actions[0].Value)

	assert.NotNil(t, actions.HasFlags)
	assert.Len(t, actions.HasFlags, 1)
	assert.Len(t, actions.HasFlags[0].RequiredFlags, 2)
	assert.Equal(t, "decypher", actions.HasFlags[0].RequiredFlags[0])
	assert.Equal(t, "otherFlag", actions.HasFlags[0].RequiredFlags[1])

	assert.NotNil(t, actions.OnAdminGained)
	assert.Len(t, actions.OnAdminGained, 1)
	assert.NotNil(t, actions.OnAdminGained[0].Target)
	assert.Equal(t, "advExamplePC", *actions.OnAdminGained[0].Target)

	assert.NotNil(t, actions.Instantly)

	assert.NotNil(t, actions.DoesNotHaveFlags)
	assert.Len(t, actions.DoesNotHaveFlags, 1)
	assert.Len(t, actions.DoesNotHaveFlags[0].Flags, 2)
	assert.Equal(t, "SomeFlag", actions.DoesNotHaveFlags[0].Flags[0])
	assert.Equal(t, "MoreFlags", actions.DoesNotHaveFlags[0].Flags[1])

	assert.NotNil(t, actions.OnDisconnect)
}

func TestLoadTestLoadActionSet2(t *testing.T) {
	actions, err := LoadXML[hacknet.ConditionalActionSet](basePath + "\\Actions\\HackerScriptActions.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Len(t, actions.OnConnect, 2)
	assert.Len(t, actions.OnConnect[0].Actions, 1)
	assert.IsType(t, actions2.LaunchHackScript{}, actions.OnConnect[0].Actions[0].Value)
	hackScript := actions.OnConnect[0].Actions[0].Value.(actions2.LaunchHackScript)
	assert.Equal(t, "HackerScripts/AllyHack.txt", hackScript.Script)
	assert.Equal(t, "advExamplePC", *hackScript.DelayHost)
	assert.Equal(t, 2.5, hackScript.Delay)
	assert.Equal(t, "allyHackerSource", hackScript.Source)
	assert.Equal(t, "hackerTarget", hackScript.Target)
	assert.False(t, hackScript.RequireLogs)
}

func TestLoadFaction(t *testing.T) {
	faction, err := LoadXML[hacknet.Faction](basePath + "\\Factions\\ExampleFaction.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, "Example Faction", faction.FactionName)
	assert.Equal(t, "examplefaction", faction.ID)
	assert.Equal(t, 0, faction.StartValue)
	assert.Len(t, faction.ActionSets, 3)
	assert.Equal(t, *faction.ActionSets[0].RequiresValue, 1)

	assert.IsType(t, actions2.RunFunction{}, faction.ActionSets[0].Actions[0].Value)
	assert.IsType(t, actions2.LoadMission{}, faction.ActionSets[0].Actions[1].Value)
	assert.IsType(t, actions2.AddAsset{}, faction.ActionSets[0].Actions[2].Value)
}

func TestLoadTheme(t *testing.T) {
	theme, err := LoadXML[hacknet.Theme](basePath + "\\Themes\\ExampleTheme.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, hacknet.Mint, theme.Layout)
	assert.Equal(t, "Themes/Backgrounds/RiptideGreen.png", theme.BackgroundPath)
	assert.False(t, theme.PreserveScaling)
	assert.Len(t, theme.BackgroundFill, 3)
	assert.Equal(t, hacknet.RgbColour{0, 0, 0}, theme.BackgroundFill)
	assert.Equal(t, hacknet.RgbColour{255, 41, 63}, theme.DefaultHighlightColor)
	assert.Len(t, theme.DefaultTopBarColor, 4)
	assert.Equal(t, hacknet.RgbColour{74, 7, 14, 255}, theme.DefaultTopBarColor)

	assert.Equal(t, hacknet.RgbColour{255, 0, 0}, theme.WarningColor)
	assert.Equal(t, hacknet.RgbColour{90, 90, 90}, theme.SubtleTextColor)
	assert.Equal(t, hacknet.RgbColour{8, 8, 8}, theme.DarkBackgroundColor)
	assert.Equal(t, hacknet.RgbColour{12, 12, 12}, theme.IndentBackgroundColor)
	assert.Equal(t, hacknet.RgbColour{68, 68, 68}, theme.OutlineColor)
	assert.Equal(t, hacknet.RgbColour{65, 16, 16, 200}, theme.LockedColor)
	assert.Equal(t, hacknet.RgbColour{160, 0, 0}, theme.BrightLockedColor)
	assert.Equal(t, hacknet.RgbColour{0, 160, 0}, theme.BrightUnlockedColor)
	assert.Equal(t, hacknet.RgbColour{39, 65, 36}, theme.UnlockedColor)
	assert.Equal(t, hacknet.RgbColour{180, 180, 180}, theme.LightGray)
	assert.Equal(t, hacknet.RgbColour{222, 201, 24}, theme.ShellColor)
	assert.Equal(t, hacknet.RgbColour{105, 167, 188}, theme.ShellButtonColor)
	assert.Equal(t, hacknet.RgbColour{120, 120, 120, 0}, theme.SemiTransText)
	assert.Equal(t, hacknet.RgbColour{213, 245, 255}, theme.TerminalTextColor)
	assert.Equal(t, hacknet.RgbColour{126, 126, 126, 100}, theme.TopBarTextColor)
	assert.Equal(t, hacknet.RgbColour{2, 2, 2, 30}, theme.SuperLightWhite)
	assert.Equal(t, hacknet.RgbColour{222, 0, 0, 195}, theme.ConnectedNodeHighlight)
	assert.Equal(t, hacknet.RgbColour{213, 245, 255, 0}, theme.NetmapToolTipColor)
	assert.Equal(t, hacknet.RgbColour{0, 0, 0, 70}, theme.NetmapToolTipBackground)
	assert.Equal(t, hacknet.RgbColour{255, 255, 255}, theme.TopBarIconsColor)
	assert.Equal(t, hacknet.RgbColour{95, 220, 83}, theme.ThisComputerNode)
	assert.Equal(t, hacknet.RgbColour{255, 255, 255, 15}, theme.ScanlinesColor)
}

func TestLoadComputer(t *testing.T) {
	node, err := LoadXML[hacknet.Computer](basePath + "\\Nodes\\ExampleComputer.xml")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, "advExamplePC", node.ID)
	assert.Equal(t, "Extension Example PC", node.Name)
	assert.Equal(t, "167.194.132.7", node.IP)
	assert.Equal(t, 2, node.Security)
	assert.False(t, node.AllowsDefaultBootModule)
	assert.Equal(t, hacknet.Chip, node.Icon)
	assert.Equal(t, 1, node.Type)
	assert.Equal(t, hacknet.AdminPass{Password: "password"}, node.AdminPass)
	assert.Len(t, node.Ports, 10)
	assert.Equal(t, "21", node.Ports[0])
	assert.Equal(t, "554", node.Ports[9])
	assert.Len(t, node.PortRemap, 2)
	assert.Contains(t, node.PortRemap, "web")
	assert.Equal(t, 1234, node.PortRemap["web"])
	assert.Contains(t, node.PortRemap, "22")
	assert.Equal(t, 2, node.PortRemap["22"])
	assert.True(t, bool(node.HasTracker))
}
