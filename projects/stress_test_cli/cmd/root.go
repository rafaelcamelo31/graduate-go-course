package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/controller"
	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/repository"
	"github.com/rafaelcamelo31/graduate-go-course/projects/stress_test_cli/internal/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "stress_test_cli",
	Short: "Stress test your web app using stress_test_cli",
	Long: `Stress test your web app by providing the target URL, total requests,
and number of concurrent calls as arguments.

The CLI will generate a report with total execution time, total requests,
successful requests (status 200), and failed requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := executeStressTest(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "Target URL for stress testing (required)")
	rootCmd.Flags().IntVar(&requests, "requests", 0, "Total number of requests to send (required)")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 0, "Number of concurrent workers (required)")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")

	viper.BindPFlag("url", rootCmd.Flags().Lookup("url"))
	viper.BindPFlag("requests", rootCmd.Flags().Lookup("requests"))
	viper.BindPFlag("concurrency", rootCmd.Flags().Lookup("concurrency"))
}

func executeStressTest() error {
	httpClientRepo := repository.NewHTTPClientRepository(30 * time.Second)
	resultStoreRepo := repository.NewResultStoreRepository()

	stressTestUseCase := usecase.NewStressTestUseCase(httpClientRepo, resultStoreRepo)

	stressTestController := controller.NewStressTestController(stressTestUseCase)

	flags, err := config.NewConfig(url, requests, concurrency)
	if err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	return stressTestController.Execute(flags)
}
