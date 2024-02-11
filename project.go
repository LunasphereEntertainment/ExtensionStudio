package main

import (
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	currentProject *Project
)

type ResourcePath string

type Project struct {
	Info       *hacknet.ExtensionInfo `json:"info"`
	ActionSets map[ResourcePath]hacknet.ConditionalActionSet
	Factions   map[ResourcePath]hacknet.Faction
	Nodes      map[ResourcePath]hacknet.Computer
	Missions   map[ResourcePath]hacknet.Mission
	Themes     map[ResourcePath]hacknet.Theme
}

func LoadProject(baseDir string) (*Project, error) {
	proj := new(Project)

	var err error
	proj.Info, err = LoadXML[hacknet.ExtensionInfo](path.Join(baseDir, "ExtensionInfo.xml"))
	if err != nil {
		return proj, err
	}

	proj.ActionSets, err = resourceDiscovery[hacknet.ConditionalActionSet](path.Join(baseDir, "Actions"))
	if err != nil {
		return proj, err
	}

	proj.Factions, err = resourceDiscovery[hacknet.Faction](path.Join(baseDir, "Factions"))
	if err != nil {
		return proj, err
	}

	proj.Nodes, err = resourceDiscovery[hacknet.Computer](path.Join(baseDir, "Nodes"))
	if err != nil {
		return proj, err
	}

	proj.Missions, err = resourceDiscovery[hacknet.Mission](path.Join(baseDir, "Missions"))
	if err != nil {
		return proj, err
	}

	proj.Themes, err = resourceDiscovery[hacknet.Theme](path.Join(baseDir, "Themes"))
	if err != nil {
		return proj, err
	}

	return proj, err
}

func resourceDiscovery[T interface{}](startPath string) (map[ResourcePath]T, error) {
	resources := make(map[ResourcePath]T)

	err := filepath.WalkDir(startPath, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(filePath) == ".xml" {
			filePath = strings.TrimPrefix(filePath, startPath)

			out, err := LoadXML[T](filePath)
			if err != nil {
				return fmt.Errorf("failed to parse '%s' - err: %v", filePath, err)
			}
			resources[ResourcePath(filePath)] = *out
		}

		return nil
	})

	return resources, err
}
