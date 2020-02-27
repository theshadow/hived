package game

import (
	"math"

	"github.com/theshadow/hived"
)

// Distance will return the Manhattan distance between two coordinates.
// It does not validate if either coordinate is a cell with a valid piece
// as it's mostly here for path algorithms
func distance(a, b hived.Coordinate) int {
	return int(math.Round((math.Abs(float64(a.X()-b.X())) + math.Abs(float64(a.X()-b.X())) +
		math.Abs(float64(a.X()-b.X()))) / 2))
}

func neighbors(c hived.Coordinate) []hived.Coordinate {
	var neighbors []hived.Coordinate
	for _, loc := range hived.NeighborsMatrix {
		neighbors = append(neighbors, c.Add(loc))
	}
	return neighbors
}

func heuristic(a, b hived.Coordinate) int {
	return int(math.Round(math.Abs(float64(a.X()-b.X()))+
		math.Abs(float64(a.Y()-b.Y()))))
}

const (
	BeetleMaxDistance  = 1
	LadybugMaxDistance = 3
	PillBugMaxDistance = 1
	QueenMaxDistance   = 1
	SpiderMaxDistance  = 3
)