package agents

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func getFakeUpdateCtx(method string, uri string, body string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(
		method,
		uri,
		strings.NewReader(body),
	)

	req.Header = map[string][]string{
		"Content-type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c
}

func TestWebHandle(t *testing.T) {
	sentConfiguration := `
        a
        a
    `

	globalConf := make(GlobalAgentsConfiguration)

	ctx := getFakeUpdateCtx("POST", "/foo", sentConfiguration)

	handleUpdateConfigurationRequest(ctx, &globalConf)
}
