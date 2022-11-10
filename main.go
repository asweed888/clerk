package main

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/cmd"
	"github.com/urfave/cli/v2"
)

var Version string

func main(){

    app := cli.NewApp()
    app.Name = "clerk"
    app.Usage = cmd.Build.Usage
    app.Description = "This is a very simple declarative development framework."
    app.Flags = cmd.Build.Flags
    app.Version = Version
    app.Commands = []*cli.Command{
        cmd.Build,
    }

    if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

}
