package datastore

import (
	"fmt"
	"io/ioutil"

	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
	"gopkg.in/yaml.v2"
)

type declareRepository struct {
    Datasource string
}


func NewDeclareRepository(Datasource string) repository.DeclareRepository {
    return &declareRepository{Datasource}
}



func (r *declareRepository) Load() (*model.Declare, error) {
    f, err := ioutil.ReadFile(r.Datasource)
	if err != nil {
		return nil, fmt.Errorf(
            "ERROR: %s is not found, please create %s.",
            r.Datasource,
            r.Datasource,
        )
	}

    var declare *model.Declare
    err = yaml.Unmarshal(f, &declare)
	if err != nil {
		return nil, fmt.Errorf(
            "%s unmarshal failed",
            r.Datasource,
        )
	}

	return declare, nil

}
