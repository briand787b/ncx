package testutil

import (
	"encoding/json"
	"testing"
)

// PrintJSONForTest prints the msg, with a trailing JSONified value of v
func PrintJSONForTest(t *testing.T, msg string, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		t.Log("failed to marshal JSON: ", err)
	}

	t.Logf("%s: %s", msg, bs)
}
