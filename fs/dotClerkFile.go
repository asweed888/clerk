/* <location: fs.dotClerkFile />



*/
package fs

import (
	"fmt"
	"os"
)

type dotClerkFileMod struct {}
var DotClerkFile = &dotClerkFileMod{}


func (s *dotClerkFileMod) Create(dirpath string) error {
    outFilePath := fmt.Sprintf("./%s/.clerk", dirpath)

    if _, err := os.Stat(outFilePath); err != nil {
        file, err := os.Create(outFilePath)
        if err != nil {
            return err
        }
        defer file.Close()
    }

    return nil
}
