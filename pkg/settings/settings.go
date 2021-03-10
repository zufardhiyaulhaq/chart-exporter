package settings

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	UseServiceAccount bool `envconfig:"USE_SERVICE_ACCOUNT" default:"false"`
}

type Option func(*Settings)

func NewSettings() Settings {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		log.Fatalln(err)
	}

	return settings
}
