package model

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []*DeclareSpec
}

type DeclareSpec struct {
    Location string `yaml:"location"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
}

type DeclareUpstream struct {
    Name string `yaml:"name"`
    Upstream []*DeclareUpstream
    CodeFile []*DeclareCodeFile
}

type DeclareCodeFile struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
}


/* methods */

// DeclareSpec
func (d DeclareSpec) CreateDirectory(workdir string)

func (d DeclareSpec) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Location)
}

// DeclareUpstream
func (d DeclareUpstream) CreateDirectory(workdir string)

func (d DeclareUpstream) ChangeDirectory(prevWorkDir string) string {
    return fmt.Sprintf("%s/%s", prevWorkDir, d.Name)
}


// DeclareCodeFile
func (d DeclareCodeFile) CreateCodeFile(workdir string)

func (d DeclareCodeFile) createCodeFilePermission() fs.FileMode {
    perm32, _ := strconv.ParseUint(d.TacitConfig.FilePermission, 8, 32)
    return os.FileMode(perm32)
}
