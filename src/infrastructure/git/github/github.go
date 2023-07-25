package github

import "encoding/json"

type Github struct {
}

func (g Github) ProccessWebhook(payload []byte) (dataStruct any, err error) {
	var jsonData GitHubWebhookPullRequest

	err = json.Unmarshal(payload, &jsonData)

	return jsonData, err
}

func (g Github) Driver() string {
	return "github"
}

func (g Github) GetUserData(userCode string) (dataStruct any, err error) {
	return nil, nil
}
