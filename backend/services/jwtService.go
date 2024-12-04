package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId string, name string, admin bool) string
}

type JWTCustomClaims struct {
	Name string `json:"name"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}

type jwtServiceStruct struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtServiceStruct{
		secretKey: getSecretKey(),
		issuer: "email@example.com",
	}
}

func getSecretKey() (secret string) {
	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	return secret
}

func (jwtSrv *jwtServiceStruct) GenerateToken(userId string, username string, admin bool) string {
	claims := &JWTCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer: jwtSrv.issuer,
			IssuedAt: time.Now().Unix(),
			Id: userId,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}