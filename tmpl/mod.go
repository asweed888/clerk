package tmpl

import (
	"fmt"
	"io/ioutil"
)


func ModTmpl(level string) func(string) (string, error) {
    return func (lang string) (string, error) {
        var tmplFilePath string
        switch lang {
            case "python":
                tmplFilePath = fmt.Sprintf("./templates/python/%s.tmpl", level)
            default:
                return "", fmt.Errorf("ERROR: 無効な言語が指定されています。")
        }

        rawdata, err := ioutil.ReadFile(tmplFilePath)
        if err != nil {
            return "", fmt.Errorf("ERROR: mod1.tmplを読み込めませんでした。")
        }

        return string(rawdata), nil
    }
}
