package api

import "testing"

func TestAPI(t *testing.T) {
	n := new()
	_, err := n.NewGame()
	if err != nil {
		t.Fatalf("NewGame() failed: %s", err)
	}
}
