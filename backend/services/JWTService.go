package services

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateJWTToken(userId string, username string, isAdmin bool) string
	ValidateJWTToken(token string) (*jwt.Token, error)
	GetUserIdFromToken(token string) (userId uint64, err error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "email@example.com",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	return secret
}

func (service *jwtService) GenerateJWTToken(userId string, username string, isAdmin bool) string {
	claims := &jwtCustomClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
			Id:        userId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return signedToken
}

func (service *jwtService) ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtService) GetUserIdFromToken(tokenString string) (userId uint64, err error) {
	token, err := service.ValidateJWTToken(tokenString)
	if err != nil {
		return 0, err
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		if jti, ok := claims["jti"].(string); ok {
			id, err := strconv.ParseUint(jti, 10, 64)
			if err != nil {
				return 0, err
			}
			return id, nil
		}
		return 0, nil
	} else {
		return 0, nil
	}
}
