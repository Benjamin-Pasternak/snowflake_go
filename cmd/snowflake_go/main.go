package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Benjamin-Pasternak/snowflake_go/internal/data"
	"github.com/Benjamin-Pasternak/snowflake_go/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// load in config from application.yaml
	util.InitConfig()

	snowflake, err := data.NewSnowFlake(1)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating Snowflake")
	}

	r := gin.Default()
	r.GET("/generate-id", func(c *gin.Context) {
		id, err := snowflake.GenerateId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	if err := r.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}

}
