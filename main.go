package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"sysmon/cpu"
	"sysmon/publisher"
	"sysmon/redisClient"
)

func main() {
	// setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init redis
	if err := redisClient.InitRedis(); err != nil {
		logger.Error("Redis connection failed")
		return
	}

	// graceful shutdown
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	go redisClient.HandleGracefulShutdown()

	// main loop
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := cpu.GetCpuData()
			if err != nil {
				logger.Warn("CPU fetch failed")
				continue
			}

			if err := publisher.Publish("cpu_channel", data); err != nil {
				logger.Error("Redis publish failed")
			}

		case <-ctx.Done():
			logger.Info("Shutting down main loop")
			return
		}
	}
}
