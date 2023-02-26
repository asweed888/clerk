package interactor

import (
	"github.com/asweed888/clerk/adapter"
	"github.com/asweed888/clerk/domain/repository"
	"github.com/asweed888/clerk/infra/persistence/datastore"
	"github.com/asweed888/clerk/usecase"
)

type Interactor interface {
    NewDeclareRepository() repository.DeclareRepository
    NewDeclareUseCase() usecase.DeclareUseCase
    NewBuildCommand() adapter.BuildCommand
    NewSubCommand() adapter.SubCommand
}

type interactor struct {
    DeclareFile string
}

func NewInteractor(declareFile string) Interactor {
    return &interactor{declareFile}
}

type subCommand struct {
    adapter.BuildCommand
}

func (i *interactor) NewSubCommand() adapter.SubCommand {
    subCmd := &subCommand{}
    subCmd.BuildCommand = i.NewBuildCommand()
    return subCmd
}

func (i *interactor) NewBuildCommand() adapter.BuildCommand {
    return adapter.NewBuildCommand(i.NewDeclareUseCase())
}

func (i *interactor) NewDeclareUseCase() usecase.DeclareUseCase {
    return usecase.NewDeclareUseCase(i.NewDeclareRepository())
}

func (i *interactor) NewDeclareRepository() repository.DeclareRepository {
    return datastore.NewDeclareRepository(i.DeclareFile)
}
