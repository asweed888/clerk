package model

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
)

type DeclareCodeFile struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    TacitConfig *TacitConfig
}

func (d *DeclareCodeFile) CreateCodeFile(workdir string) error {
    path := fmt.Sprintf("%s/%s.%s", workdir, d.Name, d.TacitConfig.Ext)
    if _, err := os.Stat(path); err != nil {
        err = ioutil.WriteFile(path, []byte(""), d.createCodeFilePermission())
        if err != nil { return err }

    }
    return nil
}

func (d *DeclareCodeFile) createCodeFilePermission() fs.FileMode {
    perm32, _ := strconv.ParseUint(d.TacitConfig.FilePermission, 8, 32)
    return os.FileMode(perm32)
}

func (d *DeclareCodeFile) SetTacitConfig(conf *TacitConfig) {
    d.TacitConfig = conf
}
