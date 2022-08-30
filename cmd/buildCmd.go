package cmd

import (
	"fmt"

	_create "github.com/asweed888/clerk/utils/create"
	_get "github.com/asweed888/clerk/utils/get"
	_load "github.com/asweed888/clerk/utils/load"
	"github.com/urfave/cli/v2"
)

// type Mod1TmplArgs struct {
//     Name
//
// }

var BuildCommand = &cli.Command{
    Name: "build",
    Usage: "build clerk",
    Flags: []cli.Flag{},
    Action: func(c *cli.Context) error {

        //clerk.ymlの読み込み
        // clerkYml, _ := schema.ClerkYaml()

        schema, _ := _load.Schema()
        for _, mod1 := range schema.Spec {

            // mod1の階層のディレクトリを作成
            if err := _create.ModuleDirectory(mod1.Name); err != nil {
                return err
            }

            tmplTxt, err := _load.TemplateFile("mod1")(schema.Lang)
            if err != nil {
                return err
            }

            err = _create.Module(
                fmt.Sprintf("./clerk/%s/%s", mod1.Name, _get.ModuleMainFileName(schema.Lang)),
                tmplTxt,
                mod1,
            )
            if err != nil {
                return err
            }

            for _, mod2 := range mod1.Upstreams {
                // mod2の階層のディレクトリを作成
                if err := _create.ModuleDirectory(mod1.Name, mod2.Name); err != nil {
                    return fmt.Errorf("ディレクトリの作成に失敗")
                }

                tmplTxt, err := _load.TemplateFile("mod2")(schema.Lang)
                if err != nil {
                    return fmt.Errorf("テンプレートファイルの読み込みに失敗")
                }
                err = _create.Module(
                    fmt.Sprintf("./clerk/%s/%s/%s", mod1.Name, mod2.Name, _get.ModuleMainFileName(schema.Lang)),
                    tmplTxt,
                    map[string]interface{}{
                        "Mod1": mod1,
                        "Mod2": mod2,
                    },
                )
                if err != nil {
                    return fmt.Errorf("%s::モジュールの作成に失敗", err)
                }
                
            }

        }
        return nil
    },
}
