package clubhouse

import (
	"fmt"
	"net/http"
	"time"
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

func (c *Client) ListStoriesForProject(id string) ([]Story, error) {
	path := fmt.Sprintf("/projects/%s/stories", id)

	var stories []Story
	if err := c.get(path, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}

type CreateStoryCommentParams struct {
	// Required
	Text string `json:"text"`

	// Optional
	AuthorID   *string    `json:"author_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ExternalID *string    `json:"external_id,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type CreateExternalTicketParams struct {
	// Optional
	ExternalID  *string `json:"external_id,omitempty"`
	ExternalURL *string `json:"external_url,omitempty"`
}

type CreateLabelParams struct {
	// Required
	Name string `json:"name"`

	// Optional
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	ExternalID  *string `json:"external_id,omitempty"`
}

type CreateStoryLinkParams struct {
	// Required
	ObjectID  int64  `json:"object_id"`
	SubjectID int64  `json:"subject_id"`
	Verb      string `json:"verb"`
}

type CreateTaskParams struct {
	// Required
	Complete    bool   `json:"complete"`
	Description string `json:"description"`

	// Optional
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ExternalID *string    `json:"external_id,omitempty"`
	OwnerIDs   []string   `json:"owner_ids,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type CreateStoryParams struct {
	// Required
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   int64  `json:"project_id"`
	Archived    bool   `json:"archived"`
	StoryType   string `json:"story_type"`

	// Optional
	Comments            []CreateStoryCommentParams   `json:"comments,omitempty"`
	CompletedAtOverride *time.Time                   `json:"completed_at_override,omitempty"`
	CreatedAt           *time.Time                   `json:"created_at,omitempty"`
	Deadline            *time.Time                   `json:"deadline,omitempty"`
	EpicID              *int64                       `json:"epic_id,omitempty"`
	Estimate            *int64                       `json:"estimate,omitempty"`
	ExternalID          *string                      `json:"external_id,omitempty"`
	ExternalTickets     []CreateExternalTicketParams `json:"external_tickets,omitempty"`
	FileIDs             []int64                      `json:"file_ids,omitempty"`
	FollowerIDs         []string                     `json:"follower_ids,omitempty"`
	IterationID         *int64                       `json:"iteration_id,omitempty"`
	Labels              []CreateLabelParams          `json:"labels,omitempty"`
	LinkedFileIDs       []int64                      `json:"linked_file_ids,omitempty"`
	OwnerIDs            []string                     `json:"owner_ids,omitempty"`
	RequestedByID       *string                      `json:"requested_by_id,omitempty"`
	StartedAtOverride   *time.Time                   `json:"started_at_override,omitempty"`
	StoryLinks          []CreateStoryLinkParams      `json:"story_links,omitempty"`
	Tasks               []CreateTaskParams           `json:"tasks,omitempty"`
	UpdatedAt           *time.Time                   `json:"updated_at,omitempty"`
	WorkflowStateID     *int64                       `json:"workflow_state_id,omitempty"`
}

type UpdateStoriesParam struct {
	Name                string     `json:"name,omitempty"`
	Description         string     `json:"description,omitempty"`
	Archived            bool       `json:"archived"`
	BeforeStoryID       *int64     `json:"before_id,omitempty"`
	EpicID              *int64     `json:"epic_id,omitempty"`
	Estimate            *int64     `json:"estimate,omitempty"`
	BranchIDs           []*int64   `json:"branch_ids,omitempty"`
	CommitIDs           []*int64   `json:"commit_ids,omitempty"`
	FileIDs             []*int64   `json:"file_ids,omitempty"`
	FollowerIDs         []*int64   `json:"follower_ids,omitempty"`
	IterationID         []*int64   `json:"iteration_id,omitempty"`
	StoryType           []*int64   `json:"story_type,omitempty"`
	StartedAtOverride   *time.Time `json:"started_at_override,omitempty"`
	CompletedAtOverride *time.Time `json:"completed_at_override,omitempty"`
}

type DeleteStoriesParam struct {
	StoryIds []int64 `json:"story_ids"`
}


func (c *Client) StoryCreate(param CreateStoryParams) (*Story, error) {
	stories, err := c.StoriesCreate([]CreateStoryParams{param})

	if len(stories) != 0 {
		return stories[0], err
	}
	return nil, err
}

func (c *Client) StoriesCreate(params []CreateStoryParams) ([]*Story, error) {
	body := map[string][]CreateStoryParams{"stories": params}
	path := "/stories/bulk"

	req, err := c.makeRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var stories []*Story
	if err := c.decode(resp, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}

func (c *Client) StoryArchive(storyId int64) (*Story, error) {
	path := fmt.Sprintf("/stories/%d", storyId)

	updateStory := UpdateStoriesParam {
		Archived: true,
	}

	req, err := c.makeRequest(http.MethodPut, path, updateStory)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var story *Story
	if err := c.decode(resp, &story); err != nil {
		return nil, err
	}

	return story, nil
}

func (c *Client) StoryDelete(storyId int64) error {
	return c.StoriesDelete([]int64{storyId})
}

func (c *Client) StoriesDelete(storyIds []int64) error {
	path := "/stories/bulk"

	payload := DeleteStoriesParam{
		StoryIds: storyIds,
	}

	req, err := c.makeRequest(http.MethodDelete, path, payload)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) StoryGet(storyId int64) (*Story, error) {
	path := fmt.Sprintf("/stories/%d", storyId)

	req, err := c.makeRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var story *Story
	if err := c.decode(resp, &story); err != nil {
		return nil, err
	}

	return story, nil
}
