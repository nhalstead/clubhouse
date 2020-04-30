package clubhouse

import (
	"context"
)

type Project struct {
	Abbreviation      string   `json:"abbreviation"`
	AppURL            string   `json:"app_url"`
	Archived          bool     `json:"archived"`
	Color             string   `json:"color"`
	CreatedAt         string   `json:"created_at"`
	DaysToThermometer int64    `json:"days_to_thermometer"`
	Description       string   `json:"description"`
	EntityType        string   `json:"entity_type"`
	ExternalID        string   `json:"external_id"`
	FollowerIDs       []string `json:"follower_ids"`
	ID                int64    `json:"id"`
	IterationLength   int64    `json:"iteration_length"`
	Name              string   `json:"name"`
	ShowThermometer   bool     `json:"show_thermometer"`
	StartTime         string   `json:"start_time"`
	Stats             struct {
		NumPoints           int64 `json:"num_points"`
		NumRelatedDocuments int64 `json:"num_related_documents"`
		NumStories          int64 `json:"num_stories"`
	} `json:"stats"`
	TeamID    int64  `json:"team_id"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) ListProjects(ctx context.Context) ([]Project, error) {
	path := "/projects"

	var projects []Project
	if err := c.get(path, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
