package schema

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)


type ClerkYaml struct {
    Export bool `yaml:"export"`
	Lang   string `yaml:"lang"`
	Spec   []struct {
		Location  string `yaml:"location"`
        Comment string `yaml:"comment"`
		Upstream []struct {
			Name    string   `yaml:"name"`
			Comment string   `yaml:"comment"`
            Methods []string `yaml:"function"`
		} `yaml:"upstream"`
	} `yaml:"spec"`
}


func Read() (*ClerkYaml, error) {
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
