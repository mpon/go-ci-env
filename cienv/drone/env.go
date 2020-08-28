package drone

import (
	"fmt"
	"strconv"
)

type Client struct {
	Getenv func(string) string
}

func (client Client) Match() bool {
	return client.Getenv("DRONE") != ""
}

func (client Client) RepoOwner() string {
	return client.Getenv("DRONE_REPO_OWNER")
}

func (client Client) RepoName() string {
	return client.Getenv("DRONE_REPO_NAME")
}

func (client Client) RepoPath() string {
	return client.RepoOwner() + "/" + client.RepoName()
}

func (client Client) SHA1() string {
	return client.Getenv("DRONE_COMMIT_SHA1")
}

func (client Client) IsPR() bool {
	return client.Getenv("DRONE_PULL_REQUEST") != ""
}

func (client Client) PRNumber() (int, error) {
	pr := client.Getenv("DRONE_PULL_REQUEST")
	if pr == "" {
		return -1, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("DRONE_PULL_REQUEST is invalid. It is failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}