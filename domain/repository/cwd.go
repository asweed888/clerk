package repository

import "github.com/asweed888/clerk/domain/model"

type CwdRepository interface {
    Load() *model.Cwd
}
