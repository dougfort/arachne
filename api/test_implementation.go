package api

import "fmt"

type testAPIStruct struct {
}

func new() ArachneAPI {
	return &testAPIStruct{}
}

// NewGame requests the start of a new game
func (a *testAPIStruct) NewGame() (GameInfo, error) {
	return GameInfo{}, fmt.Errorf("not implemented")
}
