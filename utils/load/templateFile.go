package load

import (
	"fmt"
	"io/ioutil"
)


func TemplateFile(modLevel string) func(string) (string, error) {
    return func (lang string) (string, error) {
        var tmplFilePath string
        switch lang {
            case "python":
                tmplFilePath = fmt.Sprintf("./templates/python/%s.tmpl", modLevel)
            default:
                return "", fmt.Errorf("ERROR: 無効な言語が指定されています。")
        }

        rawdata, err := ioutil.ReadFile(tmplFilePath)
        if err != nil {
            return "", fmt.Errorf("ERROR: %s.tmplを読み込めませんでした。", modLevel)
        }

        return string(rawdata), nil
    }
}
