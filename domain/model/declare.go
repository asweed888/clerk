package model

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []DeclareSpec
    CodeFile []DeclareCodeFileSpec
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

type DeclareCodeFileSpec struct {
    Label string `yaml:"label"`
    Struct []struct {
        Name string `yaml:"name"`
    }
    Interface []struct {
        Name string `yaml:"name"`
    }
    Function []struct {
        Name string `yaml:"name"`
    }
}
