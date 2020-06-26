package modules

import (
	"context"
	"fmt"
	"log"

	"github.com/abibby/i3/i3gobar/icon"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

var ghClient *github.Client

func getGHClient(ctx context.Context) *github.Client {
	if ghClient == nil {
		token, err := pass("github-token")
		if err != nil {
			log.Printf("error getting github-token: %v\n", err)
		}
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		ghClient = github.NewClient(tc)
	}
	return ghClient
}

func GitHubNotifications() string {
	ctx := context.Background()
	client := getGHClient(ctx)

	notifs, _, err := client.Activity.ListNotifications(ctx, nil)
	if err != nil {
		log.Printf("%v\n", err)
		return ""
	}

	return fmt.Sprintf("%s %d", icon.Github, len(notifs))
}
