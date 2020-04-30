package clubhouse

type Task struct {
	Complete         bool     `json:"complete"`
	CompletedAt      string   `json:"completed_at"`
	CreatedAt        string   `json:"created_at"`
	Description      string   `json:"description"`
	EntityType       string   `json:"entity_type"`
	ExternalID       string   `json:"external_id"`
	GroupMentionIDs  []string `json:"group_mention_ids"`
	ID               int64    `json:"id"`
	MemberMentionIDs []string `json:"member_mention_ids"`
	MentionIDs       []string `json:"mention_ids"`
	OwnerIDs         []string `json:"owner_ids"`
	Position         int64    `json:"position"`
	StoryID          int64    `json:"story_id"`
	UpdatedAt        string   `json:"updated_at"`
}
