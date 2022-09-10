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
            "./clerk/%s/__init__.py",
            mod0.Location,
        )

        // mod0のためのディレクトリを作成
        if err := create.Directory("clerk", mod0.Location); err != nil {
            return err
        }

        // mod0のmoduleを作成
        if err := create.Module(
            modFilePath,
            get.ModuleTemplate("mod0"),
            map[string]interface{}{
                "Mod0": mod0,
            },
        ); err != nil { return err }

        for _, mod1 := range mod0.Modules {

            modFilePath := fmt.Sprintf(
                "./clerk/%s/%s/__init__.py",
                mod0.Location,
                mod1.Name,
            )

            // mod1のためのディレクトリを作成
            if err := create.Directory("clerk", mod0.Location, mod1.Name); err != nil {
                return err
            }

            // mod1を作成
            if err := create.Module(
                modFilePath,
                get.ModuleTemplate("mod1"),
                map[string]interface{}{
                    "Mod0": mod0,
                    "Mod1": mod1,
                },
            ); err != nil {
                return err
            }

            for _, mod2 := range mod1.Upstreams {

                modFilePath := fmt.Sprintf(
                    "./clerk/%s/%s/%s.py",
                    mod0.Location,
                    mod1.Name,
                    mod2.Name,
                )

                // mod2のためのディレクトリを作成
                if err := create.Directory("clerk", mod0.Location, mod1.Name, mod2.Name); err != nil {
                    return err
                }

                // moduleのファイルが存在していない場合
                if _, err := os.Stat(modFilePath); err != nil {

                    if err := create.Module(
                        modFilePath,
                        get.ModuleTemplate("mod2"),
                        map[string]interface{}{
                            "Mod0": mod0,
                            "Mod1": mod1,
                            "Mod2": mod2,
                        },
                    ); err != nil {
                        return nil
                    }

                // moduleのファイルが存在している場合
                } else {

                    if err := update.EndMod(
                        modFilePath,
                        create.EndModComment(
                            mod0.Location,
                            mod1.Name,
                            mod2.Name,
                            mod2.Comment,
                        ),
                    ); err != nil {
                        return nil
                    }

                } //end if

                // mod2を作成
                if err := create.Module(
                    modFilePath,
                    get.ModuleTemplate("mod2"),
                    map[string]interface{}{
                        "Mod0": mod0,
                        "Mod1": mod1,
                        "Mod2": mod2,
                    },
                ); err != nil {
                    return err
                }

            } // end for

        } // end for
    }

    return nil
}
