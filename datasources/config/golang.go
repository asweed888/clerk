package config

import (
	"fmt"
	"strings"

	"github.com/asweed888/clerk/domain/model"
)

type golang struct {
    Config *model.TacitConfig
    Path string
    Fname string
}

var Golang = &model.TacitConfig{
    Ext: "go",
    FileModeStr: "0644",
    ImplInitialFileContents: func(c *model.TacitConfig, path string, fname string) string {
        nativeData := golang{c, path, fname}

        if isDomainModel(c, path) {
            return nativeData.DomainModelFileContents()
        } else if isDomainRepository(c, path) {
            return nativeData.DomainRepositoryFileContents()
        }

        return fmt.Sprintf("package %s", pkgname(path))
    },
}


func (s *golang) DomainModelFileContents() string {
    tmpl := `package %s

type %s struct {}`

    return fmt.Sprintf(tmpl, pkgname(s.Path), strings.Title(s.Fname))
}


func (s *golang) DomainRepositoryFileContents() string {
    tmpl := `package %s

type %sRepository interface {}`

    return fmt.Sprintf(tmpl, pkgname(s.Path), strings.Title(s.Fname))
}
