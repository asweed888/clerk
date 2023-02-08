package usecase

import (
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type ConfigUsecase interface {
    Load(lang string) (*model.Config, error)
}

type configUsecase struct {
    repository.ConfigRepository
}

func NewConfigUsecase(r repository.ConfigRepository) ConfigUsecase {
    return &configUsecase{r}
}


func (u *configUsecase) Load(lang string) (*model.Config, error) {
    return u.ConfigRepository.Load(lang)
}
