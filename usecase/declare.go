package usecase

import (
	"github.com/asweed888/clerk/domain/model"
	"github.com/asweed888/clerk/domain/repository"
)

type DeclareUseCase  interface {
    GenerateArch(tacitConfig *model.TacitConfig)
}

type declareUseCase struct {
    repository.DeclareRepository
}

func NewDeclareUseCase(r repository.DeclareRepository) DeclareUseCase {
    return &declareUseCase{r}
}


func (u *declareUseCase) GenerateArch(tacitConfig *model.TacitConfig)  {
    decl, _ := u.DeclareRepository.Load()
    decl.SetTacitConfig(tacitConfig)

    for _, s := range decl.Spec {
        workdir := s.ChangeDirectory(".")
        s.CreateDirectory(workdir)

        if len(s.Upstream) != 0 {
            generateUpstream(workdir, s.Upstream)
        }

        if len(s.CodeFile) != 0 {
            generateCodeFile(workdir, s.CodeFile)
        }
    }
}

func generateUpstream(prevWorkDir string, upstream []*model.DeclareUpstream)  {

    for _, u := range upstream {
        workdir := u.ChangeDirectory(prevWorkDir)
        u.CreateDirectory(workdir)

        if len(u.Upstream) != 0 {
            generateUpstream(workdir, u.Upstream)
        }

        if len(u.CodeFile) != 0 {
            generateCodeFile(workdir, u.CodeFile)
        }
    }
}

func generateCodeFile(workdir string, codeFile []*model.DeclareCodeFile)  {

    for _, c := range codeFile {
        c.CreateCodeFile(workdir)
    }
}
