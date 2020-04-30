package clubhouse

type PullRequest struct {
	BranchID         int64  `json:"branch_id"`
	BranchName       string `json:"branch_name"`
	Closed           bool   `json:"closed"`
	CreatedAt        string `json:"created_at"`
	EntityType       string `json:"entity_type"`
	ID               int64  `json:"id"`
	NumAdded         int64  `json:"num_added"`
	NumCommits       int64  `json:"num_commits"`
	NumModified      int64  `json:"num_modified"`
	NumRemoved       int64  `json:"num_removed"`
	Number           int64  `json:"number"`
	TargetBranchID   int64  `json:"target_branch_id"`
	TargetBranchName string `json:"target_branch_name"`
	Title            string `json:"title"`
	UpdatedAt        string `json:"updated_at"`
	URL              string `json:"url"`
}
