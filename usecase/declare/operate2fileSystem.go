package declare

import (
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type Operate2fileSystem interface {
    Load(datasource string) (*model.Declare, error)
}

type operate2fileSystem struct {
    repository.DeclareRepository
}

func NewDeclareOperate2fileSystem(r repository.DeclareRepository) Operate2fileSystem {
    return &operate2fileSystem{r}
}

func (u *operate2fileSystem) Load(datasource string) (*model.Declare, error) {
    return u.DeclareRepository.Load(datasource)
}
