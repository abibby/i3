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
		log.Printf("%v\n", err)
		return ""
	}

	prs, _, err := client.PullRequests.List(ctx, "merotechnologies", "web-dashboard", &github.PullRequestListOptions{
		State: "open",
	})
	if err != nil {
		log.Printf("%v\n", err)
		return ""
	}

	a, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, "merotechnologies", "web-dashboard", &github.ListWorkflowRunsOptions{
		Actor: "abibby",
	})

	type TimeStatus struct {
		time   time.Time
		status string
	}
	m := map[string]TimeStatus{}
	for _, run := range a.WorkflowRuns {
		ts, ok := m[run.GetHeadBranch()]
		if !ok || run.GetCreatedAt().Time.After(ts.time) {
			m[run.GetHeadBranch()] = TimeStatus{
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
		status, ok := m[pr.GetHead().GetRef()]
		if ok && status.status != "success" {
			failedActions++
		}
	}

	return fmt.Sprintf("%s %d !%d", icon.Github, len(notifs), failedActions)
}

// (*github.PullRequest)(0xc00036c000)(github.PullRequest{
// 	ID:499160072,
// 	Number:815,
// 	State:"open",
// 	Locked:false,
// 	Title:"WEB-611 fix subdomain when non tenant scoped",
// 	Body:"really hacky fix for adding tenant subdomains to routes

// I forgot that in web we always have magic tenant applied, and when working on backend we don't have that.

// I was unable to find where and how org's slug is applied to routes when working thru web routes. Anyone knows?
// ",
// 	CreatedAt:time.Time{wall:, ext:},
// 	UpdatedAt:time.Time{wall:, ext:},
// 	Labels:[
// 		github.Label{ID:2266512410, URL:"https://api.github.com/repos/merotechnologies/web-dashboard/labels/size/M", Name:"size/M", Color:"7F7203", Default:false, NodeID:"MDU6TGFiZWwyMjY2NTEyNDEw"}
// 	],
// 	User:github.User{Login:"dkhorev", ID:21098285, NodeID:"MDQ6VXNlcjIxMDk4Mjg1", AvatarURL:"https://avatars1.githubusercontent.com/u/21098285?v=4", HTMLURL:"https://github.com/dkhorev", GravatarID:"", Type:"User", SiteAdmin:false, URL:"https://api.github.com/users/dkhorev", EventsURL:"https://api.github.com/users/dkhorev/events{/privacy}", FollowingURL:"https://api.github.com/users/dkhorev/following{/other_user}", FollowersURL:"https://api.github.com/users/dkhorev/followers", GistsURL:"https://api.github.com/users/dkhorev/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/dkhorev/orgs", ReceivedEventsURL:"https://api.github.com/users/dkhorev/received_events", ReposURL:"https://api.github.com/users/dkhorev/repos", StarredURL:"https://api.github.com/users/dkhorev/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/dkhorev/subscriptions"}, Draft:false, MergeCommitSHA:"bbfd7faf59e6379edc820f6061140943907273ee", URL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815", HTMLURL:"https://github.com/merotechnologies/web-dashboard/pull/815", IssueURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/815", StatusesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/statuses/d676461c2e1373b51bbd159a61cea5883088cd38", DiffURL:"https://github.com/merotechnologies/web-dashboard/pull/815.diff", PatchURL:"https://github.com/merotechnologies/web-dashboard/pull/815.patch", CommitsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815/commits", CommentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/815/comments", ReviewCommentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815/comments", ReviewCommentURL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/comments{/number}", Assignees:[], AuthorAssociation:"CONTRIBUTOR", NodeID:"MDExOlB1bGxSZXF1ZXN0NDk5MTYwMDcy", RequestedReviewers:[], RequestedTeams:[github.Team{ID:3377978, NodeID:"MDQ6VGVhbTMzNzc5Nzg=", Name:"Web Dev", Description:"", URL:"https://api.github.com/organizations/42920552/team/3377978", Slug:"web-dev", Permission:"pull", Privacy:"closed", MembersURL:"https://api.github.com/organizations/42920552/team/3377978/members{/member}", RepositoriesURL:"https://api.github.com/organizations/42920552/team/3377978/repos"}], Links:github.PRLinks{Self:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815"}, HTML:github.PRLink{HRef:"https://github.com/merotechnologies/web-dashboard/pull/815"}, Issue:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/815"}, Comments:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/815/comments"}, ReviewComments:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815/comments"}, ReviewComment:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/comments{/number}"}, Commits:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls/815/commits"}, Statuses:github.PRLink{HRef:"https://api.github.com/repos/merotechnologies/web-dashboard/statuses/d676461c2e1373b51bbd159a61cea5883088cd38"}}, Head:github.PullRequestBranch{Label:"merotechnologies:dk-WEB-611-fix-subdomain-when-non-tenant-scoped", Ref:"dk-WEB-611-fix-subdomain-when-non-tenant-scoped", SHA:"d676461c2e1373b51bbd159a61cea5883088cd38", Repo:github.Repository{ID:147126789, NodeID:"MDEwOlJlcG9zaXRvcnkxNDcxMjY3ODk=", Owner:github.User{Login:"merotechnologies", ID:42920552, NodeID:"MDEyOk9yZ2FuaXphdGlvbjQyOTIwNTUy", AvatarURL:"https://avatars2.githubusercontent.com/u/42920552?v=4", HTMLURL:"https://github.com/merotechnologies", GravatarID:"", Type:"Organization", SiteAdmin:false, URL:"https://api.github.com/users/merotechnologies", EventsURL:"https://api.github.com/users/merotechnologies/events{/privacy}", FollowingURL:"https://api.github.com/users/merotechnologies/following{/other_user}", FollowersURL:"https://api.github.com/users/merotechnologies/followers", GistsURL:"https://api.github.com/users/merotechnologies/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/merotechnologies/orgs", ReceivedEventsURL:"https://api.github.com/users/merotechnologies/received_events", ReposURL:"https://api.github.com/users/merotechnologies/repos", StarredURL:"https://api.github.com/users/merotechnologies/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/merotechnologies/subscriptions"}, Name:"web-dashboard", FullName:"merotechnologies/web-dashboard", Description:"📉 Removing the uncertainty from building maintenance", Homepage:"https://mero.co", DefaultBranch:"sprint/october19", CreatedAt:github.Timestamp{2018-09-02 22:36:03 +0000 UTC}, PushedAt:github.Timestamp{2020-10-07 12:18:54 +0000 UTC}, UpdatedAt:github.Timestamp{2020-10-06 20:06:04 +0000 UTC}, HTMLURL:"https://github.com/merotechnologies/web-dashboard", CloneURL:"https://github.com/merotechnologies/web-dashboard.git", GitURL:"git://github.com/merotechnologies/web-dashboard.git", SSHURL:"git@github.com:merotechnologies/web-dashboard.git", SVNURL:"https://github.com/merotechnologies/web-dashboard", Language:"PHP", Fork:false, ForksCount:0, OpenIssuesCount:88, StargazersCount:2, WatchersCount:2, Size:12329, Archived:false, Disabled:false, Private:true, HasIssues:true, HasWiki:true, HasPages:false, HasProjects:true, HasDownloads:true, URL:"https://api.github.com/repos/merotechnologies/web-dashboard", ArchiveURL:"https://api.github.com/repos/merotechnologies/web-dashboard/{archive_format}{/ref}", AssigneesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/assignees{/user}", BlobsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/blobs{/sha}", BranchesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/branches{/branch}", CollaboratorsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/collaborators{/collaborator}", CommentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/comments{/number}", CommitsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/commits{/sha}", CompareURL:"https://api.github.com/repos/merotechnologies/web-dashboard/compare/{base}...{head}", ContentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/contents/{+path}", ContributorsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/contributors", DeploymentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/deployments", DownloadsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/downloads", EventsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/events", ForksURL:"https://api.github.com/repos/merotechnologies/web-dashboard/forks", GitCommitsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/commits{/sha}", GitRefsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/refs{/sha}", GitTagsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/tags{/sha}", HooksURL:"https://api.github.com/repos/merotechnologies/web-dashboard/hooks", IssueCommentURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/comments{/number}", IssueEventsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/events{/number}", IssuesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues{/number}", KeysURL:"https://api.github.com/repos/merotechnologies/web-dashboard/keys{/key_id}", LabelsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/labels{/name}", LanguagesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/languages", MergesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/merges", MilestonesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/milestones{/number}", NotificationsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/notifications{?since,all,participating}", PullsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls{/number}", ReleasesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/releases{/id}", StargazersURL:"https://api.github.com/repos/merotechnologies/web-dashboard/stargazers", StatusesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/statuses/{sha}", SubscribersURL:"https://api.github.com/repos/merotechnologies/web-dashboard/subscribers", SubscriptionURL:"https://api.github.com/repos/merotechnologies/web-dashboard/subscription", TagsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/tags", TreesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/trees{/sha}", TeamsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/teams"}, User:github.User{Login:"merotechnologies", ID:42920552, NodeID:"MDEyOk9yZ2FuaXphdGlvbjQyOTIwNTUy", AvatarURL:"https://avatars2.githubusercontent.com/u/42920552?v=4", HTMLURL:"https://github.com/merotechnologies", GravatarID:"", Type:"Organization", SiteAdmin:false, URL:"https://api.github.com/users/merotechnologies", EventsURL:"https://api.github.com/users/merotechnologies/events{/privacy}", FollowingURL:"https://api.github.com/users/merotechnologies/following{/other_user}", FollowersURL:"https://api.github.com/users/merotechnologies/followers", GistsURL:"https://api.github.com/users/merotechnologies/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/merotechnologies/orgs", ReceivedEventsURL:"https://api.github.com/users/merotechnologies/received_events", ReposURL:"https://api.github.com/users/merotechnologies/repos", StarredURL:"https://api.github.com/users/merotechnologies/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/merotechnologies/subscriptions"}}, Base:github.PullRequestBranch{Label:"merotechnologies:sprint/october19", Ref:"sprint/october19", SHA:"07925d9f1b20596c8bf3499e600798bb1885c4ce", Repo:github.Repository{ID:147126789, NodeID:"MDEwOlJlcG9zaXRvcnkxNDcxMjY3ODk=", Owner:github.User{Login:"merotechnologies", ID:42920552, NodeID:"MDEyOk9yZ2FuaXphdGlvbjQyOTIwNTUy", AvatarURL:"https://avatars2.githubusercontent.com/u/42920552?v=4", HTMLURL:"https://github.com/merotechnologies", GravatarID:"", Type:"Organization", SiteAdmin:false, URL:"https://api.github.com/users/merotechnologies", EventsURL:"https://api.github.com/users/merotechnologies/events{/privacy}", FollowingURL:"https://api.github.com/users/merotechnologies/following{/other_user}", FollowersURL:"https://api.github.com/users/merotechnologies/followers", GistsURL:"https://api.github.com/users/merotechnologies/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/merotechnologies/orgs", ReceivedEventsURL:"https://api.github.com/users/merotechnologies/received_events", ReposURL:"https://api.github.com/users/merotechnologies/repos", StarredURL:"https://api.github.com/users/merotechnologies/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/merotechnologies/subscriptions"}, Name:"web-dashboard", FullName:"merotechnologies/web-dashboard", Description:"📉 Removing the uncertainty from building maintenance", Homepage:"https://mero.co", DefaultBranch:"sprint/october19", CreatedAt:github.Timestamp{2018-09-02 22:36:03 +0000 UTC}, PushedAt:github.Timestamp{2020-10-07 12:18:54 +0000 UTC}, UpdatedAt:github.Timestamp{2020-10-06 20:06:04 +0000 UTC}, HTMLURL:"https://github.com/merotechnologies/web-dashboard", CloneURL:"https://github.com/merotechnologies/web-dashboard.git", GitURL:"git://github.com/merotechnologies/web-dashboard.git", SSHURL:"git@github.com:merotechnologies/web-dashboard.git", SVNURL:"https://github.com/merotechnologies/web-dashboard", Language:"PHP", Fork:false, ForksCount:0, OpenIssuesCount:88, StargazersCount:2, WatchersCount:2, Size:12329, Archived:false, Disabled:false, Private:true, HasIssues:true, HasWiki:true, HasPages:false, HasProjects:true, HasDownloads:true, URL:"https://api.github.com/repos/merotechnologies/web-dashboard", ArchiveURL:"https://api.github.com/repos/merotechnologies/web-dashboard/{archive_format}{/ref}", AssigneesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/assignees{/user}", BlobsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/blobs{/sha}", BranchesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/branches{/branch}", CollaboratorsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/collaborators{/collaborator}", CommentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/comments{/number}", CommitsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/commits{/sha}", CompareURL:"https://api.github.com/repos/merotechnologies/web-dashboard/compare/{base}...{head}", ContentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/contents/{+path}", ContributorsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/contributors", DeploymentsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/deployments", DownloadsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/downloads", EventsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/events", ForksURL:"https://api.github.com/repos/merotechnologies/web-dashboard/forks", GitCommitsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/commits{/sha}", GitRefsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/refs{/sha}", GitTagsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/tags{/sha}", HooksURL:"https://api.github.com/repos/merotechnologies/web-dashboard/hooks", IssueCommentURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/comments{/number}", IssueEventsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues/events{/number}", IssuesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/issues{/number}", KeysURL:"https://api.github.com/repos/merotechnologies/web-dashboard/keys{/key_id}", LabelsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/labels{/name}", LanguagesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/languages", MergesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/merges", MilestonesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/milestones{/number}", NotificationsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/notifications{?since,all,participating}", PullsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/pulls{/number}", ReleasesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/releases{/id}", StargazersURL:"https://api.github.com/repos/merotechnologies/web-dashboard/stargazers", StatusesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/statuses/{sha}", SubscribersURL:"https://api.github.com/repos/merotechnologies/web-dashboard/subscribers", SubscriptionURL:"https://api.github.com/repos/merotechnologies/web-dashboard/subscription", TagsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/tags", TreesURL:"https://api.github.com/repos/merotechnologies/web-dashboard/git/trees{/sha}", TeamsURL:"https://api.github.com/repos/merotechnologies/web-dashboard/teams"}, User:github.User{Login:"merotechnologies", ID:42920552, NodeID:"MDEyOk9yZ2FuaXphdGlvbjQyOTIwNTUy", AvatarURL:"https://avatars2.githubusercontent.com/u/42920552?v=4", HTMLURL:"https://github.com/merotechnologies", GravatarID:"", Type:"Organization", SiteAdmin:false, URL:"https://api.github.com/users/merotechnologies", EventsURL:"https://api.github.com/users/merotechnologies/events{/privacy}", FollowingURL:"https://api.github.com/users/merotechnologies/following{/other_user}", FollowersURL:"https://api.github.com/users/merotechnologies/followers", GistsURL:"https://api.github.com/users/merotechnologies/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/merotechnologies/orgs", ReceivedEventsURL:"https://api.github.com/users/merotechnologies/received_events", ReposURL:"https://api.github.com/users/merotechnologies/repos", StarredURL:"https://api.github.com/users/merotechnologies/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/merotechnologies/subscriptions"}}})
