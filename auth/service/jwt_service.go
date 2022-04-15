package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

type JwtService struct {
	secret                string
	expirationTimeMinutes int
	issure                string
}

func NewJwtService(config *viper.Viper) *JwtService {
	return &JwtService{
		secret:                config.GetString("jwt.secret"),
		issure:                config.GetString("jwt.issure"),
		expirationTimeMinutes: config.GetInt("jwt.expiration_time_minutes"),
	}
}

type Claim struct {
	sum uint `json:"sum"`
	jwt.StandardClaims
}

func (service *JwtService) CreateToken(sum uint) (string, error) {
	claim := &Claim{
		sum: sum,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(service.expirationTimeMinutes) * time.Minute).Unix(),
			Issuer:    service.issure,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(service.secret))
}

func (service *JwtService) VerifyToken(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(service.secret), nil
	})
	return err == nil
}
