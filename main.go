package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asweed888/clerk/cmd"
	"github.com/urfave/cli/v2"
)

func main(){

    rawdata, err := ioutil.ReadFile("templates/python/mod1.tmpl.py")
	if err != nil {
		panic(err)
	}

    println(string(rawdata))

    app := cli.NewApp()
    app.Name = "clerk"
    app.Usage = cmd.UpCommand.Usage
    app.Description = "Comming soon!"
    app.Flags = cmd.UpCommand.Flags
    app.Version = "v0.0.0"
    app.Action = cmd.UpCommand.Action
    app.Commands = []*cli.Command{
        cmd.UpCommand,
    }


    if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

}
