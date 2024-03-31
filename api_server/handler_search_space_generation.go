package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// const defaultGeneratorId = "ssgen"

var validNodeConnectorIds = map[string]bool{
	"default":      true,
	"prim":         true,
	"min_two_conn": true,
	"none":         true,
}

var validNameGeneratorIds = map[string]bool{
	"default":       true,
	"three_letters": true,
	"cities":        true,
	"countries":     true,
	"first_names":   true,
}

func newsearchSpaceGenerationParameters(request *http.Request) (SearchSpaceGenerationParameters, error) {
	queryParams := request.URL.Query()
	nodeConnectorId := queryParams.Get("node_connector_id")
	_, ok := validNodeConnectorIds[nodeConnectorId]
	if !ok {
		return SearchSpaceGenerationParameters{},
			fmt.Errorf("invalid node_connector_id %s.\nvalid options are: %s",
				nodeConnectorId, serializeStringSet(validNodeConnectorIds))
	}

	nameGeneratorId := request.URL.Query().Get("name_generator_id")
	_, ok = validNameGeneratorIds[nameGeneratorId]
	if !ok {
		return SearchSpaceGenerationParameters{},
			fmt.Errorf("invalid name_generator_id %s.\nvalid options are: %s",
				nameGeneratorId, serializeStringSet(validNameGeneratorIds))
	}

	nodeCountString := request.URL.Query().Get("node_count")
	nodeCountInt, err := strconv.Atoi(nodeCountString)
	if err != nil {
		return SearchSpaceGenerationParameters{},
			fmt.Errorf("invalid node_count %s: %w", nodeCountString, err)
	}

	minimumNodeDistanceString := request.URL.Query().Get("minimum_node_distance")
	minimumNodeDistanceInt, err := strconv.Atoi(minimumNodeDistanceString)
	if err != nil {
		return SearchSpaceGenerationParameters{},
			fmt.Errorf("invalid minimum_node_distance %s: %w", minimumNodeDistanceString, err)
	}

	return SearchSpaceGenerationParameters{
		NodeConnectorId:     nodeConnectorId,
		NameGeneratorId:     nameGeneratorId,
		NodeCount:           nodeCountInt,
		MinimumNodeDistance: minimumNodeDistanceInt,
	}, nil
}

func handleGenerateSearchSpace(writer http.ResponseWriter, request *http.Request) {
	params, err := newsearchSpaceGenerationParameters(request)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	generator, err := newSearchSpaceGenerator(defaultGeneratorId)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	searchSpace, durationMs, err := generator.Generate(params)
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
