package main

import (
	"code/code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func baseAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() != 1 {
		return fmt.Errorf("missing operand")
	}

	path := cmd.Args().Get(0)

	isHumanFormat := cmd.Bool("human")
	isIncludeHiddens := cmd.Bool("all")

	res, err := code.GetPathSize(path, isIncludeHiddens)

	if err != nil {
		return err
	}

	fmt.Printf("%s\t%s\n", code.FormatSize(res, isHumanFormat), path)

	return nil
}

func main() {
	cmd := &cli.Command{
		Name: "hexlet-path-size",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human, H",
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    "all, a",
				Value:   false,
				Usage:   "include hidden files and directories",
				Aliases: []string{"a"},
			},
		},
		Usage:  "print size of a file or directory",
		Action: baseAction,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
