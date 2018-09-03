package game

import "sort"

// HighestMoveStruct implements the Orderer interface
type HighestMoveStruct struct{}

// NewHighestMove returns an object that implements the Orderer interface
func NewHighestMove() HighestMoveStruct {
	return HighestMoveStruct{}
}

// Order places the moves in order by highest count of sequential cards of the
// same suit after move
func (h HighestMoveStruct) Order(s []EvaluatedMoveType) error {
	sort.Slice(
		s,
		func(i, j int) bool {
			return (s[i].FromCount + s[i].ToCount) < (s[j].FromCount + s[j].ToCount)
		},
	)

	return nil
}
