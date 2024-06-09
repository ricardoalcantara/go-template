package messagebroker

import (
	"encoding/json"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type MessageType string

const (
	MessageTypeDebugMessage MessageType = "debug message"
)

type Message struct {
	Type MessageType `json:"type"`
	Data []byte      `json:"data"`
}

func NewMessageToJson(messageType MessageType, data any) Message {
	b, err := json.Marshal(&data)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal message")
	}
	return Message{
		Type: messageType,
		Data: b,
	}
}

type RabbitMQManager struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (m *RabbitMQManager) Connect() {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		log.Warn().Msg("RABBITMQ_URL not found")
		return
	}

	vhost := os.Getenv("RABBITMQ_VHOST")
	if vhost == "" {
		vhost = "/"
	}

	var err error
	attemptDelay := 5
	for {
		m.conn, err = amqp.DialConfig(url, amqp.Config{
			Vhost:  vhost,
			Locale: "en_US",
		})
		if err == nil {
			break
		}
		log.Error().Err(err).Msg("Failed to connect to RabbitMQ")
		time.Sleep(time.Duration(attemptDelay) * time.Second)
		attemptDelay *= 2
	}
}

func (m *RabbitMQManager) Close() {
	if m.conn != nil {
		m.conn.Close()
	}

	if m.channel != nil {
		m.channel.Close()
	}
}

func (m *RabbitMQManager) Channel() (*amqp.Channel, error) {
	if m.conn == nil {
		return nil, amqp.ErrClosed
	}

	if m.channel == nil {
		var err error
		m.channel, err = m.conn.Channel()
		if err != nil {
			return nil, err
		}
	}

	return m.channel, nil
}

func (m *RabbitMQManager) Consume(queue string, processor func(amqp.Delivery)) error {
	ch, err := m.Channel()
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		queue,            // queue
		"platform-proxy", // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		processor(d)
	}
	return nil
}

func (m *RabbitMQManager) Publish(exchange string, routingKey string, message Message) error {
	ch, err := m.Channel()
	if err != nil {
		return err
	}

	return ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Type:        string(message.Type),
			Body:        message.Data,
		})
}
