package clubhouse_test

import (
	"testing"
)

func TestListWorkflows(t *testing.T) {

	stories, err := client.ListWorkflows()
	if err != nil {
		t.Fatal(err)
	}

	AssertNotZero(t, len(stories), "No Workflows")
}
