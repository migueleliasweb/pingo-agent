package agents

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/migueleliasweb/pingo-agent/common"
	log "github.com/sirupsen/logrus"
)

//AuthenticationAgent Authenticates with master
type AuthenticationAgent struct {
	MasterConfiguration *MasterConfiguration
	Tags                *map[string]string
	ClientHTTP          *common.ClientHTTPPoster
	Ctx                 context.Context

	//Frequency on which the agent authenticates to master
	Frequency time.Duration
}

//generateJWT Generates and return the token
func (agent *AuthenticationAgent) generateJWT() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"agent-tags": agent.Tags,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(agent.MasterConfiguration.Secret)

	if err != nil {
		log.WithFields(log.Fields{
			"type":  "jwt-generation",
			"error": true,
		}).Panicf(err.Error())
	}

	return tokenString
}

//Authenticate Executes JWT authentication
func (agent *AuthenticationAgent) Authenticate() error {
	jwt := agent.generateJWT()

	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	client := http.Client{Timeout: 5}

	ticker := time.NewTicker(agent.Frequency)

	for {
		select {
		case <-ticker.C:
			resp, err := client.Post(
				agent.MasterConfiguration.Host,
				"application/jwt",
				strings.NewReader(jwt),
			)

			if err != nil {
				log.WithFields(log.Fields{
					"type":        "authehtication",
					"status_code": resp.StatusCode,
					"error":       true,
				}).Panicf(err.Error())

				return err
			}
		case <-agent.Ctx.Done():
			ticker.Stop()
			return nil
		}

	}
}
