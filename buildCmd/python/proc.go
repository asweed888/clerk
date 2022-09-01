package python

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/buildCmd/python/create"
	"github.com/asweed888/clerk/buildCmd/python/load"
	"github.com/asweed888/clerk/buildCmd/python/update"
	"github.com/asweed888/clerk/schema"
)

func Proc(scm *schema.ClerkYaml) error {

    var modFilePath string

    for _, mod1 := range scm.Spec {

        modFilePath = "./clerk/%s/__init__.py"
        
        // mod1の階層のディレクトリを作成
        if err := create.ModuleDirectory(mod1.Name); err != nil {
           return err
        }

        tmplTxt, err := load.TemplateFile("mod1")
        if err != nil {
            return err
        }

        err = create.Module(
            fmt.Sprintf(modFilePath, mod1.Name),
            tmplTxt,
            mod1,
        )

        for _, mod2 := range mod1.Upstreams {
            
            modFilePath = "./clerk/%s/%s/__init__.py"
            
            if err := create.ModuleDirectory(mod1.Name, mod2.Name); err != nil {
               return err
            }

            tmplTxt, err := load.TemplateFile("mod2")
            if err != nil {
                return err
            }

            err = create.Module(
                fmt.Sprintf(modFilePath, mod1.Name, mod2.Name),
                tmplTxt,
                map[string]interface{}{
                    "Mod1": mod1,
                    "Mod2": mod2,
                },
            )
            if err != nil {
                return fmt.Errorf("%s\nモジュールの作成に失敗", err)
            }

            for _, mod3 := range mod2.Upstreams {

                modFilePath = "./clerk/%s/%s/%s.py"

                if _, err := os.Stat(
                    fmt.Sprintf(modFilePath, mod1.Name, mod2.Name, mod3.Name),
                ); err != nil {
                    tmplTxt, err := load.TemplateFile("mod3")
                    if err != nil {
                        return err
                    }

                    err = create.Module(
                        fmt.Sprintf(modFilePath, mod1.Name, mod2.Name, mod3.Name),
                        tmplTxt,
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
