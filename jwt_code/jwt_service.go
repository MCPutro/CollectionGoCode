package jwt_code

import "github.com/golang-jwt/jwt/v4"

type JwtService interface {
	GenerateToken(ID string) string
	ValidateToken(Token string) (*jwt.Token, error)
}
