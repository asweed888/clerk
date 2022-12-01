/* <location: buildCmd.python />



 */
package buildCmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/asweed888/clerk/fs"
	"github.com/asweed888/clerk/schema"
	"github.com/asweed888/clerk/template"
)

type python struct {}


func (s *python) Exec(scm *schema.ClerkYaml) error {

    for _, lv0 := range scm.Spec {
        codeFilePath := fmt.Sprintf(
            "./%s/__init__.py",
            lv0.Location,
        )
        codeFilePath2 := fmt.Sprintf(
            "./%s/Clerk/__init__.py",
            lv0.Location,
        )

		// locationのディレクトリを作成する
		if err := fs.Clerk.Directory.Create(
            fmt.Sprintf("%s/Clerk", lv0.Location),
        ); err != nil { return err }

        // 作成されたディレクトリがclerkによって作成されたものである事がわかるように
        // .clerkというファイルを作成
		if err := fs.Clerk.DotClerkFile.Create(lv0.Location); err != nil { return err }

        // location rootにコメントが記載された場合は
        // locationのディレクトリのみを作成してモジュール書き出し等の処理は行わない
        if lv0.Comment != "" {
            continue
        }

		// level0のコードファイルを書き出す
		if err := fs.Clerk.CodeFile.Write(
			codeFilePath,
			template.Python.Lv0.FullTemplate(),
			map[string]interface{}{},
		); err != nil { return err }

		// level1のコードファイルを書き出す
		if err := fs.Clerk.CodeFile.Write(
			codeFilePath2,
			template.Python.Lv1.FullTemplate(),
			map[string]interface{}{
				"Level0": lv0,
			},
		); err != nil { return err }


		for _, lv2 := range lv0.Upstream {
			codeFilePath := fmt.Sprintf(
				"./%s/Clerk/%s.py",
				lv0.Location,
				strings.Title(lv2.Name),
			)

            if _, err := os.Stat(codeFilePath); err != nil {
            // moduleのファイルが存在していない場合

				// level1のコードファイルを書き出す
				if err := fs.Clerk.CodeFile.Write(
					codeFilePath,
					template.Python.Lv2.FullTemplate(),
					map[string]interface{}{
						"Level0": lv0,
						"Level2": lv2,
					},
				); err != nil { return nil }
			} else {
            // moduleのファイルが存在している場合

				// コードファイルのコメント部分を更新する
				if err := fs.Clerk.CodeFile.Replace(
					codeFilePath,
					`""" <location:[\s\S\n]*?"""`,
					template.Utils.FillTemplate(
						template.Python.Lv2.CommentTemplate(),
						map[string]interface{}{
							"Level0": lv0,
							"Level2": lv2,
						},
					),
				); err != nil { return err }
			} //end if

            // メソッドの自動書き出し機能
            // 未定義のメソッドをファイル行末に追加
            fileContent, err := fs.Clerk.CodeFile.Read(codeFilePath)
            if err != nil {
                return err
            }

			for _, method := range lv2.Methods {
				methodTemplate := template.Python.Lv2.MethodTemplate()

                // コードファイル内でmethodが定義されていない場合
				if !fs.Clerk.CodeFileMethod.IsDefined(
					fileContent,
					fmt.Sprintf(
						"def %s(",
						strings.Title(method),
					),
				) {
                    // コードファイルにmethodを追記する
					err = fs.Clerk.CodeFileMethod.Append(
						codeFilePath,
						methodTemplate,
						strings.Title(method),
					)
					if err != nil {
						return err
					}
				}
			}

		}
    }

    return nil
}
