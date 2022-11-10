package module1

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type FileSpec struct {
    Comment F_CommentSpec
    Clerk F_ClerkSpec
    Method F_MethodSpec
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


type F_CommentSpec struct {}
func (s *F_CommentSpec) Save(filePath string, newComment string) error {
    replaceTarget := regexp.MustCompile(`""" <location:[\s\S\n]*?"""`)
    bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

    fileContents := string(bytes)
    fileContents = replaceTarget.ReplaceAllString(fileContents, newComment)

    err = ioutil.WriteFile(filePath, []byte(fileContents), os.ModePerm)
    if err != nil {
        return err
    }

    return nil
}


type F_ClerkSpec struct {}
func (s *F_ClerkSpec) Save(filePath string, newValue string) error {
    replaceTarget := regexp.MustCompile(`def clerk[\s\S\n]*# end clerk`)
    bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

    fileContents := string(bytes)
    fileContents = replaceTarget.ReplaceAllString(fileContents, newValue)

    err = ioutil.WriteFile(filePath, []byte(fileContents), os.ModePerm)
    if err != nil {
        return err
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


func (s *F_MethodSpec) IsDefined(fileContent string, method string) bool {
    return strings.Contains(fileContent, fmt.Sprintf("def _%s(", method))
}
