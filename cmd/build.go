package cmd

import (
	"fmt"
	"log"

	"github.com/asweed888/clerk/cmd/build/commonjs"
	"github.com/asweed888/clerk/cmd/build/deno"
	"github.com/asweed888/clerk/cmd/build/ecmascript"
	"github.com/asweed888/clerk/cmd/build/python"
	"github.com/asweed888/clerk/schema"
	"github.com/urfave/cli/v2"
)

var Build = &cli.Command{
    Name: "build",
    Aliases: []string{"generate", "gen", "b", "g"},
    Usage: "The package is generated based on the specifications written in clerk.yml.",
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
            case "deno":
                err = deno.Proc(scm)
            case "commonjs":
                err = commonjs.Proc(scm)
            case "ecmascript":
                err = ecmascript.Proc(scm)
            default:
                return fmt.Errorf("対応していない言語が指定されています。")
        }

        if err != nil {
            return err
        }

        log.Println("generate of clerk has been completed.")
        return nil
    },
}
