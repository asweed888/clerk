package cmd

import (
	"fmt"

	"github.com/asweed888/clerk/buildCmd/commonjs"
	"github.com/asweed888/clerk/buildCmd/python"
	"github.com/asweed888/clerk/schema"
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
        scm, err := schema.Load()
        if err != nil {
            return err
        }
        switch scm.Lang {
            case "python":
                err = python.Proc(scm)
            case "commonjs":
                err = commonjs.Proc(scm)
            default:
                return fmt.Errorf("対応していない言語が指定されています。")
        }

        if err != nil {
            return err
        }

        return nil
    },
}
