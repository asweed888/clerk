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
            err = buildCmd.Golang.Exec(scm)
        case "rust":
            err = buildCmd.Rust.Exec(scm)
        case "python":
            err = buildCmd.Python.Exec(scm)
        case "deno":
            err = buildCmd.Modernjs.Exec(scm)
        case "denojs":
            err = buildCmd.Modernjs.Exec(scm)
        case "nodejs":
            err = buildCmd.Modernjs.Exec(scm)
        case "nodets":
            err = buildCmd.Modernjs.Exec(scm)
        case "jsx":
            err = buildCmd.Modernjs.Exec(scm)
        case "tsx":
            err = buildCmd.Modernjs.Exec(scm)
        case "bash":
            err = buildCmd.Shell.Exec(scm)
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
