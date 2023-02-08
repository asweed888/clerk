package datastore

import (
	"fmt"

	"github.com/asweed888/clerk/datasources/config"
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type configRepository struct {}



func NewConfigRepository() repository.ConfigRepository {
    return &configRepository{}
}


func (r *configRepository) Load(lang string) (*model.Config, error) {
    switch lang {
    case "go":
        return config.Golang, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
