package main

import (
	"github.com/asweed888/clerk/interactor"
	"github.com/spf13/cobra"
)

var Version string


func main(){

    i := interactor.NewInteractor("./clerk.yml")
    subcmd := i.NewSubCommand()


    app := &cobra.Command{
        Use: "clerk",
        Short: "This is a very simple declarative development framework.",
        Version: Version,
    }

    app.AddCommand(subcmd.BuildCmd())
    app.Execute()
}
