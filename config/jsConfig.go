/* <location: config.jsConfig />

modernjsの設定

*/
package config

import (
	"fmt"

	"github.com/asweed888/clerk/schema"
)

type jsConfig struct {}


type jsConfigContents struct {
    CodeFileName string
    ExportTo string
}

func (s *jsConfig) Get(scm *schema.ClerkYaml) (*jsConfigContents, error) {
    switch scm.Lang {
    case "deno":
        return &jsConfigContents{
            CodeFileName: "mod.js",
            ExportTo: "default ",
        }, nil
    case "ecmascript":
        return &jsConfigContents{
            CodeFileName: "index.js",
            ExportTo: "default ",
        }, nil
    case "jsx":
        return &jsConfigContents{
            CodeFileName: "index.js",
            ExportTo: "",
        }, nil
    default:
        return nil, fmt.Errorf("Invalid language designation")
    }
}
