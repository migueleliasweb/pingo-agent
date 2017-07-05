package agents

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

//WebAgent The WebAgent
type WebAgent struct {
	Address string
}

func (agent WebAgent) handleUpdateConfigurationRequest(context echo.Context) error {
	globalConf := context.Get("global-configuration").(*GlobalAgentsConfiguration)
	bindErr := context.Bind(globalConf)

	if bindErr != nil {
		log.WithFields(log.Fields{
			"type":  "update-configuration-unmarshall-payload",
			"error": true,
		}).Debug(bindErr)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"updated": "false",
		})
	}

	log.WithFields(log.Fields{
		"type":  "update-configuration",
		"error": false,
	}).Debug("Configuration updated.")

	return context.JSON(http.StatusOK, map[string]string{
		"updated": "true",
	})
}

//Execute Executes the agent
func (agent WebAgent) Execute() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", agent.handleUpdateConfigurationRequest)

	// Start server
	e.Logger.Fatal(e.Start(agent.Address))
}
