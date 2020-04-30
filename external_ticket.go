package clubhouse

type ExternalTicket struct {
	EpicIDs     []int64 `json:"epic_ids"`
	ExternalID  string  `json:"external_id"`
	ExternalURL string  `json:"external_url"`
	ID          string  `json:"id"`
	StoryIDs    []int64 `json:"story_ids"`
}
