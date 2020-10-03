package clubhouse_test

import (
	"github.com/nhalstead/clubhouse"
	"testing"
)

func TestListStories(t *testing.T) {

	stories, err := client.ListStoriesForProject("2")
	if err != nil {
		t.Fatal(err)
	}

	AssertNotZero(t, len(stories), "No Stories")
}

func TestCreateGetDeleteStory(t *testing.T) {
	// Select a project to run the tests in.
	projects, _ := client.ListProjects()
	AssertNotNil(t, projects, "Projects Exist")
	project := projects[0]

	// Create a new Story, Ensure that the name is Test
	name := "Testing 123!"
	newStory, err := client.StoryCreate(clubhouse.CreateStoryParams{
		Name: name,
		Description: "something to describe this issue",
		Archived: false,
		StoryType: clubhouse.StoryTypeFeature,
		ProjectID: project.ID,
	})
	AssertNoError(t, err, "Story Created without error")
	AssertNotNil(t, newStory, "Story response exists")
	AssertStringEqual(t, newStory.Name, name, "Ensure the name was set.")

	// Let the API breath a bit
	wait(1)

	// Get the Story, Ensure that we have a result
	story, err := client.StoryGet(newStory.ID)
	AssertNoError(t, err, "Get the new story without error")
	AssertStringEqual(t, newStory.Name, story.Name, "Ensure the name was set.")

	// Let the API breath a bit
	wait(1)

	// Delete the story and ensure we don't have an error
	archiveStory, err := client.StoryArchive(newStory.ID)
	AssertNoError(t, err, "Archive the story without error")
	AssertInt64Equal(t, archiveStory.ID, newStory.ID, "Ensure the ID are the same.")
	AssertStringEqual(t, newStory.Name, story.Name, "Ensure the name was set.")

	// Delete the story and ensure we don't have an error
	err = client.StoryDelete(newStory.ID)
	AssertNoError(t, err, "Delete the story without error")

	// Test that the story is deleted
	_, err = client.StoryGet(newStory.ID)
	AssertError(t, err, "Fail get the new story since it's deleted")
}
