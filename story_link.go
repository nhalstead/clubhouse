package clubhouse

type StoryLink struct {
	CreatedAt  string `json:"created_at"`
	EntityType string `json:"entity_type"`
	ID         int64  `json:"id"`
	ObjectID   int64  `json:"object_id"`
	SubjectID  int64  `json:"subject_id"`
	Type       string `json:"type"`
	UpdatedAt  string `json:"updated_at"`
	Verb       string `json:"verb"`
}
