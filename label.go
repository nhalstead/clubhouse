package clubhouse

type Label struct {
	AppURL      string `json:"app_url"`
	Archived    bool   `json:"archived"`
	Color       string `json:"color"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	EntityType  string `json:"entity_type"`
	ExternalID  string `json:"external_id"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Stats       struct {
		NumEpics              int64 `json:"num_epics"`
		NumPointsCompleted    int64 `json:"num_points_completed"`
		NumPointsInProgress   int64 `json:"num_points_in_progress"`
		NumPointsTotal        int64 `json:"num_points_total"`
		NumRelatedDocuments   int64 `json:"num_related_documents"`
		NumStoriesCompleted   int64 `json:"num_stories_completed"`
		NumStoriesInProgress  int64 `json:"num_stories_in_progress"`
		NumStoriesTotal       int64 `json:"num_stories_total"`
		NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	} `json:"stats"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) AddLabelToMultipleStories(ids []int64, params CreateLabelParams) error {
	path := "/stories/bulk"
	body := map[string]interface{}{
		"story_ids":  ids,
		"labels_add": []CreateLabelParams{params},
	}
	if err := c.put(path, body, nil); err != nil {
		return err
	}
	return nil
}
