package clubhouse_test

import (
	"testing"
)

func TestListProjects(t *testing.T) {

	projects, err := client.ListProjects()
	if err != nil {
		t.Fatal(err)
	}

	AssertNotZero(t, len(projects), "No Projects")
}

