package repository

import "github.com/asweed888/clerk/domain/model"

type DeclareRepository interface {
    Load(datasource string) (*model.Declare, error)
}
