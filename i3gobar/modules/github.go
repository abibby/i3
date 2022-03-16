package modules

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abibby/i3/i3gobar/icon"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	"robpike.io/filter"
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
		log.Printf("Failed to fetch notifications: %v\n", err)
		return ""
	}
	failedActions, err := failedWorkflows(ctx, client)
	if err != nil {
		log.Printf("Failed to fetch workflows: %v\n", err)
		return ""
	}

	notMergable, err := getNotMergableCount(ctx, client)
	if err != nil {
		log.Printf("Failed to fetch non mergable PRs: %v\n", err)
		return ""
	}

	return fmt.Sprintf("%s %d !%d m%d", icon.Github, len(notifs), failedActions, notMergable)
}

func failedWorkflows(ctx context.Context, client *github.Client) (int, error) {
	prs, _, err := client.PullRequests.List(ctx, "merotechnologies", "web-dashboard", &github.PullRequestListOptions{
		State: "open",
	})
	if err != nil {
		return 0, err
	}

	a, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, "merotechnologies", "web-dashboard", &github.ListWorkflowRunsOptions{
		Actor: "abibby",
	})
	if err != nil {
		return 0, err
	}

	type TimeStatus struct {
		time   time.Time
		status string
	}
	statusMap := map[string]TimeStatus{}
	for _, run := range a.WorkflowRuns {
		ts, ok := statusMap[run.GetHeadBranch()]
		if !ok || run.GetCreatedAt().Time.After(ts.time) {
			statusMap[run.GetHeadBranch()] = TimeStatus{
				time:   run.GetCreatedAt().Time,
				status: run.GetConclusion(),
			}
		}
	}

	filter.ChooseInPlace(&prs, func(pr *github.PullRequest) bool {
		return pr.GetUser().GetLogin() == "abibby"
	})

	failedActions := 0
	for _, pr := range prs {
		status, ok := statusMap[pr.GetHead().GetRef()]
		if ok && status.status == "failure" {
			failedActions++
		}
	}
	return failedActions, nil
}

func getNotMergableCount(ctx context.Context, client *github.Client) (int, error) {
	prs, _, err := client.PullRequests.List(ctx, "merotechnologies", "web-dashboard", &github.PullRequestListOptions{
		State: "open",
	})
	if err != nil {
		return 0, err
	}

	notMergable := 0
	for _, pr := range prs {
		if pr.GetUser().GetLogin() == "abibby" && !pr.GetMergeable() {
			notMergable++
		}
	}

	return notMergable, nil
}
