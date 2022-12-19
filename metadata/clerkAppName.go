/* <location: metadata.clerkAppName />

clerk.ymlがあるディレクトリ名

*/
package metadata

import (
	"os"
	"path/filepath"
)

type clerkAppNameMod struct {}
var ClerkAppName = &clerkAppNameMod{}


func (s *clerkAppNameMod) Get() (string, error) {
    wd, err := os.Getwd()
    if err != nil {
        return "", err
    }
    return filepath.Base(wd), err
}
