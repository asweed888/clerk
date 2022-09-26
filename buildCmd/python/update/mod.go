package update

import (
	"io/ioutil"
	"os"
	"regexp"
)

func EndMod(filePath string, newComment string) error {
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


func EndModComment(filePath string, newComment string) error {
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


func EndModClerk(filePath string, newValue string) error {
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
