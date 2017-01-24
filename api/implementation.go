package api

import "fmt"

type arachneAPI struct{}

// New returns an object that implements the ArachneAPI interface
func New() ArachneAPI {
	return &arachneAPI{}
}

// NewGame requests the start of a new game
func (a *arachneAPI) NewGame() (GameInfo, error) {
	return GameInfo{}, fmt.Errorf("not implemented")
}
