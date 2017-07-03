package agents

import (
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func configureAuthenticationAgent() *AuthenticationAgent {
	agent := AuthenticationAgent{
		MasterConfiguration: &MasterConfiguration{
			Host:   "127.0.0.1:8080",
			Secret: "foobarfoobarfoobarfoobar",
		},
		Tags:      &AgentTags{"foo": "bar"},
		Frequency: time.Duration(1 * time.Second),
	}

	return &agent
}

func TestGenerateJWT(t *testing.T) {
	a := configureAuthenticationAgent()
	token := a.generateJWT()

	_, err := jwt.ParseWithClaims(
		token,
		jwt.MapClaims{"agent-tags": a.Tags},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(a.MasterConfiguration.Secret), nil
		},
	)

	if err != nil {
		t.Errorf("Error while parsing JWT: %s", err)
	}
}
