package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type persistenceManager interface {
	CreateSearchSpace(CreateSearchSpaceParameters) (string, error)
	ListSearchSpaces() ([]SearchSpaceIdentifier, error)
	GetSearchSpace(string) (CreateSearchSpaceParameters, error)
	DeleteSearchSpace(string) error
}

type SearchSpaceIdentifier struct {
	Name           string    `json:"name"`
	Id             string    `json:"id"`
	GenerationDate time.Time `json:"generation_date"`
	NodeCount      int       `json:"node_count"`
	GridSizeX      int       `json:"grid_size_x"`
	GridSizeY      int       `json:"grid_size_y"`
}

func newPersistenceManager(managerId string) (persistenceManager, error) {
	switch managerId {
	case "memory":
		return newMemoryPersistenceManager()
	default:
		return nil, fmt.Errorf("invalid search generator id: %s", managerId)
	}
}

type memoryPersistenceManager struct {
	mutex        sync.RWMutex
	searchSpaces map[string]CreateSearchSpaceParameters
}

func newMemoryPersistenceManager() (*memoryPersistenceManager, error) {
	return &memoryPersistenceManager{
		mutex:        sync.RWMutex{},
		searchSpaces: map[string]CreateSearchSpaceParameters{},
	}, nil
}

func (m *memoryPersistenceManager) CreateSearchSpace(searchSpace CreateSearchSpaceParameters) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	id := uuid.New().String()
	m.searchSpaces[id] = searchSpace

	return id, nil
}

func (m *memoryPersistenceManager) ListSearchSpaces() ([]SearchSpaceIdentifier, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	identifiers := []SearchSpaceIdentifier{}
	for id, searchSpace := range m.searchSpaces {
		identifiers = append(identifiers, SearchSpaceIdentifier{
			Id:             id,
			Name:           searchSpace.Name,
			GenerationDate: searchSpace.GenerationDate,
			NodeCount:      len(searchSpace.Nodes),
			GridSizeX:      searchSpace.GridSizeX,
			GridSizeY:      searchSpace.GridSizeY,
		})
	}
	return identifiers, nil
}

func (m *memoryPersistenceManager) GetSearchSpace(id string) (CreateSearchSpaceParameters, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	searchSpace, ok := m.searchSpaces[id]
	if !ok {
		return CreateSearchSpaceParameters{}, fmt.Errorf("failed to get search space '%s'. invalid id", id)
	}

	return searchSpace, nil
}

func (m *memoryPersistenceManager) DeleteSearchSpace(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.searchSpaces, id)
	return nil
}
