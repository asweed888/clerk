package model

import (
	"fmt"
)

type DeclareUpstream struct {
    Name string `yaml:"name"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
    TacitConfig *TacitConfig
}

func (d *DeclareUpstream) CreateDirectory(workdir string) error {
    return createDirectory(workdir)
}

func (d *DeclareUpstream) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Name)
}

func (d *DeclareUpstream) SetTacitConfig(conf *TacitConfig) {
    d.TacitConfig = conf
}
