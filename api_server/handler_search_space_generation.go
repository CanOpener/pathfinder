package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/canopener/pathfinder/api_server/searchSpaceGeneration"
)

type searchSpaceDescriptor interface{}
type searchSpaceParameters struct {
	algorithm     string
	nameGenerator string
	nodeCount     int
}

func (p searchSpaceParameters) Algorithm() string     { return p.algorithm }
func (p searchSpaceParameters) NameGenerator() string { return p.nameGenerator }
func (p searchSpaceParameters) NodeCount() int        { return p.nodeCount }

var validAlgorithms = map[string]bool{
	"default": true,
	"prim":    true,
}

var validNameGenerators = map[string]bool{
	"default":  true,
	"3letters": true,
}

func newSearchSpaceParameters(request *http.Request) (searchSpaceParameters, error) {
	queryParams := request.URL.Query()
	algorithm := queryParams.Get("algorithm")
	_, ok := validAlgorithms[algorithm]
	if !ok {
		return searchSpaceParameters{},
			fmt.Errorf("invalid algorithm %s.\nvalid options are: %s", algorithm, serializeStringSet(validAlgorithms))
	}

	nameGenerator := request.URL.Query().Get("name_generator")
	_, ok = validNameGenerators[nameGenerator]
	if !ok {
		return searchSpaceParameters{},
			fmt.Errorf("invalid name_generator %s.\nvalid options are: %s", nameGenerator, serializeStringSet(validNameGenerators))
	}

	nodeCountString := request.URL.Query().Get("node_count")
	nodeCountInt, err := strconv.Atoi(nodeCountString)
	if err != nil {
		return searchSpaceParameters{},
			fmt.Errorf("invalid node_count %s: %w", nodeCountString, err)
	}

	return searchSpaceParameters{
		algorithm:     algorithm,
		nameGenerator: nameGenerator,
		nodeCount:     nodeCountInt,
	}, nil
}

func generateSearchSpace(params searchSpaceParameters) (searchSpaceDescriptor, int, error) {
	startTime := time.Now()
	searchSpace, err := searchSpaceGeneration.Generate(params)
	endTime := time.Now()
	if err != nil {
		return nil, 0, fmt.Errorf("generation failed: %w", err)
	}

	return searchSpace, int(endTime.Sub(startTime).Milliseconds()), nil
}

func handleGenerateSearchSpace(writer http.ResponseWriter, request *http.Request) {
	params, err := newSearchSpaceParameters(request)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	searchSpace, durationMs, err := generateSearchSpace(params)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	successResponse := struct {
		Success          bool                  `json:"success"`
		SearchSpace      searchSpaceDescriptor `json:"search_space"`
		GenerationTimeMs int                   `json:"generation_time_ms"`
	}{
		Success:          true,
		SearchSpace:      searchSpace,
		GenerationTimeMs: durationMs,
	}

	response, err := json.Marshal(successResponse)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func serializeStringSet(stringSet map[string]bool) string {
	var keys []string
	for key := range stringSet {
		keys = append(keys, key)
	}

	return strings.Join(keys, ", ")
}
