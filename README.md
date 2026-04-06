# 🚀 SysMon

Lightweight system monitoring service built with Golang that streams
real-time system metrics (CPU, Memory, Disk) to Redis.

## Features

-   CPU monitoring (total + per core)
-   Memory usage tracking
-   Disk usage tracking
-   Redis Pub/Sub for real-time updates
-   Modular architecture
-   Structured logging using slog
-   Graceful shutdown support

## Project Structure

sysmon/ ├── cpu/ ├── memory/ ├── disk/ ├── publisher/ ├── redisClient/
├── tools/ ├── main.go

## Setup

git clone `https://github.com/anighost1/sysmon-go.git`{=html} cd sysmon go mod tidy

## Run

go run .

## Build

go build -o sysmon.exe

## Redis Channels

metrics:cpu metrics:memory metrics:disk

