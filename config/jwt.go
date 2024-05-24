package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("aowjdvnh9128d9218nsda12")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
