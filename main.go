package main

import (
	"fmt"
	"github.com/urfave/cli"
	"my-git/git"
	"my-git/selectPrompt"
	"my-git/textPrompt"
	"os"
)

func branchesCmd(c *cli.Context) error {
	var options []selectPrompt.Option
	for _, branch := range git.Branches() {
		var option = selectPrompt.Option{Text: branch.Name, Value: branch.Name, Selected: branch.Currect}
		options = append(options, option)
	}
	options = append(options, selectPrompt.Option{Text: "Create new branch", Custom: true})

	selectedOption := selectPrompt.Init(options)

	if selectedOption.Custom {
		branch := textPrompt.Init("Please enter a name for your branch:")
		command := append(git.CheckoutNewBranch, branch)
		_, err := git.RunInlineCommand(command...)
		if err != nil {
			return err
		}
		fmt.Printf("Switched to new branch %q \n", branch)
	} else {
		branch := selectedOption.Value
		command := append(git.CheckoutBranch, branch)
		_, err := git.RunInlineCommand(command...)
		if err != nil {
			return err
		}
		fmt.Printf("Switched to branch %q \n", branch)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "my-git"
	app.Version = "0.0.1"
	app.Usage = "set of git conveniences"

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "branches",
			Aliases:     []string{"br"},
			Usage:       "manage branches",
			ArgsUsage:   "[D]",
			Description: "create/delete/checkout branches",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "checkout, c",
					Usage: "Checkout the branch",
				},
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "Delete the branch",
				},
			},
			Action: branchesCmd,
		},
	}

	app.Run(os.Args)
}
