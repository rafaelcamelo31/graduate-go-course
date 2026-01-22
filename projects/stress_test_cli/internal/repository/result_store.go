package repository

import (
	"sync"

	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/config"
)

type ResultStoreRepository interface {
	AddResult(result *config.Result)
	GetResults() []config.Result
	GetStatusCodes() (int, int)
}

type resultStoreRepository struct {
	mu      sync.RWMutex
	results []config.Result
}

func NewResultStoreRepository() ResultStoreRepository {
	return &resultStoreRepository{
		results: make([]config.Result, 0),
	}
}

func (r *resultStoreRepository) AddResult(result *config.Result) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.results = append(r.results, *result)
}

func (r *resultStoreRepository) GetResults() []config.Result {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.results
}

func (r *resultStoreRepository) GetStatusCodes() (int, int) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.results) == 0 {
		return 0, 0
	}

	successful := 0
	failed := 0
	totalTime := int64(0)

	for _, result := range r.results {
		if result.Error != nil {
			failed++
		} else if result.StatusCode == 200 {
			successful++
		} else {
			failed++
		}

		totalTime += result.ResponseTime
	}

	return successful, failed
}
