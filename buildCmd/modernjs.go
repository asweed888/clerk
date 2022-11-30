/* <location: buildCmd.modernjs />



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

type modernjs struct {}


func (s *modernjs) Exec(scm *schema.ClerkYaml) error {

    for _, lv0 := range scm.Spec {
        codeFilePath := fmt.Sprintf(
            "./%s/mod.go",
            lv0.Location,
        )

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

		// level0のコードファイルを書き出す
		if err := fs.Clerk.CodeFile.Write(
			codeFilePath,
			template.Golang.Lv0.FullTemplate(),
			map[string]interface{}{
				"Level0": lv0,
			},
		); err != nil { return err }


		for _, lv1 := range lv0.Upstream {
			codeFilePath := fmt.Sprintf(
				"./%s/%s.go",
				lv0.Location,
				lv1.Name,
			)

            if _, err := os.Stat(codeFilePath); err != nil {
            // moduleのファイルが存在していない場合

				// level1のコードファイルを書き出す
				if err := fs.Clerk.CodeFile.Write(
					codeFilePath,
					template.Golang.Lv1.FullTemplate(),
					map[string]interface{}{
						"Level0": lv0,
						"Level1": lv1,
					},
				); err != nil { return nil }
			} else {
            // moduleのファイルが存在している場合

				// コードファイルのコメント部分を更新する
				if err := fs.Clerk.CodeFile.Replace(
					codeFilePath,
					`/\* <location:[\s\S\n]*?\*/`,
					template.Utils.FillTemplate(
						template.Golang.Lv1.CommentTemplate(),
						map[string]interface{}{
							"Level0": lv0,
							"Level1": lv1,
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

			for _, method := range lv1.Methods {
				methodTemplate := template.Golang.Lv1.MethodTemplate()

                // コードファイル内でmethodが定義されていない場合
				if !fs.Clerk.CodeFileMethod.IsDefined(
					fileContent,
					fmt.Sprintf(
						"func (s *%s) %s(",
						lv1.Name,
						strings.Title(method),
					),
				) {
                    // コードファイルにmethodを追記する
					err = fs.Clerk.CodeFileMethod.Append(
						codeFilePath,
						methodTemplate,
						lv1.Name,
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
