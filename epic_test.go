package clubhouse_test

import (
	"testing"
)

func TestListEpics(t *testing.T) {

	epics, err := client.ListEpics()
	if err != nil {
		t.Fatal(err)
	}

	AssertNotZero(t, len(epics), "No Epics")
}
