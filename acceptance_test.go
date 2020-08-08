package votearama

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cucumber/godog"
)

type testContext struct {
	Response *http.Response
}

var testCtx *testContext

func (t *testContext) iSendRequestTo(method, path string) error {
	var err error
	client := &http.Client{
		Timeout: time.Second * 1,
	}
	url, _ := url.Parse("http://web:8080" + path)
	t.Response, err = client.Do(
		&http.Request{
			URL:    url,
			Method: method,
		},
	)

	return err
}

func (t *testContext) theResponseCodeWillBe(code int) error {
	if t.Response.StatusCode != code {
		return fmt.Errorf("expected response with code %d but got %d", code, t.Response.StatusCode)
	}
	return nil
}

func (t *testContext) theBodyWillBe(expected string) error {
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(t.Response.Body)
	if err != nil {
		return err
	}

	if buf.String() != expected {
		return fmt.Errorf("expected response with body \"%s\" but got \"%s\"", expected, buf.String())
	}

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		testCtx = &testContext{}
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		testCtx = &testContext{}
	})
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, testCtx.iSendRequestTo)
	ctx.Step(`^the response code will be (\d+)$`, testCtx.theResponseCodeWillBe)
	ctx.Step(`^the body will be "([^"]*)"$`, testCtx.theBodyWillBe)
}
