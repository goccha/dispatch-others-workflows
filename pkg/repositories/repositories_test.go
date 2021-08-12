package repositories

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	payload := map[string]interface{}{
		"goccha/test-repository": map[string]interface{}{
			"tag": "v0.0.1",
		},
		"goccha/test-repository2": map[string]interface{}{
			"tag": "v0.0.2",
		},
	}
	if req, err := Parse(payload); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%v\n", req)
	}
}
