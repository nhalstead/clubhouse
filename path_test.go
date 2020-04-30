package clubhouse

import "testing"

func TestBuildPath(t *testing.T) {
	want := "/stories/bulk?token=mytoken"
	got := buildPath("/stories/bulk", "token", "mytoken")
	if got != want {
		t.Errorf("buildPath(...) = %q, wanted %q", got, want)
	}
}
