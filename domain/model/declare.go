package model

import (
	"os"
)

type Declare struct {
    Lang string `yaml:"lang"`
    Arch string `yaml:"arch"`
    Spec []*DeclareSpec
    TacitConfig *TacitConfig
}


func createDirectory(workdir string) error {
    if err := os.MkdirAll(workdir, os.ModePerm); err != nil { return err }

    return nil
}
