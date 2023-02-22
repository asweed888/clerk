package cmd

import (
	"github.com/asweed888/clerk/usecase"
	"github.com/spf13/cobra"
)

type BuildCmd interface {
    GenerateArch(cmd *cobra.Command, args []string) error
}

type buildCmd struct {
    DeclareUseCase usecase.DeclareUseCase
}

func NewBuildCmd(u usecase.DeclareUseCase) BuildCmd {
    return &buildCmd{u}
}

func (c *buildCmd) GenerateArch(cmd *cobra.Command, args []string) error {
    return nil
}
