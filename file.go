package clubhouse

type File struct {
	ContentType      string   `json:"content_type"`
	CreatedAt        string   `json:"created_at"`
	Description      string   `json:"description"`
	EntityType       string   `json:"entity_type"`
	ExternalID       string   `json:"external_id"`
	Filename         string   `json:"filename"`
	GroupMentionIDs  []string `json:"group_mention_ids"`
	ID               int64    `json:"id"`
	MemberMentionIDs []string `json:"member_mention_ids"`
	MentionIDs       []string `json:"mention_ids"`
	Name             string   `json:"name"`
	Size             int64    `json:"size"`
	StoryIDs         []int64  `json:"story_ids"`
	ThumbnailURL     string   `json:"thumbnail_url"`
	UpdatedAt        string   `json:"updated_at"`
	UploaderID       string   `json:"uploader_id"`
	URL              string   `json:"url"`
}
