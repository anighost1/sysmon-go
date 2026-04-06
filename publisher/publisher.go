package publisher

import (
	"encoding/json"

	"sysmon/redisClient"
)

func Publish(channel string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return redisClient.Client.Publish(redisClient.Ctx, channel, jsonData).Err()
}
