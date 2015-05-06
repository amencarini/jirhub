# JirHub

Tools for the Jira & GitHub developer, inspired by [git whistles](https://github.com/mezis/git-whistles).

## Install

Make sure you have hub installed:

```
brew install hub
```

Install this!

```
go get github.com/amencarini/jirhub
```

Then create a configuration file in your home dir with name `.jirhub-config.json`, and fill it as following:

```json
{
	"username": "amencarini",
	"password": "wow_muchs3cure",
	"targetBranch": "notMaster",
	"jiraHost": "yourcompany.atlassian.net"
}
```

Use `targetBranch` if you want a default target branch other than master.

Then when in a directory with the project associated to Jira, use `jirhub pr` and profit!
