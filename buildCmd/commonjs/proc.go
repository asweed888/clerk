package commonjs

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/buildCmd/commonjs/create"
	"github.com/asweed888/clerk/buildCmd/commonjs/get"
	"github.com/asweed888/clerk/buildCmd/commonjs/update"
	"github.com/asweed888/clerk/schema"
)

func Proc(scm *schema.ClerkYaml) error {

    var modFilePath string

    for _, mod1 := range scm.Spec {

        modFilePath = "./clerk/%s/index.js"

        // mod1の階層のディレクトリを作成
        if err := create.ModuleDirectory(mod1.Name); err != nil {
           return err
        }

        err := create.Module(
            fmt.Sprintf(modFilePath, mod1.Name),
            get.ModuleTemplate("mod1"),
            mod1,
        )
        if err != nil {
            return err
        }

        for _, mod2 := range mod1.Upstreams {

            modFilePath = "./clerk/%s/%s/index.js"

            if err := create.ModuleDirectory(mod1.Name, mod2.Name); err != nil {
               return err
            }

            err = create.Module(
                fmt.Sprintf(modFilePath, mod1.Name, mod2.Name),
                get.ModuleTemplate("mod2"),
                map[string]interface{}{
                    "Mod1": mod1,
                    "Mod2": mod2,
                },
            )
            if err != nil {
                return fmt.Errorf("%s\nモジュールの作成に失敗", err)
            }

            for _, mod3 := range mod2.Upstreams {

                modFilePath = "./clerk/%s/%s/%s.js"

                if _, err := os.Stat(
                    fmt.Sprintf(modFilePath, mod1.Name, mod2.Name, mod3.Name),
                ); err != nil {

                    err = create.Module(
                        fmt.Sprintf(modFilePath, mod1.Name, mod2.Name, mod3.Name),
                        get.ModuleTemplate("mod3"),
                        map[string]interface{}{
                            "Mod1": mod1,
                            "Mod2": mod2,
                            "Mod3": mod3,
                        },
                    )
                    if err != nil {
                        return fmt.Errorf("%s\nモジュールの作成に失敗", err)
                    }

                } else {
                    err = update.Mod3(
                        fmt.Sprintf(modFilePath, mod1.Name, mod2.Name, mod3.Name),
                        create.Mod3Comment(mod1.Name, mod2.Name, mod3.Name, mod3.Comment),
                    )
                    if err != nil {
                        return fmt.Errorf("%s", err)
                    }

                }

            }
        }
    }

    return nil
}
