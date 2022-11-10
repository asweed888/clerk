package python

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/cmd/build/python/module"
	"github.com/asweed888/clerk/cmd/build/python/module0"
	"github.com/asweed888/clerk/cmd/build/python/module1"
	"github.com/asweed888/clerk/schema"
)

func Proc(scm *schema.ClerkYaml) error {

    for _, mod0 := range scm.Spec {
        modFilePath := fmt.Sprintf(
            "./%s/__init__.py",
            mod0.Location,
        )

        // mod0のためのディレクトリを作成
        if err := module0.Directory.Create(mod0.Location); err != nil {
            return err
        }

        // location rootにコメントが記載された場合は
        // locationのディレクトリのみを作成してモジュール書き出し等の処理は行わない
        if mod0.Comment != "" {
            continue
        }

        // mod0のmoduleを作成
        if err := module.File.Save(
            modFilePath,
            module0.Template.Get(),
            map[string]interface{}{
                "Mod0": mod0,
            },
        ); err != nil { return err }


        for _, mod1 := range mod0.Upstream {
            modFilePath := fmt.Sprintf(
                "./%s/%s.py",
                mod0.Location,
                mod1.Name,
            )

            if _, err := os.Stat(modFilePath); err != nil {
            // moduleのファイルが存在していない場合

                if err := module.File.Save(
                    modFilePath,
                    module1.Template.Get(),
                    map[string]interface{}{
                        "Mod0": mod0,
                        "Mod1": mod1,
                    },
                ); err != nil { return nil }

            } else {
            // moduleのファイルが存在している場合

                if err := module1.File.Comment.Save(
                    modFilePath,
                    module.Template.Fill(
                        module1.Template.Comment.Get(),
                        map[string]interface{}{
                            "Mod0": mod0,
                            "Mod1": mod1,
                        },
                    ),
                ); err != nil {
                    return nil
                }

                if err := module1.File.Clerk.Save(
                    modFilePath,
                    module.Template.Fill(
                        module1.Template.Clerk.Get(),
                        map[string]interface{}{
                            "Mod1": mod1,
                        },
                    ),
                ); err != nil {
                    return nil
                }

            } //end if

            // メソッドの自動書き出し機能
            // 未定義のメソッドをファイル行末に追加
            fileContent, err := module1.File.Load(modFilePath)
            if err != nil {
                return err
            }

            for _, method := range mod1.Methods {
                methodTemplate := module1.Template.Method.Get()

                if !module1.File.Method.IsDefined(fileContent, method) {
                    err = module1.File.Method.Save(modFilePath, methodTemplate, method)
                    if err != nil {
                        return err
                    }
                }
            }

        } //end for mod1

    } //end for mod0

    return nil
}
