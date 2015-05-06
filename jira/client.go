package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/amencarini/jirhub/config"
)

func GetTicket(key string, conf config.Configuration) Ticket {
	ticket := Ticket{Key: key}

	credentials := strings.Join([]string{conf.Username, conf.Password}, ":")

	encryptedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	url := fmt.Sprintf("https://%s/rest/api/2/issue/%s", conf.JiraHost, ticket.Key)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encryptedCredentials))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &ticket)

	return ticket
}
