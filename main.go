package main

import (
	"github.com/amencarini/jirhub/commands"

	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jirhub"
	app.Usage = "The Jira/GitHub Developer Swiss Army Knife"
	app.Author = "Alessandro Mencarini"
	app.Email = "alessandro.mencarini@gmail.com"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:      "pull-request",
			ShortName: "pr",
			Usage:     "Opens a GitHub pull request for the current branch",
			Action: func(c *cli.Context) {
				commands.PullRequest()
			},
		},
		{
			Name:      "select",
			ShortName: "s",
			Usage:     "Switchs to a branch, local or remote, given a Jira ticket ID",
			Action: func(c *cli.Context) {
				commands.Select()
			},
		},
	}

	app.Run(os.Args)
}
