package main

import (
	"github.com/asweed888/clerk/cmd"
	"github.com/spf13/cobra"
)

var Version string

var commands = []*cobra.Command{
    cmd.Build,
    cmd.Watch,
}

func main(){

    app := &cobra.Command{
        Use: "clerk",
        Short: "This is a very simple declarative development framework.",
        Version: Version,
    }

    app.AddCommand(commands...)
    app.Execute()
}
