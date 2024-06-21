package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	url := "abcd"
	logger, _ := zap.NewProduction()
	fmt.Println("before")
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	fmt.Println("after")
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	sugar.Debug("debugging")
	sugar.Debug("debugging")

	sugar.Level()
	sugar.Log(1, "good")
	sugar.Log(2, "good")
	// sugar.Log(3, "good")
	// sugar.Log(4, "good")

	// sugar.WithLazy("lazy loading").Log(1, "chained")

	sugar.WithLazy(func() zap.Field {
		return zap.String("lazy field", expensiveComputation())
	}).Debug("message")

	sugar.Info(expensiveComputation())
	sugar.Debug(expensiveComputation())


	// sugar.
}

func expensiveComputation() string {
	time.Sleep(2 * time.Second)
	fmt.Println("called")
	return "good"
}
