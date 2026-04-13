package subscriber

import (
	"context"

	"sysmon/redisClient"
	"sysmon/sse"
)

func StartSubscriber() {
	pubsub := redisClient.Client.Subscribe(
		context.Background(),
		"metrics:cpu",
		"metrics:memory",
		"metrics:disk",
	)

	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			// log.Println("Received:", msg.Channel, msg.Payload)

			// push to SSE clients
			sse.Broadcast(msg.Payload)
		}
	}()
}
