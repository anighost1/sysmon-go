package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"sysmon/cpu"
	"sysmon/disk"
	"sysmon/memory"
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

			// CPU
			if data, err := cpu.GetCpuData(); err == nil {
				if err := publisher.Publish("metrics:cpu", data); err != nil {
					logger.Error("CPU publish failed")
				}
			} else {
				logger.Warn("CPU fetch failed")
			}

			// MEMORY
			if data, err := memory.GetMemoryData(); err == nil {
				if err := publisher.Publish("metrics:memory", data); err != nil {
					logger.Error("Memory publish failed")
				}
			} else {
				logger.Warn("Memory fetch failed")
			}

			// DISK
			if data, err := disk.GetDiskData(); err == nil {
				if err := publisher.Publish("metrics:disk", data); err != nil {
					logger.Error("Disk publish failed")
				}
			} else {
				logger.Warn("Disk fetch failed")
			}

		case <-ctx.Done():
			logger.Info("Shutting down main loop")
			return
		}
	}
}
