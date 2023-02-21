package interactor

import (
	"github.com/asweed888/clerk/domain/repository"
	"github.com/asweed888/clerk/usecase"
)

type Interactor interface {
    NewDeclareRepository() repository.DeclareRepository
    NewDeclareUseCase() usecase.DeclareUseCase
}
