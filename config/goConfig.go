/* <location: config.goConfig />

goの設定

*/
package config

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/asweed888/clerk/metadata"
)

type goConfigMod struct {}
var GoConfig = &goConfigMod{}


type goConfigContents struct {
    AppName string
    GoModName string
}


func goModName() (string, error) {
    f, _ := os.Open("../go.mod")
	bu := bufio.NewReaderSize(f, 1024)
    line, _, err := bu.ReadLine()
    if err == io.EOF {
        return "", err
    }

    return strings.Replace(string(line), "module ", "", 1), nil
}


func (s *goConfigMod) Get() (*goConfigContents, error){
    an, err := metadata.ClerkAppName.Get()
    if err != nil {
        return &goConfigContents{}, err
    }

    gmn, err := goModName()
    if err != nil {
        return &goConfigContents{}, err
    }

    return &goConfigContents{
        AppName: an,
        GoModName: gmn,
    }, nil
}
