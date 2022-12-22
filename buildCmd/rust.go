/* <location: buildCmd.rust />



 */
package buildCmd

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/fs"
	"github.com/asweed888/clerk/schema"
	"github.com/asweed888/clerk/template"
)

type rustMod struct {}
var Rust = &rustMod{}


func (s *rustMod) Exec(scm *schema.ClerkYaml) error {

    for _, lv0 := range scm.Spec {
        codeFilePath := fmt.Sprintf(
            "./%s/mod.rs",
            lv0.Location,
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

        // export: trueの場合
        if scm.Export {
            if err := fs.CodeFile.Write(
                "./mod.rs",
                template.Rust.Lv0.RootExportFile(),
                map[string]interface{}{
                    "Spec": scm.Spec,
                },
            ); err != nil { return err }
        }

        // level0のコードファイルを書き出す
        if err := fs.CodeFile.Write(
            codeFilePath,
            template.Rust.Lv0.ExportFile(),
            map[string]interface{}{
                "Level0": lv0,
            },
        ); err != nil { return err }


		for _, lv1 := range lv0.Upstream {
            codeFilePath := fmt.Sprintf(
                "./%s/%s.rs",
                lv0.Location,
                lv1.Name,
            )

            if _, err := os.Stat(codeFilePath); err != nil {
            // moduleのファイルが存在していない場合

				// level1のコードファイルを書き出す
				if err := fs.CodeFile.Write(
					codeFilePath,
					template.Rust.Lv1.CommentTemplate(),
					map[string]interface{}{
						"Level0": lv0,
						"Level1": lv1,
					},
				); err != nil { return err }

            } else {
            // moduleのファイルが存在している場合

				// コードファイルのコメント部分を更新する
				if err := fs.CodeFile.Replace(
					codeFilePath,
					`/\* <location:[\s\S\n]*?\*/`,
					template.Utils.FillTemplate(
						template.Rust.Lv1.CommentTemplate(),
						map[string]interface{}{
							"Level0": lv0,
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
				methodTemplate := template.Rust.Lv1.MethodTemplate()

                // コードファイル内でmethodが定義されていない場合
				if !fs.CodeFileMethod.IsDefined(
					fileContent,
					fmt.Sprintf(
						"pub fn %s(",
						method,
					),
				) {
                    // コードファイルにmethodを追記する
					err = fs.CodeFileMethod.Append(
						codeFilePath,
						methodTemplate,
						method,
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
