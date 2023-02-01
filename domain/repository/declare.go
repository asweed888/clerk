package repository

import "github.com/asweed888/clerk/domain/model"

type DeclareRepository interface {
    Load() (*model.Declare, error)
}
