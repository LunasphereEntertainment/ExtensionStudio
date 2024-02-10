package main

import (
	"encoding/json"
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Project struct {
	Path string `json:"-"`

	Info       *hacknet.ExtensionInfo                                    `json:"info"`
	ActionSets []hacknet.ExternalReference[hacknet.ConditionalActionSet] `json:"actions"`
	Factions   []hacknet.ExternalReference[hacknet.Faction]              `json:"factions"`
	Missions   []hacknet.ExternalReference[hacknet.Mission]              `json:"missions"`
	Nodes      []hacknet.ExternalReference[nodes.Computer]               `json:"nodes"`
	Themes     []hacknet.ExternalReference[hacknet.Theme]                `json:"themes"`
}

type ProjectListing struct {
	ID   uuid.UUID `json:"id"`
	Path string    `json:"path"`
}

var (
	recentProjects projectListing
)

type projectListing []ProjectListing

func getApplicationDir() string {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := path.Join(cfgPath, "ExtensionStudio")
	err = os.MkdirAll(appDir, 0644)
	if err != nil {
		panic(err)
	}

	return appDir
}

func (pl projectListing) Find(id uuid.UUID) (*ProjectListing, error) {
	for _, proj := range pl {
		if proj.ID == id {
			return &proj, nil
		}
	}

	return nil, fmt.Errorf("recent project does not exist")
}

func (pl projectListing) FindByPath(path string) (*ProjectListing, error) {
	for _, proj := range pl {
		if proj.Path == path {
			return &proj, nil
		}
	}

	return nil, fmt.Errorf("recent project does not exist")
}

func (pl projectListing) Save() {
	f, err := os.OpenFile(path.Join(getApplicationDir(), "projects.json"), os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(f).Encode(pl)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func init() {
	recentProjects = make(projectListing, 0)

	f, err := os.OpenFile(path.Join(getApplicationDir(), "projects.json"), os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		if os.IsNotExist(err) {
			recentProjects.Save()
			return
		} else {
			panic(err)
		}
	}

	err = json.NewDecoder(f).Decode(&recentProjects)
	if err != nil {
		log.Println(err)
	}
}

func resourceDiscovery[T interface{}](projectId uuid.UUID, startPath string) []hacknet.ExternalReference[T] {
	resources := make([]hacknet.ExternalReference[T], 0)

	err := filepath.WalkDir(startPath, func(filePath string, d os.DirEntry, err error) error {
		if filepath.Ext(filePath) == ".xml" {
			filePath = strings.TrimPrefix(filePath, startPath)
			resources = append(resources, hacknet.ExternalReference[T]{Path: filePath, ProjectID: projectId})
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

	proj.Info, err = LoadXML[hacknet.ExtensionInfo]{path.Join(projectPath, "ExtensionInfo.xml")}
	proj.ActionSets = resourceDiscovery[hacknet.ConditionalActionSet]{path.Join(projectPath, "Actions")}
	proj.Factions = resourceDiscovery[hacknet.Faction]{path.Join(projectPath, "Factions")}
	proj.Missions = resourceDiscovery[hacknet.Mission]{path.Join(projectPath, "Missions")}
	proj.Nodes = resourceDiscovery[nodes.Computer]{path.Join(projectPath, "Nodes")}
	proj.Themes = resourceDiscovery[hacknet.Theme]{path.Join(projectPath, "Themes")}

	_, err = recentProjects.FindByPath(projectPath)
	if err != nil {
		recentProjects = append(recentProjects, ProjectListing{ID: uuid.New(), Path: projectPath})
		recentProjects.Save()
	}

	return proj, err
}

func LoadRecentProject(id uuid.UUID) (*Project, error) {
	proj, err := recentProjects.Find(id)
	if err != nil {
		return nil, err
	}

	return LoadProject(proj.Path)
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

	// Create directories (if not exist)
	err = os.MkdirAll(path.Join(projectPath, "Actions"), 0644)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(path.Join(projectPath, "Missions"), 0644)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(path.Join(projectPath, "Nodes"), 0644)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(path.Join(projectPath, "Factions"), 0644)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(path.Join(projectPath, "Themes"), 0644)
	if err != nil {
		return nil, err
	}

	recentProjects = append(recentProjects, ProjectListing{
		ID:   uuid.New(),
		Path: projectPath,
	})
	recentProjects.Save()

	return LoadProject(projectPath)
}

// DeleteProject deletes a project directory and all it's contents.
// A/N: used to be HIGHLY DESTRUCTIVE USE WITH CAUTION
func DeleteProject(projectPath string) error {
	//err := os.RemoveAll(projectPath)
	//if err != nil {
	//	return err
	//}

	var i = -1
	for _, p := range recentProjects {
		if p.Path == projectPath {
			break
		}
		i++
	}

	if i >= 0 {
		recentProjects = append(recentProjects[0:i-1], recentProjects[i+1:]...)
	}

	recentProjects.Save()

	return nil
}
