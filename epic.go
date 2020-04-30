package clubhouse

import "context"

type Epic struct {
	AppURL              string           `json:"app_url"`
	Archived            bool             `json:"archived"`
	Comments            []Comment        `json:"comments"`
	Completed           bool             `json:"completed"`
	CompletedAt         string           `json:"completed_at"`
	CompletedAtOverride string           `json:"completed_at_override"`
	CreatedAt           string           `json:"created_at"`
	Deadline            string           `json:"deadline"`
	Description         string           `json:"description"`
	EntityType          string           `json:"entity_type"`
	EpicStateID         int64            `json:"epic_state_id"`
	ExternalID          string           `json:"external_id"`
	ExternalTickets     []ExternalTicket `json:"external_tickets"`
	FollowerIDs         []string         `json:"follower_ids"`
	GroupMentionIDs     []string         `json:"group_mention_ids"`
	ID                  int64            `json:"id"`
	Labels              []Label          `json:"labels"`
	MemberMentionIDs    []string         `json:"member_mention_ids"`
	MentionIDs          []string         `json:"mention_ids"`
	MilestoneID         int64            `json:"milestone_id"`
	Name                string           `json:"name"`
	OwnerIDs            []string         `json:"owner_ids"`
	PlannedStartDate    string           `json:"planned_start_date"`
	Position            int64            `json:"position"`
	ProjectIDs          []int64          `json:"project_ids"`
	RequestByID         string           `json:"requested_by_id"`
	Started             bool             `json:"started"`
	StartedAt           string           `json:"started_at"`
	StartedAtOverride   string           `json:"started_at_override"`
	State               string           `json:"state"`
	Stats               struct {
		AverageCycleTime      int64  `json:"average_cycle_time"`
		AverageLeadTime       int64  `json:"average_lead_time"`
		LastStoryUpdate       string `json:"last_story_update"`
		NumPoints             int64  `json:"num_points"`
		NumPointsDone         int64  `json:"num_points_done"`
		NumPointsStarted      int64  `json:"num_points_started"`
		NumPointsUnstarted    int64  `json:"num_points_unstarted"`
		NumRelatedDocuments   int64  `json:"num_related_documents"`
		NumStoriesDone        int64  `json:"num_stories_done"`
		NumStoriesStarted     int64  `json:"num_stories_started"`
		NumStoriesUnestimated int64  `json:"num_stories_unestimated"`
		NumStoriesUnstarted   int64  `json:"num_stories_unstarted"`
	} `json:"stats"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) ListEpics(ctx context.Context) ([]Epic, error) {
	path := "/epics"

	var epics []Epic
	if err := c.get(path, &epics); err != nil {
		return nil, err
	}

	return epics, nil
}
