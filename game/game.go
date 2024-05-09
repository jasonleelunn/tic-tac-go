package game

import "errors"

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

func (s *State) MarkGridCell(cellNumber string) error {
	c := s.grid[cellNumber]

	if c == true {
		return errors.New("This grid cell is already taken!")
	}

	s.grid[cellNumber] = true

	return nil
}
