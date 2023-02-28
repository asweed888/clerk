package config

import "github.com/asweed888/clerk/domain/model"


var Bash = &model.TacitConfig{
    FileModeStr: "0755",
    ImplInitialFileContents: func(c *model.TacitConfig, path string, fname string) string {
        return "#!/bin/bash"
    },
}
