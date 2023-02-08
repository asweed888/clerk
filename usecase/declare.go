package usecase

import (
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type DeclareUsecace interface {
    Load(datasource string) (*model.Declare, error)
}

type declareUsecace struct {
    repository.DeclareRepository
}

func NewDeclareUsecace(r repository.DeclareRepository) DeclareUsecace {
    return &declareUsecace{r}
}


func (u *declareUsecace) Load(datasource string) (*model.Declare, error) {
    return u.DeclareRepository.Load(datasource)
}
