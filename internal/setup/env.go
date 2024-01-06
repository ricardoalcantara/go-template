package setup

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Env() {
	err := godotenv.Load()

	if err != nil {
		logrus.Info("Fail loading .env file")
	}
}
