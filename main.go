package main

import (
	"fmt"
	"os"

	"github.com/asweed888/clerk/cmd"
	"github.com/urfave/cli/v2"
)

func main(){

    app := cli.NewApp()
    app.Name = "clerk"
    app.Usage = cmd.BuildCommand.Usage
    app.Description = "Comming soon!"
    app.Flags = cmd.BuildCommand.Flags
    app.Version = "v0.0.0"
    app.Action = cmd.BuildCommand.Action
    app.Commands = []*cli.Command{
        cmd.BuildCommand,
    }


    if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

}
