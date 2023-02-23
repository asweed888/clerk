package model

import "fmt"

type DeclareSpec struct {
    Location string `yaml:"location"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
}

func (d DeclareSpec) CreateDirectory(workdir string)

func (d DeclareSpec) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Location)
}
