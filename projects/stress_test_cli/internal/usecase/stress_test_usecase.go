package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/repository"
)

type StressTestUseCase interface {
	ExecuteStressTest(ctx context.Context, cfg *config.Config) (*config.StressTestReport, error)
}

type stressTestUseCase struct {
	httpClientRepo  repository.HTTPClientRepository
	resultStoreRepo repository.ResultStoreRepository
}

func NewStressTestUseCase(
	httpClientRepo repository.HTTPClientRepository,
	resultStoreRepo repository.ResultStoreRepository,
) StressTestUseCase {
	return &stressTestUseCase{
		httpClientRepo:  httpClientRepo,
		resultStoreRepo: resultStoreRepo,
	}
}

func (s *stressTestUseCase) ExecuteStressTest(ctx context.Context, cfg *config.Config) (*config.StressTestReport, error) {
	startTime := time.Now()

	requestsPerWorker := cfg.Requests / cfg.Concurrency
	remainder := cfg.Requests % cfg.Concurrency

	var wg sync.WaitGroup
	wg.Add(cfg.Concurrency)

	for i := 0; i < cfg.Concurrency; i++ {
		go func(workerID int) {
			defer wg.Done()

			requests := requestsPerWorker
			if i < remainder {
				requests++
			}

			for j := 0; j < requests; j++ {
				statusCode, duration, err := s.httpClientRepo.SendRequest(ctx, cfg.URL)

				result := config.Result{
					StatusCode:   statusCode,
					ResponseTime: duration,
					Error:        err,
				}

				s.resultStoreRepo.AddResult(&result)
			}
		}(i)
	}

	wg.Wait()

	totalElapsed := time.Since(startTime).Seconds()
	successful, failed := s.resultStoreRepo.GetStatusCodes()

	report := &config.StressTestReport{
		TargetURL:          cfg.URL,
		TotalRequests:      cfg.Requests,
		ConcurrencyLevel:   cfg.Concurrency,
		TotalElapsedTime:   totalElapsed,
		SuccessfulRequests: successful,
		FailedRequests:     failed,
	}

	return report, nil
}
