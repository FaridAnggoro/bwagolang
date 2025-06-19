package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
	secretKey string
}

func NewService() *jwtService {
	return &jwtService{
		secretKey: "BWASTARTUP_s3cr3t_k3y",
	}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
