package config

import (
	"github.com/asweed888/clerk/domain/model"
)

type golangUtils struct {
    utils
}


var Golang = &model.TacitConfig{
    Ext: "go",
    FileModeStr: "0644",
    ImplInitialFileContents: func(c *model.TacitConfig, path string, fname string) string {

        ut := &golangUtils{utils{c, path, fname}}

        if ut.IsDomainModel() {
            return ut.DomainModelFileContents()
        } else if ut.IsDomainRepository() {
            return ut.DomainRepositoryFileContents()
        } else if ut.IsInfra() {
            return ut.InfraFileContents()
        } else if ut.IsUseCase() {
            return ut.UseCaseFileContents()
        } else if ut.IsPresentation() {
            return ut.PresentationFileContents()
        } else if ut.IsDi() {
            return ut.DiFileContents()
        } else {
            return ut.DefaultFileContents()
        }
    },
}


func (u *golangUtils) DomainModelFileContents() string {
    tmpl := `package {{ .Pkgname }}

type {{ .Fname | ToTitle }} struct {}`


    return u.utils.CodeFileContents(tmpl)
}


func (u *golangUtils) DomainRepositoryFileContents() string {
    tmpl := `package {{ .Pkgname }}

type {{ .Fname | ToTitle }}Repository interface {}`

    return u.utils.CodeFileContents(tmpl)
}

func (u *golangUtils) InfraFileContents() string {
    tmpl := `package {{ .Pkgname }}

type {{ .Fname }}Repository struct {}

func New{{ .Fname | ToTitle }}Repository() repository.{{ .Fname | ToTitle }}Repository {
    return &{{ .Fname }}Repository{}
}`

    return u.utils.CodeFileContents(tmpl)
}


func (u *golangUtils) UseCaseFileContents() string {
    tmpl := `package {{ .Pkgname }}

type {{ .Fname | ToTitle }}UseCase interface {}

type {{ .Fname }}UseCase struct {
    repository.{{ .Fname | ToTitle }}Repository
}

func New{{ .Fname | ToTitle }}UseCase(r repository.{{ .Fname | ToTitle }}Repository) {{ .Fname | ToTitle }}UseCase {
    return &{{ .Fname }}UseCase{r}
}`

    return u.utils.CodeFileContents(tmpl)
}

func (u *golangUtils) PresentationFileContents() string {
    tmpl := `package {{ .Pkgname }}

type {{ .Fname | ToTitle }}{{ .Pkgname | ToTitle }} interface {}

type {{ .Fname }}{{ .Pkgname | ToTitle }} struct {
    usecase.{{ .Fname | ToTitle }}UseCase
}

func New{{ .Fname | ToTitle }}{{ .Pkgname | ToTitle }}(u usecase.{{ .Fname | ToTitle }}UseCase) {{ .Fname | ToTitle }}{{ .Pkgname | ToTitle }} {
    return &{{ .Fname }}{{ .Pkgname | ToTitle }}{u}
}`

	if u.Fname == "app" || u.Fname == "mod" {
		return u.DefaultFileContents()
	} else {
		return u.utils.CodeFileContents(tmpl)
	}
}

func (u *golangUtils) DiFileContents() string {
	tmpl := `package di

type DiContainer interface {}

type diContainer struct {}

func NewDiContainer() DiContainer {
	return &diContainer{}
}`

	return tmpl
}

func (u *golangUtils) DefaultFileContents() string {
    tmpl := `package {{ .Pkgname }}`
    return u.utils.CodeFileContents(tmpl)
}
