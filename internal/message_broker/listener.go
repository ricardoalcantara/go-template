package messagebroker

import (
	"os"

	"github.com/rs/zerolog/log"
)

var rabbitmq *RabbitMQManager

func Start() {
	go start()

	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		Publish(Message{
	// 			Type: MessageTypeDebugMessage,
	// 			Data: "Hello World!",
	// 		})
	// 	}
	// }()
}

func start() {
	rabbitmq = &RabbitMQManager{}
	rabbitmq.Connect()

	err := rabbitmq.Consume(os.Getenv("RABBITMQ_QUEUE"), processor)
	if err != nil {
		log.Error().Err(err).Msg("Failed to consume message")
	}
}
