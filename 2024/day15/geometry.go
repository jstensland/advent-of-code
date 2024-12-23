package day15

type State int

func (s State) String() string {
	return map[State]string{
		Robot:    "@",
		Wall:     "#",
		Box:      "O",
		BoxLeft:  "[",
		BoxRight: "]",
		Empty:    ".",
	}[s]
}

const (
	Robot State = iota
	Wall
	Box
	BoxLeft  // for wide boxes
	BoxRight // for wide boxes
	Empty
)

type Grid struct {
	data     [][]State
	robotLoc Location
	movesVec []MoveVector
	Width    int
	Height   int
}

func (g *Grid) GetLoc(l Location) State {
	return g.data[l.Row][l.Col]
}

func (g *Grid) String() string {
	out := ""
	for row := range g.Height {
		for col := range g.Width {
			out += g.data[row][col].String()
		}
		out += "\n"
	}
	return out
}

// Copy returns a copy of the grid.
func (g *Grid) Copy() *Grid {
	out := Grid{
		data:     [][]State{},
		robotLoc: g.robotLoc,
		movesVec: []MoveVector{},
		Width:    g.Width,
		Height:   g.Height,
	}
	out.movesVec = append(out.movesVec, g.movesVec...)
	for row := range g.Height {
		newRow := make([]State, g.Width)
		for col := range g.Width {
			newRow[col] = g.data[row][col]
		}
		out.data = append(out.data, newRow)
	}
	return &out
}

type Location struct {
	Row int
	Col int
}

// Add will add a vector to a location and returns the resulting location.
func (loc Location) Add(vec MoveVector) Location {
	return Location{Row: loc.Row + vec.deltaRow, Col: loc.Col + vec.deltaCol}
}

// boxLoc is the two sides of a box
type boxLocation struct {
	left  Location
	right Location
}

type MoveVector struct {
	deltaRow int
	deltaCol int
}

//nolint:gochecknoglobals // reference sentinel values for possilbe moves
var (
	upVec    = MoveVector{-1, 0}
	rightVec = MoveVector{0, 1}
	downVec  = MoveVector{1, 0}
	leftVec  = MoveVector{0, -1}
)

func (mv MoveVector) String() string {
	return map[MoveVector]string{
		upVec:    "^",
		rightVec: ">",
		downVec:  "v",
		leftVec:  "<",
	}[mv]
}
