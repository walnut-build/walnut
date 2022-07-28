package wfile

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type WalnutFile struct {
	Include []string                   `yaml:"include"`
	Tasks   map[string]TaskDescription `yaml:"tasks"`
}

type TaskDescription struct {
	DependsOn []string `yaml:"depends_on"`
}

func LoadFile(path string) (*WalnutFile, error) {
	wfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := &WalnutFile{}
	err = yaml.Unmarshal(wfile, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
