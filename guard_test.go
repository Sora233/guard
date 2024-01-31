package guard

import (
	"encoding/json"
	"testing"
)

func TestGuard(t *testing.T) {
	var d1, d2 Guard[struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}]
	var ok = `{"id":123, "name":"abc"}`
	var malformed = `["bad", "format"]`

	if err := json.Unmarshal([]byte(ok), &d1); err != nil {
		t.Errorf("json unmarshal fail")
	} else if !d1.IsSuccess() || d1.Get().ID != 123 || d1.Get().Name != "abc" {
		t.Errorf("json data mismatched")
	}
	if err := json.Unmarshal([]byte(malformed), &d2); err != nil {
		t.Errorf("json unmarshal fail %v", err.Error())
	} else if d2.IsSuccess() {
		t.Errorf("json data mismatched")
	}
}
