package config

import (
	"os"

	"github.com/getsentry/sentry-go"
	"gopkg.in/yaml.v2"
)

type App struct {
	Openai struct {
		Gpt3Token string `yaml:"gpt3Token"`
	}
}

func (app *App) Gpt3Token() string {
	return app.Openai.Gpt3Token
}

func InstanceApp() *App {
	var app App
	yml := YAML{}
	content, err := os.ReadFile(yml.PathFile())
	if err != nil {
		sentry.CaptureMessage("App / InstanceApp() / os.ReadFile(yml.PathFile()) / " + err.Error())
	}

	err = yaml.Unmarshal(content, &app)
	if err != nil {
		sentry.CaptureMessage("App / InstanceApp() / yaml.Unmarshal(content, &app) / " + err.Error())
	}

	return &app
}
