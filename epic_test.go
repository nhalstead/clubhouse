package clubhouse_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestListEpics(t *testing.T) {
	ctx := context.Background()

	epics, err := client.ListEpics(ctx)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(epics, "", "    ")
	fmt.Println(string(b))
}
