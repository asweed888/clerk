package usecase

import (
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type DeclareUseCase  interface {
    GenerateArch() error
}

type declareUseCase struct {
    repository.DeclareRepository
}

func NewDeclareUseCase(r repository.DeclareRepository) DeclareUseCase {
    return &declareUseCase{r}
}


func (u *declareUseCase) GenerateArch() error {
    decl, _ := u.DeclareRepository.Load()

    for _, s := range decl.Spec {
        workdir := s.ChangeDirectory(".")
        if err := s.CreateDirectory(workdir); err != nil { return err }

        if len(s.Upstream) != 0 {
            if err := generateUpstream(workdir, s.Upstream); err != nil { return err }
        }

        if len(s.CodeFile) != 0 {
            if err := generateCodeFile(workdir, s.CodeFile); err != nil { return err }
        }
    }

    return nil
}

func generateUpstream(prevWorkDir string, upstream []*model.DeclareUpstream) error {

    for _, u := range upstream {
        workdir := u.ChangeDirectory(prevWorkDir)
        if err := u.CreateDirectory(workdir); err != nil { return err }

        if len(u.Upstream) != 0 {
            if err := generateUpstream(workdir, u.Upstream); err != nil { return err }
        }

        if len(u.CodeFile) != 0 {
            if err := generateCodeFile(workdir, u.CodeFile); err != nil { return err }
        }
    }
    return nil
}

func generateCodeFile(workdir string, codeFile []*model.DeclareCodeFile) error {

    for _, c := range codeFile {
        c.CreateCodeFile(workdir)
    }
    return nil
}
