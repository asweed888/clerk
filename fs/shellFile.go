/* <location: fs.shellFile />

実行可能なshellのためのファイル

*/
package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type shellFile struct {}


func (s *shellFile) Create(lang string, location string, upstream string, method string) error {
    outFilePath := fmt.Sprintf("./%s/Clerk/%s/%s", location, upstream, method)
    fileContents := fmt.Sprintf(`#!/bin/%s

echo "this is clerk's default return value"`, lang)

    if _, err := os.Stat(outFilePath); err != nil {
        perm := "0755"
        perm32, _ := strconv.ParseUint(perm, 8, 32)
        err = ioutil.WriteFile(outFilePath, []byte(fileContents), os.FileMode(perm32))
        if err != nil {
            return err
        }
    }

    return nil
}
