package metrics

import (
	"context"
	"log"
	"time"

	"github.com/shurcooL/githubv4"
)

// Metrics contain all open-source related metrics.
type Metrics struct {
	Stars                     int           // No. of stars in repository
	Forks                     int           // No. of forks in repository
	Watchers                  int           // No. of watchers in repository
	AvgIssueResponseTime      time.Duration // Avg. duration for first response in an issue
	AvgOpenIssueTime          time.Duration // Avg. duration for issues to remain open
	AvgIssueResolutionTime    time.Duration // Avg. time it takes for issues to be closed
	AvgPRComments             float64       // No. of comments per pull request
	BusFactor                 int           // No. of developers to destroy progress
	ClosedIssues              int           // No. of closed issues
	CodeCommits               int           // No. of commits
	CodeMergeDuration         time.Duration // Duration between PR and merge
	CodeReviewEfficiency      float64       // Number of merged PR / closed PRs
	CommunityAge              time.Duration // Duration since repo was first released
	ContributionAcceptance    float64       // Ratio of merged PR / total PRs
	Contributors              int           // No. of contributors
	IssueComments             float64       // No. of comments per issue
	IssueResolutionEfficiency float64       // No. of closed issues / abandoned issues
	NewContributions          float64       // Percent of contributions from first-timers / accepted contributions over time
	PonyFactor                float64       // Min. number of developers performing 50% of the commits
	PRCommentDuration         time.Duration // Duration between PR creation and recent comment
	PRMadeClosed              float64       // No. of closed PRs / total PRs
	PROpen                    int           // No. of open PRs
	ReleaseVelocity           time.Duration // Avg. time between releases
	SizeCodeBase              int           // Lines of code
	UpdateAge                 time.Duration // Time since last update
}

// ResourceStats indicates the remaining rate limit and the query cost
type ResourceStats struct {
	RateLimit int // Maximum rate limit
	Cost      int // Query cost
	Remaining int // Remaining API calls at the current time

}

// Get returns a Metrics structure.
func Get(ctx context.Context, client *githubv4.Client, request map[string]interface{}) (Metrics, ResourceStats) {

	err := client.Query(ctx, &q, request)
	if err != nil {
		log.Panic(err)
	}

	metrics := Metrics{
		Stars:                  q.Repository.Stargazers.TotalCount,
		Forks:                  q.Repository.ForkCount,
		Watchers:               q.Repository.Watchers.TotalCount,
		ClosedIssues:           q.Repository.ClosedIssues.TotalCount,
		CodeCommits:            q.Repository.CommitCount.Target.Commit.History.TotalCount,
		CodeReviewEfficiency:   getCodeReviewEfficiency(),
		PRMadeClosed:           getPRMadeClosed(),
		CommunityAge:           time.Since(q.Repository.CreatedAt),
		ContributionAcceptance: getContributionAcceptance(),
		PROpen:                 q.Repository.OpenPRs.TotalCount,
		UpdateAge:              time.Since(q.Repository.UpdatedAt),
	}

	resourceStats := ResourceStats{
		RateLimit: q.RateLimit.Limit,
		Cost:      q.RateLimit.Cost,
		Remaining: q.RateLimit.Remaining,
	}

	return metrics, resourceStats
}

// getPRMadeClosed is a helper function to obtain the prMadeClosed metric
func getPRMadeClosed() float64 {
	return float64(q.Repository.MergedPRs.TotalCount+q.Repository.ClosedPRs.TotalCount) / float64(q.Repository.PullRequests.TotalCount)
}

// getCodeReviewEfficiency is a helper function to obtain the codeReviewEfficiency metric
func getCodeReviewEfficiency() float64 {
	return float64(q.Repository.MergedPRs.TotalCount) / float64(q.Repository.ClosedPRs.TotalCount)
}

// getContributionAcceptance is a helper function to obtain the contributionAcceptance metric
func getContributionAcceptance() float64 {
	return float64(q.Repository.MergedPRs.TotalCount) / float64(q.Repository.PullRequests.TotalCount)
}
