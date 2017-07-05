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
		"Content-Type": {echo.MIMEApplicationJSON},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c
}

func TestWebHandle(t *testing.T) {
	sentConfiguration := `
        {
            "http": {
                "target": "www.google.com",
                "timeout": 10
            }
        }
    `

	globalConf := GlobalAgentsConfiguration{}
	ctx := getFakeUpdateCtx("POST", "/foo", sentConfiguration)

	ctx.Set("global-configuration", &globalConf)

	wa := WebAgent{}

	wa.handleUpdateConfigurationRequest(ctx)

	if len(globalConf) == 0 {
		t.Error("Global configuration is empty.")
	}

	//I really tried to use DeepEqual but I couldn't make it to work
	// if reflect.DeepEqual(globalConf, expectedGlobalConf) {
	// 	t.Error("Actual configuration differs from expected.")
	// }

	if globalConf["http"]["target"] != "www.google.com" {
		t.Error("Error on configuration target.")
	}

	if globalConf["http"]["timeout"].(float64) != 10 {
		t.Error("Error on configuration timeout.")
	}
}
