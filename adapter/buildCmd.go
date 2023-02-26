package adapter

import (
	"github.com/asweed888/clerk/usecase"
	"github.com/spf13/cobra"
)

type BuildCommand interface {
    BuildCmd() *cobra.Command
}

type buildCommand struct {
    DeclareUseCase usecase.DeclareUseCase
}

func NewBuildCommand(u usecase.DeclareUseCase) BuildCommand {
    return &buildCommand{u}
}

func (c *buildCommand) BuildCmd() *cobra.Command {
    return &cobra.Command{
        Use: "build",
        Aliases: []string{"generate", "gen", "b", "g"},
        Short: "The package is generated based on the specifications written in clerk.yml.",
        RunE: func(cmd *cobra.Command, args []string) error {
            c.DeclareUseCase.GenerateArch()
            return nil
        },
    }
}
