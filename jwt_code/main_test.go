package jwt_code

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
)

func TestJwt(t *testing.T) {
	jwtService := NewJwtServiceImpl("secret-key", "MCP")

	//generate JWT
	token := jwtService.GenerateToken("email@gmail.com")
	fmt.Println("token :", token)

	fmt.Println("-----------------------------------")
	//validation JWT
	tek, err := jwtService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

	if err != nil {
		fmt.Println(err)
		return
	}

	if tek.Valid {
		claims := tek.Claims.(jwt.MapClaims)
		log.Println("Claim[id]  : ", claims["ID"])
		log.Println("Claim[iss] : ", claims["iss"])
	}
}
