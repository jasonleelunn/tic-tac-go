package game

import (
	"testing"
)

func TestUpdateWinningToken(t *testing.T) {
	g := New()

	g.updateWinningToken()

	if g.winningToken != empty {
		t.Fatal("winningToken should be empty")
	}

	g.grid = map[int]token{
		0: cross,
		1: cross,
		2: cross,
		3: empty,
		4: empty,
		5: empty,
		6: empty,
		7: empty,
		8: empty,
	}

	g.updateWinningToken()

	if g.winningToken != cross {
		t.Fatalf("winningToken is %s, want %s", g.winningToken, cross)
	}

	g.grid = map[int]token{
		0: empty,
		1: cross,
		2: naught,
		3: cross,
		4: naught,
		5: empty,
		6: naught,
		7: empty,
		8: empty,
	}

	g.updateWinningToken()

	if g.winningToken != naught {
		t.Fatalf("winningToken is %s, want %s", g.winningToken, naught)
	}
}
