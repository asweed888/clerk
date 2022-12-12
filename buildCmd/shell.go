/* <location: buildCmd.shell />



 */
package buildCmd

import (
	"fmt"
	"strings"

	"github.com/asweed888/clerk/fs"
	"github.com/asweed888/clerk/schema"
)

type shellMod struct {}
var Shell = &shellMod{}


func (s *shellMod) Exec(scm *schema.ClerkYaml) error {
    for _, lv0 := range scm.Spec {

		// locationのディレクトリを作成する
		if err := fs.Directory.Create(lv0.Location); err != nil { return err }

        // 作成されたディレクトリがclerkによって作成されたものである事がわかるように
        // .clerkというファイルを作成
		if err := fs.DotClerkFile.Create(lv0.Location); err != nil { return err }

        // location rootにコメントが記載された場合は
        // locationのディレクトリのみを作成してモジュール書き出し等の処理は行わない
        if lv0.Comment != "" {
            continue
        }


		for _, lv2 := range lv0.Upstream {
            for _, method := range lv2.Methods {

				// method fileの書き出し先のディレクトリを作成
				if err := fs.Directory.Create(
					fmt.Sprintf("%s/%s", lv0.Location, strings.Title(lv2.Name)),
				); err != nil { return err }

				fs.ShellFile.Create(scm.Lang, lv0.Location, strings.Title(lv2.Name), strings.Title(method))

            }
        }
    }

    return nil
}
