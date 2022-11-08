package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Mod1Struct struct {
    Method Mod1Method
}

type Mod1Method struct {}

var Mod1 = &Mod1Struct{}

func (m *Mod1Struct) Read(modFilePath string) (string, error) {
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


func (m *Mod1Struct) Write(modFilePath string, template string, method string) error {
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


func (m *Mod1Method) IsDefined(fileContent string, method string) bool {
    return strings.Contains(fileContent, fmt.Sprintf("def _%s(", method))
}
