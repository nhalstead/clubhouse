package clubhouse

type Branch struct {
	CreatedAt       string        `json:"created_at"`
	Deleted         bool          `json:"deleted"`
	EntityType      string        `json:"entity_type"`
	ID              int64         `json:"id"`
	MergedBranchIDs []int64       `json:"merged_branch_ids"`
	Name            string        `json:"name"`
	Persistent      bool          `json:"persistent"`
	PullRequests    []PullRequest `json:"pull_requests"`
	RepositoryID    int64         `json:"repository_id"`
	UpdatedAt       string        `json:"updated_at"`
	URL             string        `json:"url"`
}
