package clubhouse

type Commit struct {
	AuthorEmail     string   `json:"author_email"`
	AuthorID        string   `json:"author_id"`
	AuthorIdentity  Identity `json:"author_identity"`
	CreatedAt       string   `json:"created_at"`
	EntityType      string   `json:"entity_type"`
	Hash            string   `json:"hash"`
	ID              int64    `json:"id"`
	MergedBranchIDs []int64  `json:"merged_branch_ids"`
	Message         string   `json:"message"`
	RepositoryID    int64    `json:"repository_id"`
	Timestamp       string   `json:"timestamp"`
	UpdatedAt       string   `json:"updated_at"`
	URL             string   `json:"url"`
}
