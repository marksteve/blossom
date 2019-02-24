package metrics

import "time"

var q struct {
	Repository struct {
		ForkCount  int
		Stargazers struct {
			TotalCount int
		}
		Watchers struct {
			TotalCount int
		}
		CreatedAt time.Time
		UpdatedAt time.Time
	} `graphql:"repository(owner: $owner, name: $name)"`
	RateLimit struct {
		Limit     int
		Cost      int
		Remaining int
	}
}
