package utils

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sonda2208/sso-test/model"
)

func LoadConfig(prefix string) (*model.ServiceSettings, error) {
	conf := &model.ServiceSettings{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = envconfig.Process(prefix, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
