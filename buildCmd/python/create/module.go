package create

import (
	"fmt"
	"os"
	"text/template"
)


func ModuleDirectory(dirs ...string) error {
    if len(dirs) >= 3 {
        if err := os.MkdirAll(fmt.Sprintf("%s/%s/%s", dirs[0], dirs[1], dirs[2]), os.ModePerm); err != nil {
            return err
        }
    } else {
        if err := os.MkdirAll(fmt.Sprintf("%s/%s", dirs[0], dirs[1]), os.ModePerm); err != nil {
            return err
        }
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
