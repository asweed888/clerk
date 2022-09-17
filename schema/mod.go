package schema

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)


type ClerkYaml struct {
	Lang string `yaml:"lang"`
    Export bool `yaml:"export"`
	Spec []struct {
		Location string `yaml:"location"`
		Services  []struct {
			Name      string `yaml:"name"`
			Upstreams []struct {
				Name    string `yaml:"name"`
				Comment string `yaml:"comment"`
			} `yaml:"upstreams"`
		} `yaml:"services"`
	} `yaml:"spec"`
}


func Load() (*ClerkYaml, error) {
    f, err := ioutil.ReadFile("./clerk.yml")
	if err != nil {
		return nil, fmt.Errorf("ERROR: clerk.yml is not found, please create clerk.yml.")
	}

    yml := &ClerkYaml{}
    err = yaml.Unmarshal(f, &yml)
	if err != nil {
		return nil, fmt.Errorf("clerk.yml unmarshal failed")
	}

	return yml, nil
}
