package clubhouse

import (
	"context"
	"fmt"
	"net/http"
)

const (
	StoryTypeFeature string = "feature"
	StoryTypeBug     string = "bug"
	StoryTypeChore   string = "chore"
)

type Story struct {
	AppURL               string           `json:"app_url"`
	Archived             bool             `json:"archived"`
	Blocked              bool             `json:"blocked"`
	Blocker              bool             `json:"blocker"`
	Branches             []Branch         `json:"branches"`
	Comments             []Comment        `json:"comments"`
	Commits              []Commit         `json:"commits"`
	Completed            bool             `json:"completed"`
	CompletedAt          string           `json:"completed_at"`
	CompletedAtOverride  string           `json:"completed_at_override"`
	CreatedAt            string           `json:"created_at"`
	CycleTime            int64            `json:"cycle_time"`
	Deadline             string           `json:"deadline"`
	Description          string           `json:"description"`
	EntityType           string           `json:"entity_type"`
	EpicID               int64            `json:"epic_id"`
	Estimate             int64            `json:"estimate"`
	ExternalID           string           `json:"external_id"`
	ExternalTickets      []ExternalTicket `json:"external_tickets"`
	Files                []File           `json:"files"`
	FollowerIDs          []string         `json:"follower_ids"`
	GroupMentionIDs      []string         `json:"group_mention_ids"`
	ID                   int64            `json:"id"`
	IterationID          int64            `json:"iteration_id"`
	Labels               []Label          `json:"label"`
	LeadTime             int64            `json:"lead_time"`
	LinkedFiles          []LinkedFile     `json:"linked_files"`
	MemberMentionIDs     []string         `json:"member_mention_ids"`
	MentionIDs           []string         `json:"mention_ids"`
	MovedAt              string           `json:"moved_at"`
	Name                 string           `json:"name"`
	OwnerIDs             []string         `json:"owner_ids"`
	Position             int64            `json:"position"`
	PreviousIterationIDs []int64          `json:"previous_iteration_ids"`
	ProjectID            int64            `json:"project_id"`
	PullRequests         []PullRequest    `json:"pull_requests"`
	RequestedByID        string           `json:"requested_by_id"`
	Started              bool             `json:"started"`
	StartedAt            string           `json:"started_at"`
	StartedAtOverride    string           `json:"started_at_override"`
	Stats                struct {
		NumRelatedDocuments int64 `json:"num_related_documents"`
	} `json:"stats"`
	StoryLinks      []StoryLink `json:"story_links"`
	StoryType       string      `json:"story_type"`
	Tasks           []Task      `json:"tasks"`
	UpdatedAt       string      `json:"updated_at"`
	WorkflowStateID int64       `json:"workflow_state_id"`
}

func (c *Client) ListStoriesForProject(ctx context.Context, id string) ([]Story, error) {
	path := fmt.Sprintf("/projects/%s/stories", id)

	var stories []Story
	if err := c.get(path, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}

func (c *Client) StoriesCreateMultiple(ctx context.Context, stories []Story) error {
	body := map[string][]Story{"stories": stories}
	path := buildPath("/stories/bulk", "token", c.Token)

	req, err := c.makeRequest(http.MethodPost, path, body)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return fmt.Errorf("Client.StoriesCreateMultiple: %w", err)
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Client.StoriesCreateMultiple: code=%d", resp.StatusCode)
	}

	return nil
}
