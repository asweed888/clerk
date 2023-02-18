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

func (u *operate2fileSystem) load(datasource string) (*model.Declare, error) {
    return u.DeclareRepository.Load(datasource)
}


func (u *operate2fileSystem) GenerateArch(datasource string)  {
    decl, _ := u.load(datasource)

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
