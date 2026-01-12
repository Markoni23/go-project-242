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
	res, err := code.GetPathSize(path)

	if err != nil {
		return err
	}

	fmt.Printf("%s\t%s\n", res, path)

	return nil
}

func main() {
	cmd := &cli.Command{
		Name:   "hexlet-path-size",
		Usage:  "print size of a file or directory",
		Action: baseAction,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
