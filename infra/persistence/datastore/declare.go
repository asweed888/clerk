package datastore

import (
	"fmt"
	"io/ioutil"

	"github.com/asweed888/clerk/datasources/config"
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
	"gopkg.in/yaml.v2"
)

type declareRepository struct {
    DeclareFile string
}


func NewDeclareRepository(declareFile string) repository.DeclareRepository {
    return &declareRepository{declareFile}
}



func (r *declareRepository) Load() (*model.Declare, error) {
    f, err := ioutil.ReadFile(r.DeclareFile)
	if err != nil {
		return nil, fmt.Errorf(
            "ERROR: %s is not found, please create %s.",
            r.DeclareFile,
            r.DeclareFile,
        )
	}

    var declare *model.Declare
    err = yaml.Unmarshal(f, &declare)
	if err != nil {
		return nil, fmt.Errorf(
            "%s unmarshal failed",
            r.DeclareFile,
        )
	}

    if err != nil {
        return nil, fmt.Errorf("Failed to load the config")
    }

	return declare, nil
}


func internalConfigLoad(lang string) (*model.TacitConfig, error) {
    switch lang {
    case "go":
        return config.Golang, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
