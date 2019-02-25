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
		CreatedAt      time.Time
		UpdatedAt      time.Time
		CommitComments struct {
			TotalCount int
		}
		ClosedIssues struct {
			TotalCount int
		} `graphql:"closedIssues: issues(states:CLOSED)"`
		OpenIssues struct {
			TotalCount int
		} `graphql:"openIssues: issues(states:OPEN)"`
		ClosedPRs struct {
			TotalCount int
		} `graphql:"closedPRs: pullRequests(states:CLOSED)"`
		MergedPRs struct {
			TotalCount int
		} `graphql:"mergedPRs: pullRequests(states:MERGED)"`
		OpenPRs struct {
			TotalCount int
		} `graphql:"openPRs: pullRequests(states:OPEN)"`
		PullRequests struct {
			TotalCount int
		}
		CommitCount struct {
			Target struct {
				Commit struct {
					History struct {
						TotalCount int
					}
				} `graphql:"... on Commit"`
			}
		} `graphql:"commitCount: defaultBranchRef"`
	} `graphql:"repository(owner: $owner, name: $name)"`
	RateLimit struct {
		Limit     int
		Cost      int
		Remaining int
	}
}
