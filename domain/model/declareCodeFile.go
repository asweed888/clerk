package model

import (
	"io/ioutil"
	"os"
)

type DeclareCodeFile struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    TacitConfig *TacitConfig
}

func (d *DeclareCodeFile) CreateCodeFile(workdir string) error {
    path := d.TacitConfig.CodeFileFullPath(workdir, d.Name)
    ifc := d.TacitConfig.InitialFileContents(workdir, d.Name)

    if _, err := os.Stat(path); err != nil {
        err = ioutil.WriteFile(path, []byte(ifc), d.TacitConfig.FileMode())
        if err != nil { return err }

    }
    return nil
}


func (d *DeclareCodeFile) SetTacitConfig(conf *TacitConfig) {
    d.TacitConfig = conf
}
