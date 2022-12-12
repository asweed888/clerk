/* <location: fs.codeFile />



 */
package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type codeFileMod struct {}
var CodeFile = &codeFileMod{}


func (s *codeFileMod) Read(codeFilePath string) (string, error) {
    file, err := os.Open(codeFilePath)
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


func (s *codeFileMod) Write(codeFilePath string, tmplString string, inputData interface{}) error {
    funcMap := template.FuncMap{
        "ToTitle": strings.Title,
    }
    tmpl, err := template.New("clerk").Funcs(funcMap).Parse(tmplString)
    if err != nil {
        return err
    }
    file, err := os.Create(codeFilePath)
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


func (s *codeFileMod) Replace(codeFilePath string, fromString string, toString string) error {
    replaceTarget := regexp.MustCompile(fromString)
    bytes, err := ioutil.ReadFile(codeFilePath)
	if err != nil {
		return err
	}

    fileContents := string(bytes)
    fileContents = replaceTarget.ReplaceAllString(fileContents, toString)

    err = ioutil.WriteFile(codeFilePath, []byte(fileContents), os.ModePerm)
    if err != nil {
        return err
    }

    return nil
}
