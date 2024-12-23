package day15

// Moves exposes parsed grid moves for tests
func (g *Grid) Moves() []MoveVector {
	return g.movesVec
}

func (g *Grid) RobotLocationV2() Location {
	return g.robotLoc
}
