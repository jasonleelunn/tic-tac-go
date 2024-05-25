package server

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jasonleelunn/tic-tac-go/game"
)

const (
	sessionLength     = 2 * time.Hour
	sessionCookieName = "session_token"

	sessionTokenKey requestContextKey = 0
)

var sessions = map[string]session{}

type requestContextKey int

type session struct {
	expiry    time.Time
	gameState *game.GameState
}

func (s *session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func newSession(state *game.GameState) string {
	token := uuid.NewString()

	expiry := time.Now().Add(sessionLength)

	sessions[token] = session{
		expiry,
		state,
	}

	return token
}

func ensureSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(sessionCookieName)

		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sessionToken := c.Value
		session, exists := sessions[sessionToken]

		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctxWithSession := context.WithValue(r.Context(), sessionTokenKey, session)
		rWithSession := r.WithContext(ctxWithSession)
		next.ServeHTTP(w, rWithSession)
	}
}
