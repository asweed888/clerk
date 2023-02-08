package repository

import "github.com/asweed888/clerk/domain/model"


type ConfigRepository interface {
    Load(lang string) (*model.Config, error)
}
