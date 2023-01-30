package model

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []DeclareSpec
}

type DeclareSpec struct {
    Location string `yaml:"location"`
    Upstream []DeclareUpstream
    CodeFile []DeclareCodeFile
}

type DeclareUpstream struct {
    Name string `yaml:"name"`
    Upstream []DeclareUpstream
    CodeFile []DeclareCodeFile
}

type DeclareCodeFile struct {
    Name string `yaml:"name"`
    Label string `yaml:"label"`
    Description string `yaml:"description"`
}
