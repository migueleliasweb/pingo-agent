package agents

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func handleUpdateConfigurationRequest(context echo.Context, globalConf *GlobalAgentsConfiguration) error {
	// config := context.Get("global_configuration").(map[string]map[string]string)

	receivedConfig := make(GlobalAgentsConfiguration)

	context.Bind(receivedConfig)

	// globalConf = receivedConfig

	log.WithFields(log.Fields{
		"type":  "update-configuration",
		"error": false,
	}).Debug("Configuration updated.")

	return context.JSON(http.StatusOK, map[string]string{
		"updated": "true",
	})
}
