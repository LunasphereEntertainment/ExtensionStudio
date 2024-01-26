package main

import (
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"os"
	"path"
	"path/filepath"
)

type Project struct {
	Path string

	Info       *hacknet.ExtensionInfo
	ActionSets []hacknet.ExternalReference[hacknet.ConditionalActionSet]
	Factions   []hacknet.ExternalReference[hacknet.Faction]
	Missions   []hacknet.ExternalReference[hacknet.Mission]
	Nodes      []hacknet.ExternalReference[nodes.Computer]
	Themes     []hacknet.ExternalReference[hacknet.Theme]
}

func resourceDiscovery[T interface{}](startPath string) []hacknet.ExternalReference[T] {
	resources := make([]hacknet.ExternalReference[T], 0)

	err := filepath.WalkDir(startPath, func(path string, d os.DirEntry, err error) error {
		if filepath.Ext(path) == ".xml" {
			resources = append(resources, hacknet.ExternalReference[T](path))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return resources
}

// LoadProject loads an extension project and it's resources from the given file path.
func LoadProject(projectPath string) (proj *Project, err error) {
	// Check if project directory exists, error if not
	if _, err := os.Stat(projectPath); err != nil {
		return nil, err
	}

	proj = &Project{
		Path: projectPath,
	}

	proj.Info, err = LoadXML[hacknet.ExtensionInfo](path.Join(projectPath, "ExtensionInfo.xml"))
	proj.ActionSets = resourceDiscovery[hacknet.ConditionalActionSet](path.Join(projectPath, "Actions"))
	proj.Factions = resourceDiscovery[hacknet.Faction](path.Join(projectPath, "Factions"))
	proj.Missions = resourceDiscovery[hacknet.Mission](path.Join(projectPath, "Missions"))
	proj.Nodes = resourceDiscovery[nodes.Computer](path.Join(projectPath, "Nodes"))
	proj.Themes = resourceDiscovery[hacknet.Theme](path.Join(projectPath, "Themes"))

	return proj, err
}

// NewProject constructs a new project in the specified directory, expects an extension info to also be submitted.
func NewProject(info hacknet.ExtensionInfo, projectPath string) (proj *Project, err error) {
	// Check if project directory exists, create if not.
	if _, err := os.Stat(projectPath); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(projectPath, 644)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// TODO: project already exists, confirm override
		return nil, fmt.Errorf("project directory already exists")
	}

	// Write Extension Info to file
	err = SaveXML(path.Join(projectPath, "ExtensionInfo.xml"), info)
	if err != nil {
		return nil, err
	}
	return LoadProject(projectPath)
}

// DeleteProject deletes a project directory and all it's contents. THIS IS HIGHLY DESTRUCTIVE USE WITH CAUTION
func DeleteProject(projectPath string) error {
	return os.RemoveAll(projectPath)
}
