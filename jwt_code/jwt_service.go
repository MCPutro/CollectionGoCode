package jwt_code

import "github.com/golang-jwt/jwt/v4"

type JwtService interface {
	GenerateToken(UserId string) string
	ValidateToken(token string) (*jwt.Token, error)
}
