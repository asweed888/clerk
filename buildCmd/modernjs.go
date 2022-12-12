/* <location: buildCmd.modernjs />



 */
package buildCmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/asweed888/clerk/config"
	"github.com/asweed888/clerk/fs"
	"github.com/asweed888/clerk/schema"
	"github.com/asweed888/clerk/template"
)

type modernjsMod struct {}
var Modernjs = &modernjsMod{}


func (s *modernjsMod) Exec(scm *schema.ClerkYaml) error {

    jsConfig, err := config.JsConfig.Get(scm)
    if err != nil { return err }

    for _, lv0 := range scm.Spec {
        codeFilePath := fmt.Sprintf(
            "./%s/%s.%s",
            lv0.Location,
            jsConfig.CodeFileName,
            jsConfig.CodeFileExt,
        )

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

		// level0のコードファイルを書き出す
		if err := fs.CodeFile.Write(
			codeFilePath,
			template.Modernjs.Lv0.FullTemplate(),
			map[string]interface{}{
				"Level0": lv0,
                "JsConfig": jsConfig,
			},
		); err != nil { return err }


		for _, lv1 := range lv0.Upstream {
			codeFilePath := fmt.Sprintf(
				"./%s/%s.%s",
				lv0.Location,
				lv1.Name,
                jsConfig.CodeFileExt,
			)

            if _, err := os.Stat(codeFilePath); err != nil {
            // moduleのファイルが存在していない場合

				// level1のコードファイルを書き出す
				if err := fs.CodeFile.Write(
					codeFilePath,
					template.Modernjs.Lv1.FullTemplate(),
					map[string]interface{}{
						"Level0": lv0,
						"Level1": lv1,
					},
				); err != nil { return nil }
			} else {
            // moduleのファイルが存在している場合

				// コードファイルのコメント部分を更新する
				if err := fs.CodeFile.Replace(
					codeFilePath,
					`/\* <location:[\s\S\n]*?\*/`,
					template.Utils.FillTemplate(
						template.Modernjs.Lv1.CommentTemplate(),
						map[string]interface{}{
							"Level0": lv0,
							"Level1": lv1,
						},
					),
				); err != nil { return err }

				// コードファイルのexport部分を更新する
				if err := fs.CodeFile.Replace(
					codeFilePath,
					`export default[\s\S\n]*?\/\/ end export`,
					template.Utils.FillTemplate(
						template.Modernjs.Lv1.ExportTemplate(),
						map[string]interface{}{
							"Level1": lv1,
						},
					),
				); err != nil { return err }

			} //end if

            // メソッドの自動書き出し機能
            // 未定義のメソッドをファイル行末に追加
            fileContent, err := fs.CodeFile.Read(codeFilePath)
            if err != nil {
                return err
            }

			for _, method := range lv1.Methods {
				methodTemplate := template.Modernjs.Lv1.MethodTemplate()

                // コードファイル内でmethodが定義されていない場合
				if !fs.CodeFileMethod.IsDefined(
					fileContent,
					fmt.Sprintf(
						"function %s(",
						strings.Title(method),
					),
				) {
                    // コードファイルにmethodを追記する
					err = fs.CodeFileMethod.Append(
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
