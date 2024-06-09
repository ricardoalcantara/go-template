package messagebroker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const (
	MessageTypeDomainUpdate MessageType = "domain update"
	MessageTypeDomainDelete MessageType = "domain delete"
)

func processor(delivery amqp.Delivery) {
	switch delivery.Type {
	case string(MessageTypeDebugMessage):
		log.Info().Any("data", delivery.Body).Msg("Debug message")
	}
}
