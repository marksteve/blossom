package metrics

import (
	"context"
	"log"
	"time"

	"github.com/shurcooL/githubv4"
)

// Metrics contain all open-source related metrics.
type Metrics struct {
	stars                     int           // No. of stars in repository
	forks                     int           // No. of forks in repository
	watchers                  int           // No. of watchers in repository
	avgIssueResponseTime      time.Duration // Avg. duration for first response in an issue
	avgOpenIssueTime          time.Duration // Avg. duration for issues to remain open
	avgIssueResolutionTime    time.Duration // Avg. time it takes for issues to be closed
	avgPRComments             float64       // No. of comments per pull request
	busFactor                 int           // No. of developers to destroy progress
	closedIssues              int           // No. of closed issues
	codeCommits               int           // No. of commits
	codeMergeDuration         time.Duration // Duration between PR and merge
	codeReviewEfficiency      float64       // Number of merged PR / closed PRs
	communityAge              time.Duration // Duration since repo was first released
	contributionAcceptance    float64       // Ratio of merged PR / total PRs
	contributors              int           // No. of contributors
	issueComments             float64       // No. of comments per issue
	issueResolutionEfficiency float64       // No. of closed issues / abandoned issues
	newContributions          float64       // Percent of contributions from first-timers / accepted contributions over time
	ponyFactor                float64       // Min. number of developers performing 50% of the commits
	prCommentDuration         time.Duration // Duration between PR creation and recent comment
	prMadeClosed              float64       // No. of closed PRs / total PRs
	prOpen                    int           // No. of open PRs
	releaseVelocity           time.Duration // Avg. time between releases
	sizeCodeBase              int           // Lines of code
	updateAge                 time.Duration // Time since last update
}

// Get obtains all Metrics.
func Get(ctx context.Context, client *githubv4.Client, request map[string]interface{}) Metrics {

	err := client.Query(ctx, &q, request)
	if err != nil {
		log.Panic(err)
	}

	return Metrics{
		stars:    q.Repository.Stargazers.TotalCount,
		forks:    q.Repository.ForkCount,
		watchers: q.Repository.Watchers.TotalCount,
	}
}
