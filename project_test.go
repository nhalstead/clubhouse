package clubhouse_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/haleyrc/clubhouse"
)

var client *clubhouse.Client

func TestMain(m *testing.M) {
	client = clubhouse.New(os.Getenv("CLUBHOUSE_TOKEN"))
	flag.BoolVar(&client.Debug, "debug", false, "Debug tests")
	flag.Parse()
	os.Exit(m.Run())
}

func TestListProjects(t *testing.T) {
	ctx := context.Background()

	projects, err := client.ListProjects(ctx)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(projects, "", "    ")
	fmt.Println(string(b))
}
