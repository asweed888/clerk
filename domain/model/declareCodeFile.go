package model

import (
	"io/fs"
	"os"
	"strconv"
)

type DeclareCodeFile struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Declare *Declare
}

func (d DeclareCodeFile) CreateCodeFile(workdir string)

func (d DeclareCodeFile) createCodeFilePermission() fs.FileMode {
    config := d.Declare.TacitConfig
    perm32, _ := strconv.ParseUint(config.FilePermission, 8, 32)
    return os.FileMode(perm32)
}
