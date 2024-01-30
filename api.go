package main

import (
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func RestApi(r *mux.Router) {
	r.PathPrefix("/projects").Methods(http.MethodGet).HandlerFunc(listProjects)
	r.PathPrefix("/projects").Methods(http.MethodPost).HandlerFunc(newProject)
	r.PathPrefix("/projects/{id}").Methods(http.MethodPut).HandlerFunc(openProject)
	r.PathPrefix("/projects/{id}").Methods(http.MethodDelete).HandlerFunc(deleteProject)
	r.PathPrefix("/info").Methods(http.MethodGet).HandlerFunc(loadResource[hacknet.ExtensionInfo])
	r.PathPrefix("/info").Methods(http.MethodPost).HandlerFunc(saveResource[hacknet.ExtensionInfo])
	r.PathPrefix("/actions").Methods(http.MethodGet).HandlerFunc(loadResource[hacknet.ConditionalActionSet])
	r.PathPrefix("/actions").Methods(http.MethodPost).HandlerFunc(saveResource[hacknet.ConditionalActionSet])
	r.PathPrefix("/nodes").Methods(http.MethodGet).HandlerFunc(loadResource[nodes.Computer])
	r.PathPrefix("/nodes").Methods(http.MethodPost).HandlerFunc(saveResource[nodes.Computer])
	r.PathPrefix("/factions").Methods(http.MethodGet).HandlerFunc(loadResource[hacknet.Faction])
	r.PathPrefix("/factions").Methods(http.MethodPost).HandlerFunc(saveResource[hacknet.Faction])
	r.PathPrefix("/missions").Methods(http.MethodGet).HandlerFunc(loadResource[hacknet.Mission])
	r.PathPrefix("/missions").Methods(http.MethodPost).HandlerFunc(saveResource[hacknet.Mission])
	r.PathPrefix("/themes").Methods(http.MethodGet).HandlerFunc(loadResource[hacknet.Theme])
	r.PathPrefix("/themes").Methods(http.MethodPost).HandlerFunc(saveResource[hacknet.Theme])
}

func listProjects(w http.ResponseWriter, r *http.Request) {
	serialize(w, recentProjects)
}

func openProject(w http.ResponseWriter, r *http.Request) {
	projectId, err := uuid.Parse(mux.Vars(r)["id"])
	// when none/invalid projectID specified
	if projectId == uuid.Nil || err != nil {
		// try loading a specified directory instead
		openProjectDir(w, r)
		return
	}

	// otherwise, continue to load recent project
	openRecentProject(w, r, projectId)
}

func openRecentProject(w http.ResponseWriter, r *http.Request, projectId uuid.UUID) {
	project, err := LoadRecentProject(projectId)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	serialize(w, project)
}

func openProjectDir(w http.ResponseWriter, r *http.Request) {
	listing := deserialize[ProjectListing](w, r)
	if listing == nil {
		return
	}

	if len(listing.Path) == 0 {
		httpError(w, http.StatusBadRequest, fmt.Errorf("invalid path/no path specified"))
		return
	}

	proj, err := LoadProject(listing.Path)
	if err != nil {
		panic(err)
	}

	serialize(w, proj)
}

type newProjectRequest struct {
	Path string                `json:"path"`
	Info hacknet.ExtensionInfo `json:"info"`
}

func newProject(w http.ResponseWriter, r *http.Request) {
	req := deserialize[newProjectRequest](w, r)
	if req == nil {
		httpError(w, http.StatusBadRequest, fmt.Errorf("invalid project creation object"))
		return
	}

	proj, err := NewProject(req.Info, req.Path)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	serialize(w, proj)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	projectId, err := uuid.Parse(mux.Vars(r)["id"])

	// when none/invalid projectID specified
	if projectId == uuid.Nil || err != nil {
		httpError(w, http.StatusBadRequest, fmt.Errorf("no project id/uuid specified"))
		return
	}

	listing, err := recentProjects.Find(projectId)
	if err != nil {
		httpError(w, http.StatusNotFound, err)
	}

	DeleteProject(listing.Path)

}

func loadResource[T interface{}](w http.ResponseWriter, r *http.Request) {
	resourcePath, _ := url.QueryUnescape(r.URL.Query().Get("path"))
	if len(resourcePath) == 0 {
		httpError(w, http.StatusNotFound, fmt.Errorf("resource could not be found"))
		return
	}

	res, err := LoadXML[T](resourcePath)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
	}
	serialize(w, res)
}

func saveResource[T interface{}](w http.ResponseWriter, r *http.Request) {
	resourcePath, _ := url.QueryUnescape(r.URL.Query().Get("path"))
	if len(resourcePath) == 0 {
		httpError(w, http.StatusNotFound, fmt.Errorf("resource could not be found"))
		return
	}

	req := deserialize[T](w, r)
	if req == nil {
		return
	}
	//req := new(T)
	//err := json.NewDecoder(r.Body).Decode(req)
	//if err != nil {
	//	httpError(w, http.StatusInternalServerError, err)
	//	return
	//}

	err := SaveXML(resourcePath, *req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}
}
