package auth

import (
	"github.com/dchest/uniuri"
	"time"
)

type TokenStore struct {
	tokens *map[string]time.Time
}

func NewTokenStore() *TokenStore {
	tokens := make(map[string]time.Time, 0)

	return &TokenStore{
		tokens: &tokens,
	}
}

func (ts *TokenStore) NewToken() string {
	secret := uniuri.NewLen(64)

	tokens := *ts.tokens
	if _, exists := tokens[secret]; exists { // TODO: We have a race condition here
		return ts.NewToken()
	}

	expires := time.Now().Add(24 * 2 * time.Hour)

	tokens[secret] = expires
	return secret
}

func (ts *TokenStore) CheckToken(authToken string) bool {
	tokens := *ts.tokens
	if expires, exists := tokens[authToken]; !exists {
		return false
	} else {
		if expires.Unix() < time.Now().Unix() {
			// Clean up expired tokens
			delete(tokens, authToken)

			return false
		}

		return true
	}
}

func (ts *TokenStore) RemoveToken(authToken string) {
	tokens := *ts.tokens
	if _, exists := tokens[authToken]; exists { // TODO: We have a race condition here
		delete(tokens, authToken)
	}
}
