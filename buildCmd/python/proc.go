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

    for _, clerk := range scm.Spec {
        location := clerk.Location
        modFilePath := fmt.Sprintf(
            "./%s/__init__.py",
            location,
        )

        // clerkRootのためのディレクトリを作成
        if err := create.Directory(location); err != nil {
            return err
        }

        // clerkRootのモジュールを作成
        err := create.Module(
            modFilePath,
            get.ModuleTemplate("clerkRoot"),
            map[string]interface{}{
                "Clerk": clerk,
            },
        )
        if err != nil {
            return err
        }

        for _, mod1 := range clerk.Modules {

            modFilePath := fmt.Sprintf(
                "./%s/%s/__init__.py",
                location,
                mod1.Name,
            )

            // mod1のためのディレクトリを作成
            if err := create.Directory(location, mod1.Name); err != nil {
                return err
            }

            // mod1を作成
            err := create.Module(
                modFilePath,
                get.ModuleTemplate("mod1"),
                map[string]interface{}{
                    "Location": location,
                    "Mod1": mod1,
                },
            )
            if err != nil {
                return err
            }

            for _, mod2 := range mod1.Upstreams {

                modFilePath := fmt.Sprintf(
                    "./%s/%s/%s/__init__.py",
                    location,
                    mod1.Name,
                    mod2.Name,
                )

                // mod2のためのディレクトリを作成
                if err := create.Directory(location, mod1.Name, mod2.Name); err != nil {
                    return err
                }

                // mod2を作成
                err := create.Module(
                    modFilePath,
                    get.ModuleTemplate("mod2"),
                    map[string]interface{}{
                        "Location": location,
                        "Mod1": mod1,
                        "Mod2": mod2,
                    },
                )
                if err != nil {
                    return err
                }

                for _, mod3 := range mod2.Upstreams {

                    modFilePath := fmt.Sprintf(
                        "./%s/%s/%s/%s.py",
                        location,
                        mod1.Name,
                        mod2.Name,
                        mod3.Name,
                    )

                    // moduleのファイルが存在していない場合
                    if _, err := os.Stat(modFilePath); err != nil {

                        err := create.Module(
                            modFilePath,
                            get.ModuleTemplate("mod3"),
                            map[string]interface{}{
                                "Location": location,
                                "Mod1": mod1,
                                "Mod2": mod2,
                                "Mod3": mod3,
                            },
                        )
                        if err != nil {
                            return nil
                        }

                    // moduleのファイルが存在している場合
                    } else {

                        err := update.Mod3(
                            modFilePath,
                            create.Mod3Comment(
                                location,
                                mod1.Name,
                                mod2.Name,
                                mod3.Name,
                                mod3.Comment,
                            ),
                        )
                        if err != nil {
                            return nil
                        }

                    } //end if

                } //end for

            } // end for

        } // end for
    }

    return nil
}
