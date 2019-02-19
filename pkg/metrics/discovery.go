package metrics

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

// DiscoveryMetrics measures how well the project is being noticed.  Although
// stars and forks do not correlate with downloads or usage, it is a good
// indicator if people are finding your open-source project.
type DiscoveryMetrics struct {
	stars    int
	forks    int
	watchers int
}

var q struct {
	Repository struct {
		ForkCount  int
		Stargazers struct {
			TotalCount int
		}
		Watchers struct {
			TotalCount int
		}
	} `graphql:"repository(owner: $owner, name: $name)"`
}

// GetDiscoveryMetrics returns all metrics for a particular repository.
func GetDiscoveryMetrics(ctx context.Context, client *githubv4.Client, request map[string]interface{}) DiscoveryMetrics {

	err := client.Query(ctx, &q, request)
	if err != nil {
		log.Panic(err)
	}
	return DiscoveryMetrics{
		stars:    q.Repository.Stargazers.TotalCount,
		forks:    q.Repository.ForkCount,
		watchers: q.Repository.Watchers.TotalCount,
	}
}
