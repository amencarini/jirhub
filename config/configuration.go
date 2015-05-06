package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Configuration struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	TargetBranch string `json:"targetBranch"`
	JiraHost     string `json:"jiraHost"`
}

func GetConfiguration() Configuration {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(usr.HomeDir, ".jirhub-config.json"))

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)

	c := Configuration{}
	err = decoder.Decode(&c)

	if err != nil {
		fmt.Println("error:", err)
	}

	return c
}
