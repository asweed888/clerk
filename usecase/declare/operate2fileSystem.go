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


func (u *operate2fileSystem) GenerateArch(datasource string) {
    decl, _ := u.load(datasource)

    for _, d := range decl.Spec {
        path := d.CreateDirectory(".")
        generateUpstream(path, d.Upstream)
        generateCodeFile(path, d.CodeFile)
    }
}

func generateUpstream(path string, upstream []*model.DeclareUpstream) {

}

func generateCodeFile(path string, codeFile []*model.DeclareCodeFile) {

}
