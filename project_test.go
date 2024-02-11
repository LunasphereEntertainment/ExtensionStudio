package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadProject(t *testing.T) {
	proj, err := LoadProject("C:\\Program Files (x86)\\Steam\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, proj.Info.Title, "Intro Extension")
	assert.Len(t, proj.ActionSets, 7)
	assert.Len(t, proj.Nodes, 10)
	assert.Len(t, proj.Missions, 29)
	assert.Len(t, proj.Factions, 2)
}

/*func TestNewProject(t *testing.T) {
	dest := path.Join(os.TempDir(), "Test_New_Extension")

	extInfo := &hacknet.ExtensionInfo{
		Title:              "Test New Extension",
		AllowSaves:         true,
		Language:           hacknet.English,
		StartingNodes:      hacknet.CSVList{"myTestComp"},
		StartingMission:    "Missions/StartingMission.xml",
		StartingActions:    "NONE",
		Description:        "This is a test extension! I shouldn't still be here.",
		StartsWithTutorial: false,
		HasIntroStartup:    true,
		StartingTheme:      "Hacknet_Blue",
		IntroStartupSong:   "Whiplash",
	}

	proj, err := NewProject(*extInfo, dest)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, proj.Info)
	assert.Equal(t, "Test New Extension", proj.Info.Title)

	// check the new directories exists
	if _, err := os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project directory correctly")
		}
	}
	if _, err := os.Stat(path.Join(dest, "Actions")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project structure correctly [actions]")
		}
	}
	if _, err := os.Stat(path.Join(dest, "Factions")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project structure correctly [factions]")
		}
	}
	if _, err := os.Stat(path.Join(dest, "Missions")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project structure correctly [missions]")
		}
	}
	if _, err := os.Stat(path.Join(dest, "Nodes")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project structure correctly [missions]")
		}
	}
	if _, err := os.Stat(path.Join(dest, "Themes")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create project structure correctly [themes]")
		}
	}

	// check extension info exists
	if _, err := os.Stat(path.Join(dest, "ExtensionInfo.xml")); err != nil {
		if os.IsNotExist(err) {
			t.Fatal("did not create extension info correctly")
		}
	}

	// check we can't accidentally overwrite the project
	_, err = NewProject(*extInfo, dest)
	if err == nil || err.Error() != "project directory already exists" {
		t.Fatal("we shouldn't be overriding projects")
	}
}

func TestDeleteProject(t *testing.T) {
	dest := path.Join(os.TempDir(), "Test_New_Extension")

	err := DeleteProject(dest)
	if err != nil {
		t.Fatal(err)
	}
}
*/
