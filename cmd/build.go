package cmd

import (
	"fmt"
	"log"

	"github.com/asweed888/clerk/buildCmd"
	"github.com/asweed888/clerk/schema"
	"github.com/spf13/cobra"
)


var Build = &cobra.Command {
    Use: "build",
    Aliases: []string{"generate", "gen", "b", "g"},
    Short: "The package is generated based on the specifications written in clerk.yml.",
    RunE: func(cmd *cobra.Command, args []string) error {

        //clerk.ymlの読み込み
        scm, err := schema.Read()
        if err != nil {
            return err
        }
        switch scm.Lang {
        case "go":
            err = buildCmd.Clerk.Golang.Exec(scm)
        case "python":
            err = buildCmd.Clerk.Python.Exec(scm)
        case "deno":
            err = buildCmd.Clerk.Modernjs.Exec(scm)
        case "ecmascript":
            err = buildCmd.Clerk.Modernjs.Exec(scm)
        case "jsx":
            err = buildCmd.Clerk.Modernjs.Exec(scm)
        default:
            return fmt.Errorf("Invalid language designation")
        }

        if err != nil {
            return err
        }

        log.Println("generate of clerk has been completed.")
        return nil
    },
}
