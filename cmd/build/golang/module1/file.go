package module1

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type FileSpec struct {
    Comment F_CommentSpec
    Method F_MethodSpec
}

var File = &FileSpec{}


type F_CommentSpec struct {}
func (s *F_CommentSpec) Save(filePath string, newComment string) error {
    replaceTarget := regexp.MustCompile(`/\* <location:[\s\S\n]*\*/`)
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


type F_MethodSpec struct {}
func (s *F_MethodSpec) Save(modFilePath string, template string, methods ...string) error {
    methodDefine := fmt.Sprintf(template, methods[0], methods[1])
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
