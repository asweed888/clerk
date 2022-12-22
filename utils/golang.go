/* <location: utils.golang />



*/
package utils

import (
	"strings"
)

type golangMod struct {}
var Golang = &golangMod{}


func (s *golangMod) Exportable(isExport bool, modName string) string {
    if isExport {
        return strings.Title(modName)
    } else {
        return modName
    }
}
