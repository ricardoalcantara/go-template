package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	models.ConnectDataBase()
}

func main() {
	log.Info().Msg("Playground")
}
