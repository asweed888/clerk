package repository

import (
	"github.com/asweed888/clerk/domain/model"
)

type TacitConfigRepository interface {
    Load(lang string) (*model.TacitConfig, error)
}
