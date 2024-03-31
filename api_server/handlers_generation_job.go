package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func createGenerationJob(writer http.ResponseWriter, request *http.Request) {
	parameters, err := generationJobParametersFromRequest(request)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	job_id, err := mainGenerationJobManager.NewJob(parameters)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success         bool   `json:"success"`
		GenerationJobId string `json:"generation_job_id"`
	}{
		Success:         true,
		GenerationJobId: job_id,
	})
}

func showGenerationJob(writer http.ResponseWriter, request *http.Request) {
	jobId := mux.Vars(request)["id"]
	jobStatus, err := mainGenerationJobManager.Jobstatus(jobId)
	if err != nil {
		sendErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sendSuccessJsonResponse(writer, struct {
		Success bool                `json:"success"`
		Status  GenerationJobStatus `json:"status"`
	}{
		Success: true,
		Status:  jobStatus,
	})
}
