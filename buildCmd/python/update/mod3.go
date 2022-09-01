package update

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func Mod3(filePath string, newComment string) error {
    replaceTarget := regexp.MustCompile(`""" <location:[\s\S\n]*"""`)
    bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

    fileContents := string(bytes)
    fmt.Println(fileContents)
    fileContents = replaceTarget.ReplaceAllString(fileContents, newComment)
    fmt.Println(fileContents)
    
    err = ioutil.WriteFile(filePath, []byte(fileContents), os.ModePerm)
    if err != nil {
        return err
    }

    return nil
}
