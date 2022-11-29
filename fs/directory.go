/* <location: fs.directory />

指定されたディレクトリへの操作

*/
package fs

import "os"

type directory struct {}


func (s *directory) Create(dirpath string) error {
    if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
        return err
    }

    return nil
}
