package metrics

import (
	"github.com/google/go-github/github"
)

// DiscoveryMetrics measures how well the project is being noticed.  Although
// stars and forks do not correlate with downloads or usage, it is a good
// indicator if people are finding your open-source project.
type DiscoveryMetrics struct {
	stars    int
	forks    int
	watchers int
}

// GetDiscoveryMetrics returns all metrics for a particular repository.
func GetDiscoveryMetrics(r *github.Repository) DiscoveryMetrics {
	return DiscoveryMetrics{
		stars:    getStargazers(r),
		forks:    getForks(r),
		watchers: getWatchers(r),
	}
}

func getStargazers(r *github.Repository) int {
	return r.GetStargazersCount()
}

func getForks(r *github.Repository) int {
	return r.GetForksCount()
}

func getWatchers(r *github.Repository) int {
	return r.GetWatchersCount()
}
