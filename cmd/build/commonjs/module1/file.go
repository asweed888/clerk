package module1

import (
	"io/ioutil"
	"os"
	"regexp"
)

type FileSpec struct {
    Comment F_CommentSpec
    Clerk F_ClerkSpec
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


type F_ClerkSpec struct {}
func (s *F_ClerkSpec) Save(filePath string, newValue string) error {
    replaceTarget := regexp.MustCompile(`module\.exports\.clerk[\s\S\n]*\/\/ end clerk`)
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
