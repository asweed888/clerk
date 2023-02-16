package model

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []*DeclareSpec
    InternalConfig *InternalConfig
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

type InternalConfig struct {
    Ext string
    FilePermission string
}
