package agents

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func handleUpdateConfigurationRequest(context echo.Context) error {
	config := context.Get("global_configuration").(map[string]map[string]string)
	json.Unmarshal([]byte(""), &config)

	log.WithFields(log.Fields{
		"type":  "update-configuration",
		"error": false,
	}).Debug("Configuration updated.")

	return context.JSON(http.StatusOK, map[string]string{
		"updated": "true",
	})
}
