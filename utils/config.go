package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sonda2208/sso-test/model"
)

func LoadConfigFromFile(path string) (*model.ServiceSettings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	setting := &model.ServiceSettings{}
	err = json.Unmarshal(data, setting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}
