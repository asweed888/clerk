package config

import (
	"fmt"
	"strings"

	"github.com/asweed888/clerk/domain/model"
)


var Golang = &model.TacitConfig{
    Ext: "go",
    FileModeStr: "0644",
    ImplInitialFileContents: func(c *model.TacitConfig, path string, fname string) string {

        idx := strings.LastIndexByte(path, '/') + 1
        pkgname := path[idx:]

        return fmt.Sprintf("package %s", pkgname)
    },
}
