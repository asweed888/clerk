package config

import (
	"fmt"
	"regexp"

	"github.com/asweed888/clerk/domain/model"
)


var Golang = &model.TacitConfig{
    Ext: "go",
    FileModeStr: "0644",
    ImplInitialFileContents: func(c *model.TacitConfig, path string, fname string) string {
        re := regexp.MustCompile(`[^/]+$`)
        pkgname := re.FindString(path)
        return fmt.Sprintf("package %s", pkgname)
    },
}
