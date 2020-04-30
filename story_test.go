package clubhouse_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestListStories(t *testing.T) {
	ctx := context.Background()

	// projects, err := client.ListProjects(ctx)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	stories, err := client.ListStoriesForProject(ctx, "2")
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(stories, "", "    ")
	fmt.Println(string(b))
}
