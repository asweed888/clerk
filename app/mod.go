package app

import (
	"github.com/asweed888/clerk/domain/repository"
	"github.com/asweed888/clerk/infra/persistence/datastore"
	"github.com/asweed888/clerk/presentation"
	"github.com/asweed888/clerk/usecase"
)

type App interface {
    NewDeclareRepository() repository.DeclareRepository
    NewDeclareUseCase() usecase.DeclareUseCase
    NewBuildCommand() presentation.BuildCommand
    NewSubCommand() presentation.SubCommand
}

type app struct {
    DeclareFile string
}

func NewApp(declareFile string) App {
    return &app{declareFile}
}

type subCommand struct {
    presentation.BuildCommand
}

func (a *app) NewSubCommand() presentation.SubCommand {
    subCmd := &subCommand{}
    subCmd.BuildCommand = a.NewBuildCommand()
    return subCmd
}

func (a *app) NewBuildCommand() presentation.BuildCommand {
    return presentation.NewBuildCommand(a.NewDeclareUseCase())
}

func (a *app) NewDeclareUseCase() usecase.DeclareUseCase {
    return usecase.NewDeclareUseCase(a.NewDeclareRepository())
}

func (a *app) NewDeclareRepository() repository.DeclareRepository {
    return datastore.NewDeclareRepository(a.DeclareFile)
}
