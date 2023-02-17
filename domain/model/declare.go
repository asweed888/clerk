package model

import "fmt"

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []*DeclareSpec
    TacitConfig *TacitConfig
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

type TacitConfig struct {
    Ext string
    FilePermission string
}



/* methods */
func (d DeclareSpec) CreateDirectory(path string) string {
    p := fmt.Sprintf("%s/%s", path, d.Location)
    return p
}

func (d DeclareUpstream) CreateDirectory(path string) string {
    p := fmt.Sprintf("%s/%s", path, d.Name)
    return p
}
