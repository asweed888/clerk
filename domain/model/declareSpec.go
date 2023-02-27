package model

import (
	"fmt"
	"os"
)

type DeclareSpec struct {
    Location string `yaml:"location"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
    TacitConfig *TacitConfig
}

func (d *DeclareSpec) CreateDirectory(workdir string) error {
    if err := os.MkdirAll(workdir, os.ModePerm); err != nil { return err }

    return nil
}

func (d *DeclareSpec) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Location)
}

func (d *DeclareSpec) SetTacitConfig(conf *TacitConfig) {
    d.TacitConfig = conf
}
