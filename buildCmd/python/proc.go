package python

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/buildCmd/python/create"
	"github.com/asweed888/clerk/buildCmd/python/get"
	"github.com/asweed888/clerk/buildCmd/python/update"
	"github.com/asweed888/clerk/schema"
)

func Proc(scm *schema.ClerkYaml) error {

    for _, mod0 := range scm.Spec {
        modFilePath := fmt.Sprintf(
            "./clerks/%s/__init__.py",
            mod0.Location,
        )

        // mod0のためのディレクトリを作成
        if err := create.Directory("clerks", mod0.Location); err != nil {
            return err
        }

        // mod0のmoduleを作成
        if err := create.Module(
            modFilePath,
            get.ModuleTemplate("mod0"),
            map[string]interface{}{
                "ApiRoot": get.ApiRoot(scm.Export),
                "Mod0": mod0,
            },
        ); err != nil { return err }

        for _, mod1 := range mod0.Upstream {
            modFilePath := fmt.Sprintf(
                "./clerks/%s/%s.py",
                mod0.Location,
                mod1.Name,
            )

            if _, err := os.Stat(modFilePath); err != nil {

                if err := create.Module(
                    modFilePath,
                    get.ModuleTemplate("mod1"),
                    map[string]interface{}{
                        "Mod0": mod0,
                        "Mod1": mod1,
                    },
                ); err != nil {
                    return nil
                }

            // moduleのファイルが存在している場合
            } else {

                if err := update.EndModComment(
                    modFilePath,
                    get.CompletedTemplate(
                        get.ModuleTemplate("mod1_comment"),
                        map[string]interface{}{
                            "Mod0": mod0,
                            "Mod1": mod1,
                        },
                    ),
                ); err != nil {
                    return nil
                }

                if err := update.EndModClerk(
                    modFilePath,
                    get.CompletedTemplate(
                        get.ModuleTemplate("mod1_clerk"),
                        map[string]interface{}{
                            "Mod1": mod1,
                        },
                    ),
                ); err != nil {
                    return nil
                }

            } //end if


        } //end for mod1

    } //end for mod0

    return nil
}
