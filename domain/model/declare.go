package model

type Declare struct {
    Lang string `yaml:"lang"`
    Spec []*DeclareSpec
    TacitConfig *TacitConfig
}

type TacitConfig struct {
    Ext string
    FilePermission string
}
