package tmpl

import (
	"fmt"
	"io/ioutil"
)


func Mod1Tmpl(lang string) (string, error) {

    var tmplFilePath string
    switch lang {
        case "python":
            tmplFilePath = "./templates/python/mod1.tmpl"
        default:
            return "", fmt.Errorf("ERROR: 無効な言語が指定されています。")
    }

    rawdata, err := ioutil.ReadFile(tmplFilePath)
    if err != nil {
        return "", fmt.Errorf("ERROR: mod1.tmplを読み込めませんでした。")
    }

    return string(rawdata), nil
}
