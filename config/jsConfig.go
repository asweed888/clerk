/* <location: config.jsConfig />

modernjsの設定

*/
package config

import (
	"fmt"

	"github.com/asweed888/clerk/schema"
)

type c_jsConfig struct {}


type jsConfigContents struct {
    CodeFileName string
    CodeFileExt string
    ExportTo string
}

func (s *c_jsConfig) Get(scm *schema.ClerkYaml) (*jsConfigContents, error) {
    switch scm.Lang {
    case "deno":
        return &jsConfigContents{
            CodeFileName: "mod",
            CodeFileExt: "ts",
            ExportTo: "default ",
        }, nil
    case "denojs":
        return &jsConfigContents{
            CodeFileName: "mod",
            CodeFileExt: "js",
            ExportTo: "default ",
        }, nil
    case "nodejs":
        return &jsConfigContents{
            CodeFileName: "index",
            CodeFileExt: "js",
            ExportTo: "default ",
        }, nil
    case "nodets":
        return &jsConfigContents{
            CodeFileName: "index",
            CodeFileExt: "ts",
            ExportTo: "default ",
        }, nil
    case "jsx":
        return &jsConfigContents{
            CodeFileName: "index",
            CodeFileExt: "js",
            ExportTo: "",
        }, nil
    case "tsx":
        return &jsConfigContents{
            CodeFileName: "index",
            CodeFileExt: "ts",
            ExportTo: "",
        }, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
