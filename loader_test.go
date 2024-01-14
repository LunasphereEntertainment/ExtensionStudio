package main

import (
	"github.com/LunasphereEntertainment/ExtensionStudio/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMission(t *testing.T) {
	mission, err := LoadXML[model.Mission]("E:\\SteamLibrary\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension\\Missions\\ExampleMission.xml")
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
	assert.Equal(t, model.ExternalReference[model.Mission]("NONE"), mission.NextMission.Path)

	// Branches
	assert.Len(t, mission.Branches, 1)
	assert.Equal(t, model.ExternalReference[model.Mission]("Missions/BranchExample/TestBranchMission.xml"), mission.Branches[0])

	// Email
	assert.Equal(t, "Matt", mission.Email.Sender)
	assert.Equal(t, "Test Mission Email", mission.Email.Subject)
	assert.Len(t, mission.Email.Attachments, 3)
}
