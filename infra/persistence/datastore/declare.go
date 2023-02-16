package datastore

import (
	"fmt"
	"io/ioutil"

	"github.com/asweed888/clerk/datasources/config"
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
	"gopkg.in/yaml.v2"
)

type declareRepository struct {}


func NewDeclareRepository() repository.DeclareRepository {
    return &declareRepository{}
}



func (r *declareRepository) Load(datasource string) (*model.Declare, error) {
    f, err := ioutil.ReadFile(datasource)
	if err != nil {
		return nil, fmt.Errorf(
            "ERROR: %s is not found, please create %s.",
            datasource,
            datasource,
        )
	}

    var declare *model.Declare
    err = yaml.Unmarshal(f, &declare)
	if err != nil {
		return nil, fmt.Errorf(
            "%s unmarshal failed",
            datasource,
        )
	}

    conf, err := internalConfigLoad(declare.Lang)
    if err != nil {
        return nil, fmt.Errorf("Failed to load the config")
    }

    declare.InternalConfig = conf

	return declare, nil
}


func internalConfigLoad(lang string) (*model.InternalConfig, error) {
    switch lang {
    case "go":
        return config.Golang, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
