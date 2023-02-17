package datastore

import (
	"github.com/asweed888/clerk/datasources/metadata"
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type cwdRepository struct {}

func NewCwdRepository() repository.CwdRepository {
    return &cwdRepository{}
}


func (r *cwdRepository) Load() *model.Cwd {
    return metadata.Cwd
}
