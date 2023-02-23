package model

import "fmt"

type DeclareUpstream struct {
    Name string `yaml:"name"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
}

func (d DeclareUpstream) CreateDirectory(workdir string)

func (d DeclareUpstream) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Name)
}
