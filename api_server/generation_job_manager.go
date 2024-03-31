package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type generationJobManager interface {
	NewJob(GenerationJobParameters) (string, error)
	Jobstatus(string) (GenerationJobStatus, error)
}

type GenerationJobStatus struct {
	Status     string                  `json:"status"`
	Parameters GenerationJobParameters `json:"parameters"`
	StartTime  time.Time               `json:"startTime"`
	EndTime    time.Time               `json:"endTime"`
	Result     searchSpaceDescriptor   `json:"result"`
	Err        error                   `json:"err"`
}

func newGenerationJobManager(managerId string, generator searchSpaceGenerator) (generationJobManager, error) {
	switch managerId {
	case "mutex":
		return newMutexGenerationJobManager(generator)
	default:
		return nil, fmt.Errorf("invalid generation job manager id :%s", managerId)
	}
}

type mutexGenerationJobManager struct {
	generator   searchSpaceGenerator
	mutex       sync.RWMutex
	jobStatuses map[string]GenerationJobStatus
}

func newMutexGenerationJobManager(generator searchSpaceGenerator) (*mutexGenerationJobManager, error) {
	return &mutexGenerationJobManager{
		generator:   generator,
		mutex:       sync.RWMutex{},
		jobStatuses: map[string]GenerationJobStatus{},
	}, nil
}

func (m *mutexGenerationJobManager) NewJob(parameters GenerationJobParameters) (string, error) {
	jobId := m.initializeJob(parameters)
	go m.startJob(jobId)

	return jobId, nil
}

func (m *mutexGenerationJobManager) Jobstatus(jobId string) (GenerationJobStatus, error) {
	jobStatus := m.getJobStatus(jobId)
	if jobStatus.Status == "finished" || jobStatus.Status == "error" {
		go m.deleteJob(jobId)
	}

	return jobStatus, nil
}

func (m *mutexGenerationJobManager) initializeJob(parameters GenerationJobParameters) string {
	jobId := uuid.New().String()
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.jobStatuses[jobId] = GenerationJobStatus{
		Status:     "initializing",
		Parameters: parameters,
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		Result:     nil,
		Err:        nil,
	}

	return jobId
}

func (m *mutexGenerationJobManager) updateJobStatus(id string, newStatus string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	status, ok := m.jobStatuses[id]
	if !ok {
		log.Fatalf("failed to update job status to '%s' invalid job id: %s",
			newStatus, id)
	}

	status.Status = newStatus
	m.jobStatuses[id] = status
}

func (m *mutexGenerationJobManager) updateJobStatusFinished(id string, searchSpace searchSpaceDescriptor) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	status, ok := m.jobStatuses[id]
	if !ok {
		log.Fatalf("failed to update job status to 'finished' invalid job id: %s", id)
	}

	status.Status = "finished"
	status.EndTime = time.Now()
	status.Result = searchSpace
	m.jobStatuses[id] = status
}

func (m *mutexGenerationJobManager) updateJobStatusError(id string, err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	status, ok := m.jobStatuses[id]
	if !ok {
		log.Fatalf("failed to update job status to 'error' invalid job id: %s. err: %v",
			id, err)
	}

	status.Status = "error"
	status.EndTime = time.Now()
	status.Err = err
	m.jobStatuses[id] = status
}

func (m *mutexGenerationJobManager) deleteJob(id string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.jobStatuses[id]
	if !ok {
		log.Fatalf("failed to delete job. invalid job id: %s", id)
	}

	delete(m.jobStatuses, id)
}

func (m *mutexGenerationJobManager) getJobStatus(id string) GenerationJobStatus {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	status, ok := m.jobStatuses[id]
	if !ok {
		log.Fatalf("failed to read job. invalid job id: %s", id)
	}

	return status
}

func (m *mutexGenerationJobManager) startJob(id string) {
	m.updateJobStatus(id, "generating")
	jobStatus := m.getJobStatus(id)
	searchSpace, err := m.generator.Generate(jobStatus.Parameters)
	if err != nil {
		m.updateJobStatusError(id, err)
		return
	}

	m.updateJobStatusFinished(id, searchSpace)
}
