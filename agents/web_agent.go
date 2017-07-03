package agents

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func handleUpdateConfigurationRequest(context echo.Context, globalConf *GlobalAgentsConfiguration) error {
	bindErr := context.Bind(globalConf)

	if bindErr != nil {
		log.WithFields(log.Fields{
			"type":  "update-configuration-unmarshall-payload",
			"error": true,
		}).Debug(bindErr)
	} else {
		log.WithFields(log.Fields{
			"type":  "update-configuration",
			"error": false,
		}).Debug("Configuration updated.")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"updated": "true",
	})
}
