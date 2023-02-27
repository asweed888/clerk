package main

import (
	"github.com/asweed888/clerk/app"
	"github.com/spf13/cobra"
)

var Version string


func main(){

    a := app.NewApp("./clerk.yml")
    subcmd := a.NewSubCommand()


    cli := &cobra.Command{
        Use: "clerk",
        Short: "This is a very simple declarative development framework.",
        Version: Version,
    }

    cli.AddCommand(subcmd.BuildCmd())
    cli.Execute()
}
