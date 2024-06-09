package messagebroker

import (
	"os"

	"github.com/rs/zerolog/log"
)

func Publish(message Message) {
	if rabbitmq == nil {
		log.Warn().Msg("RabbitMQ not connected")
		return
	}

	err := rabbitmq.Publish(os.Getenv("RABBITMQ_EXCHANGE"), "", message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to publish message")
		return
	}
}
