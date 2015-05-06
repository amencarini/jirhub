package commands

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/amencarini/jirhub/config"
	"github.com/amencarini/jirhub/jira"

	"github.com/skratchdot/open-golang/open"
)

func PullRequest(conf config.Configuration) {
	// Gets the current branch name
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	if err != nil {
		log.Fatal(err)
	}

	currentBranch := string(out)

	// Look for the typical Jira ticket number format ("ABC-123") in the branch name
	r, _ := regexp.Compile("([A-Z]+-[0-9]+)")
	key := r.FindString(currentBranch)

	ticket := jira.GetTicket(key, conf)

	targetBranch := conf.TargetBranch
	if targetBranch == "" {
		targetBranch = "master"
	}

	url, err := exec.Command("hub", "compare", "-u", fmt.Sprintf("%s...%s", targetBranch, currentBranch)).Output()

	title := fmt.Sprintf("%s %s", ticket.Key, ticket.Fields.Summary)

	browserUrl := fmt.Sprintf("%s?expand=1&pull_request[title]=%s&pull_request[body]=%s", strings.Trim(string(url), " \n"), title, ticket.Fields.Description)

	if err != nil {
		fmt.Println(out)
		log.Fatal(err)
	}

	open.Run(browserUrl)
}
