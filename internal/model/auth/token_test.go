package auth

import "testing"

func TestTokenStore(t *testing.T) {
	ts := NewTokenStore()
	token := ts.NewToken()

	if token == "" {
		t.Errorf("Created invalid token: %s", token)
	}

	if ts.CheckToken(token) != true {
		t.Errorf("Token not found: %s", token)
	}

	if ts.CheckToken("foo") != false {
		t.Errorf("Found unexpected token")
	}

	ts.RemoveToken(token)
	if ts.CheckToken(token) != false {
		t.Errorf("Token not deleted: %s", token)
	}
}
