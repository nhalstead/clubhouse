package clubhouse

import "time"

type Workflow struct {
	// Like EpicWorkflow
	ID              int64           `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	EntityType      string          `json:"entity_type"`
	DefaultStateId  int64           `json:"default_state_id"`
	States          []WorkflowState `json:"states"`
	AutoAssignOwner bool            `json:"auto_assign_owner"`
	ProjectIds      []int64         `json:"project_ids"`
	TeamId          int64           `json:"team_id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type WorkflowState struct {
	EpicState
	NumStories        int64  `json:"num_stories"`
	NumStoryTemplates int64  `json:"num_story_templates"`
	Verb              string `json:"verb"`
}

func (c *Client) ListWorkflows() ([]Workflow, error) {
	path := "/workflows"

	var workflows []Workflow
	if err := c.get(path, &workflows); err != nil {
		return nil, err
	}

	return workflows, nil
}
