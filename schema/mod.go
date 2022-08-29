package schema

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ClerkYamlSchema struct {
	Lang string `yaml:"lang"`
	Spec []struct {
		Name      string `yaml:"name"`
		Upstreams []struct {
			Name      string `yaml:"name"`
			Upstreams []struct {
				Name    string `yaml:"name"`
				Comment string `yaml:"comment"`
			} `yaml:"upstreams"`
		} `yaml:"upstreams"`
	} `yaml:"spec"`
}


func ClerkYaml() (*ClerkYamlSchema, error) {
    f, err := ioutil.ReadFile("./clerk.yml")
	if err != nil {
		return nil, fmt.Errorf("ERROR: clerk.yml is not found, please create clerk.yml.")
	}

    schema := &ClerkYamlSchema{}
    err = yaml.Unmarshal(f, &schema)
	if err != nil {
		return nil, fmt.Errorf("clerk.yml unmarshal failed")
	}

	return schema, nil
}
