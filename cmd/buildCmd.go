package cmd

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/schema"
	"github.com/urfave/cli/v2"
)


var UpCommand = &cli.Command{
    Name: "build",
    Usage: "build clerk",
    Flags: []cli.Flag{},
    Action: func(c *cli.Context) error {

        //clerk.ymlの読み込み
        clerkYml, _ := schema.ClerkYaml()
        for _, mod1 := range clerkYml.Spec {

            // mod1の階層のディレクトリを作成
            if err := os.MkdirAll(fmt.Sprintf("clerk/%s", mod1.Name), os.ModePerm); err != nil {
                return err
            }


        }
        return nil
    },
}
