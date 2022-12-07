/* <location: fs.codeFileMethod />



*/
package fs

import (
	"fmt"
	"os"
	"strings"
)

type codeFileMethod struct {}


func (s *codeFileMethod) IsDefined(codeFileContent string, searchString string) bool {
    return strings.Contains(codeFileContent, searchString)
}


func (s *codeFileMethod) Append(codeFilePath string, template string, methods ...string) error {
    tmp := make([]interface{}, len(methods))
    for i, val := range methods {
        tmp[i] = val
    }
    methodDefine := fmt.Sprintf(template, tmp...)
    file, err := os.OpenFile(codeFilePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    _, err = fmt.Fprintln(file, methodDefine)
    if err != nil {
        return err
    }

    return nil
}
