package load

import (
	"fmt"
	"io/ioutil"
)


func TemplateFile(modLevel string) (string, error) {
    bytes, err := ioutil.ReadFile(fmt.Sprintf("./templates/python/%s.tmpl", modLevel))
    if err != nil {
        return "", fmt.Errorf("ERROR: %s.tmplを読み込めませんでした。", modLevel)
    }
    return string(bytes), nil
}
