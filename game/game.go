package game

type token int

type State struct {
	currentToken token
	grid         map[string]bool
	finished     bool
}

const (
	naught token = iota
	cross
)

var (
	tokens = map[token]string{
		naught: "O",
		cross:  "X",
	}
)

func New() *State {
	return &State{
		currentToken: cross,
		grid:         make(map[string]bool),
		finished:     false,
	}
}

func (s *State) GetCurrentTokenString() string {
	return tokens[s.currentToken]
}

func (s *State) ChangeCurrentToken() {
	if s.currentToken == naught {
		s.currentToken = cross
	} else {
		s.currentToken = naught
	}
}

func (s *State) GetGridCell(c string) bool {
	return s.grid[c]
}

func (s *State) MarkGridCell(c string) {
	s.grid[c] = true
}
