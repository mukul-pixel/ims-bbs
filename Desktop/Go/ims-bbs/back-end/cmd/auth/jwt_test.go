package auth

import (
	"fmt"
	"testing"

	"github.com/mukul-pixel/ims-bbs/cmd/config"
)

func TestJWTCreation(t *testing.T) {
	secret := []byte(config.Envs.JWTSecret)

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("expected token not to be empty")
	}

	fmt.Println("token:", token)
}
