package metrics

import (
	"time"

	"github.com/google/go-github/github"
)

// ActivityMetrics measures contributor and maintainer activity throughout the
// project. Most of these metrics were adopted from github.com/chaoss/metrics
type ActivityMetrics struct {
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

// GetActivityMetrics returns a set of activity metrics for a particular repository
func GetActivityMetrics(r *github.Repository) ActivityMetrics {
	return ActivityMetrics{
		communityAge: getCommunityAge(r),
		updateAge:    getUpdateAge(r),
	}
}

func getCommunityAge(r *github.Repository) time.Duration {
	return time.Since(r.GetCreatedAt().Time)
}

func getUpdateAge(r *github.Repository) time.Duration {
	return time.Since(r.GetUpdatedAt().Time)
}
