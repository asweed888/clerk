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
    Declare *Declare
}

func (d DeclareCodeFile) CreateCodeFile(workdir string) error {
    path := fmt.Sprintf("%s/%s", workdir, d.Name)
    if _, err := os.Stat(path); err != nil {
        err = ioutil.WriteFile(path, []byte(""), d.createCodeFilePermission())
        if err != nil { return err }

    }
    return nil
}

func (d DeclareCodeFile) createCodeFilePermission() fs.FileMode {
    config := d.Declare.TacitConfig
    perm32, _ := strconv.ParseUint(config.FilePermission, 8, 32)
    return os.FileMode(perm32)
}
