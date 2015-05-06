package jira

type Ticket struct {
	Key    string `json:"key"`
	Fields struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"fields"`
}
