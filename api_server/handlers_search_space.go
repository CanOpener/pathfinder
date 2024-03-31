package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func indexSearchSpaces(writer http.ResponseWriter, request *http.Request) {
	searchSpaces, err := mainPersistenceManager.ListSearchSpaces()
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success      bool                    `json:"success"`
		SearchSpaces []SearchSpaceIdentifier `json:"search_spaces"`
	}{
		Success:      true,
		SearchSpaces: searchSpaces,
	})
}

func createSearchSpace(writer http.ResponseWriter, request *http.Request) {
	parameters, err := createSearchSpaceParametersFromRequest(request)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	searchSpaceId, err := mainPersistenceManager.CreateSearchSpace(parameters)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success       bool   `json:"success"`
		SearchSpaceId string `json:"search_space_id"`
	}{
		Success:       true,
		SearchSpaceId: searchSpaceId,
	})
}

func showSearchSpace(writer http.ResponseWriter, request *http.Request) {
	searchSpaceId := mux.Vars(request)["id"]
	searchSpace, err := mainPersistenceManager.GetSearchSpace(searchSpaceId)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success     bool                        `json:"success"`
		SearchSpace CreateSearchSpaceParameters `json:"search_space"`
	}{
		Success:     true,
		SearchSpace: searchSpace,
	})
}

func deleteSearchSpace(writer http.ResponseWriter, request *http.Request) {
	searchSpaceId := mux.Vars(request)["id"]
	err := mainPersistenceManager.DeleteSearchSpace(searchSpaceId)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success bool `json:"success"`
	}{
		Success: true,
	})
}
