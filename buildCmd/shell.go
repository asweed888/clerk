/* <location: buildCmd.shell />



 */
package buildCmd

import (
	"fmt"
	"strings"

	"github.com/asweed888/clerk/fs"
	"github.com/asweed888/clerk/schema"
)

type c_shell struct {}


func (s *c_shell) Exec(scm *schema.ClerkYaml) error {
    for _, lv0 := range scm.Spec {

		// locationのディレクトリを作成する
		if err := fs.Clerk.Directory.Create(lv0.Location); err != nil { return err }

        // 作成されたディレクトリがclerkによって作成されたものである事がわかるように
        // .clerkというファイルを作成
		if err := fs.Clerk.DotClerkFile.Create(lv0.Location); err != nil { return err }

        // location rootにコメントが記載された場合は
        // locationのディレクトリのみを作成してモジュール書き出し等の処理は行わない
        if lv0.Comment != "" {
            continue
        }


		for _, lv2 := range lv0.Upstream {
            for _, method := range lv2.Methods {

				// method fileの書き出し先のディレクトリを作成
				if err := fs.Clerk.Directory.Create(
					fmt.Sprintf("%s/Clerk/%s", lv0.Location, strings.Title(lv2.Name)),
				); err != nil { return err }

				fs.Clerk.ShellFile.Create(scm.Lang, lv0.Location, strings.Title(lv2.Name), strings.Title(method))

            }
        }
    }

    return nil
}
