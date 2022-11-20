package common

import (
	"fmt"

	"github.com/asweed888/clerk/schema"
)

type JsConfigContents struct {
    ModuleMainFile string
    ExportTo string
}

func JsConfig(scm *schema.ClerkYaml) (*JsConfigContents, error) {
    switch scm.Lang {
    case "deno":
        return &JsConfigContents{
            ModuleMainFile: "mod.js",
            ExportTo: "default ",
        }, nil
    case "ecmascript":
        return &JsConfigContents{
            ModuleMainFile: "index.js",
            ExportTo: "default ",
        }, nil
    case "jsx":
        return &JsConfigContents{
            ModuleMainFile: "index.js",
            ExportTo: "",
        }, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
