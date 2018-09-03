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
func (h HighestMoveStruct) Order(eSlice []EvaluatedMoveType) ([]RankedMoveType, error) {
	var rSlice []RankedMoveType

	for _, e := range eSlice {
		rSlice = append(
			rSlice,
			RankedMoveType{EvaluatedMoveType: e, Rank: e.FromCount + e.ToCount},
		)
	}

	sort.Slice(rSlice, func(i, j int) bool { return rSlice[i].Rank < rSlice[j].Rank })

	return rSlice, nil
}
