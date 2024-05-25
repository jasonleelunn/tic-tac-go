package game

import (
	"errors"
)

type token int

type GameState struct {
	currentToken token
	winningToken token
	grid         map[int]token
}

const (
	empty token = iota
	cross
	naught
)

var (
	tokenStrings = map[token]string{
		empty:  "",
		cross:  "X",
		naught: "O",
	}

	winGroups = [8][3]int{
		// horizontal
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},

		// vertical
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},

		// diagonal
		{0, 4, 8},
		{2, 4, 6},
	}
)

func (t token) String() string {
	return tokenStrings[t]
}

func New() *GameState {
	return &GameState{
		currentToken: cross,
		grid:         make(map[int]token, 9),
	}
}

func (g *GameState) String() string {
	if g.IsFinished() {
		if g.winningToken == empty {
			return "Draw!"
		} else {
			return g.winningToken.String() + " Wins!"
		}
	}

	return g.currentToken.String() + "'s Turn"
}

func (g *GameState) IsFinished() bool {
	// grid is full, game must be over
	if len(g.grid) == 9 {
		return true
	}

	return g.winningToken != empty
}

func (g *GameState) GetWinningToken() token {
	return g.winningToken
}

func (g *GameState) GetCurrentToken() token {
	return g.currentToken
}

func (g *GameState) ChangeCurrentToken() {
	if g.currentToken == naught {
		g.currentToken = cross
	} else {
		g.currentToken = naught
	}
}

func (g *GameState) MarkGridCell(cellNumber int) error {
	c := g.grid[cellNumber]

	if c != empty {
		return errors.New("This grid cell is already taken!")
	}

	g.grid[cellNumber] = g.currentToken

	g.updateWinningToken()

	return nil
}

func (g *GameState) updateWinningToken() {
	for _, winGroup := range winGroups {
		var crossMatch int
		var naughtMatch int

		for _, cell := range winGroup {
			c := g.grid[cell]

			if c == cross {
				crossMatch++
			} else if c == naught {
				naughtMatch++
			}
		}

		if crossMatch == 3 {
			g.winningToken = cross
			break
		}

		if naughtMatch == 3 {
			g.winningToken = naught
			break
		}
	}
}
