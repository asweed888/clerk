package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type FileSpec struct {
    Module1 struct {
        Method F_MethodSpec
    }
    DotClerk DotClerkSpec
}

var File = &FileSpec{}
func (s *FileSpec) Load(modFilePath string) (string, error) {
    file, err := os.Open(modFilePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    b, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }

    return string(b), nil
}

func (s *FileSpec) Save(outFilePath string, tmplTxt string, inputData interface{}) error {
    funcMap := template.FuncMap{
        "ToTitle": strings.Title,
    }
    tmpl, err := template.New("clerk").Funcs(funcMap).Parse(tmplTxt)
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

type F_MethodSpec struct {}
func (s *F_MethodSpec) Save(modFilePath string, template string, method string) error {
    methodDefine := fmt.Sprintf(template, method)
    file, err := os.OpenFile(modFilePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    _, err = fmt.Fprintln(file, methodDefine)
    if err != nil {
        return err
    }

    return nil
}


func (s *F_MethodSpec) IsDefined(fileContent string, searchString string) bool {
    return strings.Contains(fileContent, searchString)
}


type DotClerkSpec struct {}
func (s *DotClerkSpec) Create(location string) error {
    outFilePath := fmt.Sprintf("./%s/.clerk", location)

    if _, err := os.Stat(outFilePath); err != nil {
        file, err := os.Create(outFilePath)
        if err != nil {
            return err
        }
        defer file.Close()
    }

    return nil
}
