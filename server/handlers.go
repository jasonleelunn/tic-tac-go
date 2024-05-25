package server

import (
	"net/http"
	"strconv"

	"github.com/jasonleelunn/tic-tac-go/game"
)

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	g := game.New()

	sessionToken := newSession(g)

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    sessionToken,
		Expires:  sessions[sessionToken].expiry,
		SameSite: http.SameSiteStrictMode,
	})

	game.Page(g.String()).Render(r.Context(), w)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	s, ok := r.Context().Value(sessionTokenKey).(session)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	g := s.gameState

	w.Write([]byte(g.String()))
}

func cellHandler(w http.ResponseWriter, r *http.Request) {
	s, ok := r.Context().Value(sessionTokenKey).(session)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	g := s.gameState

	if g.IsFinished() {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cellNumber, err := strconv.Atoi(r.PathValue("cell"))

	if err != nil || cellNumber < 0 || cellNumber > 8 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = g.MarkGridCell(cellNumber)

	// disallow choosing a cell more than once
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Add("HX-Trigger", "status-changed")
	w.Write([]byte(g.GetCurrentToken().String()))

	g.ChangeCurrentToken()
}
