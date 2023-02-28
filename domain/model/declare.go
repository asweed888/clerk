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
    TacitConfig *TacitConfig
}

type TacitConfig struct {
    Lang string
    Ext string
    FileModeStr string
    ImplInitialFileContents func(conf *TacitConfig, path string, fname string) string
}


func (c *TacitConfig) FileMode() fs.FileMode {
    perm32, _ := strconv.ParseUint(c.FileModeStr, 8, 32)
    return os.FileMode(perm32)
}

func (c *TacitConfig) CodeFileFullPath(path string, fname string) string {
    if c.Ext == "" {
        return fmt.Sprintf("%s/%s", path, fname)
    } else {
        return fmt.Sprintf("%s/%s.%s", path, fname, c.Ext)
    }
}

func (c *TacitConfig) InitialFileContents(path string, fname string) string {
    return c.ImplInitialFileContents(c, path, fname)
}

func createDirectory(workdir string) error {
    if err := os.MkdirAll(workdir, os.ModePerm); err != nil { return err }

    return nil
}
