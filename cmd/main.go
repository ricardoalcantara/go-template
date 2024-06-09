package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ricardoalcantara/go-template/internal/domain/foo"
	messagebroker "github.com/ricardoalcantara/go-template/internal/message_broker"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
	"github.com/ricardoalcantara/go-template/internal/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg(".env file not find")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	models.ConnectDataBase()
}

func main() {
	messagebroker.Start()

	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"),
		gin.Recovery(),
	)

	if value, ok := os.LookupEnv("CORS_ORIGIN"); ok {
		config := cors.DefaultConfig()
		if value == "*" {
			config.AllowOriginFunc = func(_ string) bool { return true }
		} else {
			config.AllowOrigins = strings.Split(value, ",")
		}

		config.AddAllowHeaders("Authorization")

		r.Use(cors.New(config))
	}

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service":    "go-template",
			"git_commit": version.GitCommit,
			"build_os":   version.BuildOS,
			"build_date": version.BuildDate,
			"start_time": version.StartTime,
			"up_time":    version.GetUptime(),
			"version":    version.Version,
		})
	})

	foo.RegisterRoutes(r)

	host := utils.GetEnv("HOST", "")
	port := utils.GetEnv("PORT", "8999")
	r.Run(host + ":" + port)
}
