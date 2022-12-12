/* <location: fs.directory />

指定されたディレクトリへの操作

*/
package fs

import "os"

type directoryMod struct {}
var Directory = &directoryMod{}


func (s *directoryMod) Create(dirpath string) error {
    if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
        return err
    }

    return nil
}
