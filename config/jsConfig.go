/* <location: config.jsConfig />

modernjsの設定

*/
package config

import (
	"fmt"

	"github.com/asweed888/clerk/metadata"
	"github.com/asweed888/clerk/schema"
)

type jsConfigMod struct {}
var JsConfig = &jsConfigMod{}


type jsConfigContents struct {
    AppName string
    CodeFileName string
    CodeFileExt string
    ExportTo string
    RootExportTo string
}


func exportTo(isExport bool, defaultValue string) string {
    if isExport {
        return "default "
    } else {
        return defaultValue
    }
}


func (s *jsConfigMod) Get(scm *schema.ClerkYaml) (*jsConfigContents, error) {

    an, err := metadata.ClerkAppName.Get()
    if err != nil {
        return &jsConfigContents{}, err
    }
    switch scm.Lang {
    case "deno":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "mod",
            CodeFileExt: "ts",
            ExportTo: "default ",
            RootExportTo: "default ",
        }, nil
    case "denojs":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "mod",
            CodeFileExt: "js",
            ExportTo: "default ",
            RootExportTo: "default ",
        }, nil
    case "nodejs":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "index",
            CodeFileExt: "js",
            ExportTo: "default ",
            RootExportTo: "default ",
        }, nil
    case "nodets":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "index",
            CodeFileExt: "ts",
            ExportTo: "default ",
            RootExportTo: "default ",
        }, nil
    case "jsx":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "index",
            CodeFileExt: "js",
            ExportTo: exportTo(scm.Export, ""),
            RootExportTo: "",
        }, nil
    case "tsx":
        return &jsConfigContents{
            AppName: an,
            CodeFileName: "index",
            CodeFileExt: "ts",
            ExportTo: exportTo(scm.Export, ""),
            RootExportTo: "",
        }, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
