package create

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func Directory(dirs ...string) error {
    if err := os.MkdirAll(strings.Join(dirs, "/"), os.ModePerm); err != nil {
        return err
    }
    return nil
}

func Module(outFilePath string, tmplTxt string, inputData interface{}) error {
    tmpl, err := template.New("clerk").Parse(tmplTxt)
    if err != nil {
        return err
    }
    file, err := os.Create(outFilePath)
    if err != nil {
        return err  //ファイルが開けなかったときエラー出力
    }
    defer file.Close()

    err = tmpl.Execute(file, inputData) // 置換して標準出力へ
	if err != nil {
		return fmt.Errorf(err.Error())
	}


	return nil
}

func EndModComment(location string, mod1Name string, mod2Name string, comment string) string {
    f := `""" <location: %s.%s.%s />

%s

"""`
    return fmt.Sprintf(f, location,  mod1Name, mod2Name, comment)
}
