package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/canopener/pathfinder/api_server/ssgen"
)

type searchSpaceDescriptor interface{}
type searchSpaceParameters struct {
	NodeConnectorId     string `json:"node_connector_id"`
	NameGeneratorId     string `json:"name_generator_id"`
	NodeCount           int    `json:"node_count"`
	MinimumNodeDistance int    `json:"minimum_node_distance"`
}

func (p searchSpaceParameters) GetNodeConnectorId() string  { return p.NodeConnectorId }
func (p searchSpaceParameters) GetNameGeneratorId() string  { return p.NameGeneratorId }
func (p searchSpaceParameters) GetNodeCount() int           { return p.NodeCount }
func (p searchSpaceParameters) GetMinimumNodeDistance() int { return p.MinimumNodeDistance }

type GenerationParameters interface {
	GetNodeConnectorId() string
	GetNameGeneratorId() string
	GetNodeCount() int
	GetMinimumNodeDistance() int
}

var validNodeConnectorIds = map[string]bool{
	"default":      true,
	"prim":         true,
	"min_two_conn": true,
	"none":         true,
}

var validNameGeneratorIds = map[string]bool{
	"default":       true,
	"three_letters": true,
}

func newSearchSpaceParameters(request *http.Request) (searchSpaceParameters, error) {
	queryParams := request.URL.Query()
	nodeConnectorId := queryParams.Get("node_connector_id")
	_, ok := validNodeConnectorIds[nodeConnectorId]
	if !ok {
		return searchSpaceParameters{},
			fmt.Errorf("invalid node_connector_id %s.\nvalid options are: %s",
				nodeConnectorId, serializeStringSet(validNodeConnectorIds))
	}

	nameGeneratorId := request.URL.Query().Get("name_generator_id")
	_, ok = validNameGeneratorIds[nameGeneratorId]
	if !ok {
		return searchSpaceParameters{},
			fmt.Errorf("invalid name_generator_id %s.\nvalid options are: %s",
				nameGeneratorId, serializeStringSet(validNameGeneratorIds))
	}

	nodeCountString := request.URL.Query().Get("node_count")
	nodeCountInt, err := strconv.Atoi(nodeCountString)
	if err != nil {
		return searchSpaceParameters{},
			fmt.Errorf("invalid node_count %s: %w", nodeCountString, err)
	}

	minimumNodeDistanceString := request.URL.Query().Get("minimum_node_distance")
	minimumNodeDistanceInt, err := strconv.Atoi(minimumNodeDistanceString)
	if err != nil {
		return searchSpaceParameters{},
			fmt.Errorf("invalid minimum_node_distance %s: %w", minimumNodeDistanceString, err)
	}

	return searchSpaceParameters{
		NodeConnectorId:     nodeConnectorId,
		NameGeneratorId:     nameGeneratorId,
		NodeCount:           nodeCountInt,
		MinimumNodeDistance: minimumNodeDistanceInt,
	}, nil
}

func generateSearchSpace(params searchSpaceParameters) (searchSpaceDescriptor, int, error) {
	startTime := time.Now()
	searchSpace, err := ssgen.Generate(params)
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
