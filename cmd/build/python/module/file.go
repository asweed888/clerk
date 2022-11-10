package module

import (
	"fmt"
	"os"
	"text/template"
)


type FileSpec struct {}

var File = &FileSpec{}

func (s *FileSpec) Save(outFilePath string, tmplTxt string, inputData interface{}) error {
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
