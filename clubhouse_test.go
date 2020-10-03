package clubhouse_test

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/nhalstead/clubhouse"
)

var client *clubhouse.Client

func TestMain(m *testing.M) {
	client = clubhouse.New(os.Getenv("CLUBHOUSE_TOKEN"))
	flag.BoolVar(&client.Debug, "debug", false, "Debug tests")
	flag.Parse()

	os.Exit(m.Run())
}

func wait(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func AssertInt64Equal(t *testing.T, a int64, b int64, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check failed: %v != %v", a, b)
	}
	t.Fatal(message)
}

func AssertStringEqual(t *testing.T, a string, b string, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check failed: %v != %v", a, b)
	}
	t.Fatal(message)
}

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check failed: %v != %v", a, b)
	}
	t.Fatal(message)
}

func AssertTrue(t *testing.T, a interface{}, message string) {
	if a == true {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != true", a)
	}
	t.Fatal(message)
}

func AssertFalse(t *testing.T, a interface{}, message string) {
	if a == false {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != false", a)
	}
	t.Fatal(message)
}

func AssertNil(t *testing.T, a interface{}, message string) {
	if a == nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != nil", a)
	}
	t.Fatal(message)
}

func AssertNotNil(t *testing.T, a interface{}, message string) {
	if a != nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != nil", a)
	}
	t.Fatal(message)
}

func AssertNotZero(t *testing.T, a interface{}, message string) {
	if a != 0 {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != 0", a)
	}
	t.Fatal(message)
}

func AssertNoError(t *testing.T, e error, message string) {
	if e == nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != 0", e.Error())
	}
	t.Fatal(message, "\n\t>", e.Error())
}

func AssertError(t *testing.T, e error, message string) {
	if e != nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Check Failed: %v != 0", e.Error())
	}
	t.Fatal(message, "\n\t>", e.Error())
}