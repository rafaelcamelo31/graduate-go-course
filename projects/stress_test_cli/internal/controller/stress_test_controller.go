package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/usecase"
)

type StressTestController struct {
	useCase usecase.StressTestUseCase
}

func NewStressTestController(useCase usecase.StressTestUseCase) *StressTestController {
	return &StressTestController{
		useCase: useCase,
	}
}

func (c *StressTestController) Execute(flags *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	report, err := c.useCase.ExecuteStressTest(ctx, flags)
	if err != nil {
		return fmt.Errorf("stress test execution failed: %w", err)
	}

	c.printReport(report)

	return nil
}

func (c *StressTestController) printReport(report *config.StressTestReport) {
	fmt.Println("\n====================================")
	fmt.Println("    STRESS TEST REPORT")
	fmt.Println("====================================")
	fmt.Printf("Target URL:           %s\n", report.TargetURL)
	fmt.Printf("Total Requests:       %d\n", report.TotalRequests)
	fmt.Printf("Concurrency Level:    %d\n", report.ConcurrencyLevel)
	fmt.Printf("Total Elapsed Time:   %.3fs\n\n", report.TotalElapsedTime)

	fmt.Println("Results:")
	fmt.Printf("Successful (200):  %d\n", report.SuccessfulRequests)
	fmt.Printf("Failed:            %d\n\n", report.FailedRequests)

	fmt.Println("====================================")
}
