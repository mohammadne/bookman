package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jwtPkg "github.com/golang-jwt/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	uuid "github.com/satori/go.uuid"
)

type TokenType uint8

const (
	Access TokenType = iota
	Refresh
)

type Jwt interface {
	CreateJwt(userId uint64) (*models.Jwt, error)
	ExtractTokenMetadata(tokenString string, tokenType TokenType) (*models.AccessDetails, error)
	TokenValid(tokenString string, tokenType TokenType) error
}

type jwt struct {
	config *Config
	logger logger.Logger
}

func New(config *Config, logger logger.Logger) Jwt {
	return &jwt{config: config, logger: logger}
}

func (jwt *jwt) CreateJwt(userId uint64) (*models.Jwt, error) {
	accessExpires := time.Duration(jwt.config.AccessExpires) * time.Hour
	refreshExpires := time.Duration(jwt.config.RefreshExpires) * time.Hour

	tokenDetail := &models.Jwt{
		AccessToken: &models.Token{
			UUID:    uuid.NewV4().String(),
			Expires: time.Now().Add(accessExpires).Unix(),
		},
		RefreshToken: &models.Token{
			UUID:    uuid.NewV4().String(),
			Expires: time.Now().Add(refreshExpires).Unix(),
		},
	}

	accessErr := createToken(userId, jwt.config.AccessSecret, tokenDetail.AccessToken)
	if accessErr != nil {
		return nil, accessErr
	}

	refreshErr := createToken(userId, jwt.config.RefreshSecret, tokenDetail.RefreshToken)
	if refreshErr != nil {
		return nil, refreshErr
	}

	return tokenDetail, nil
}

func createToken(userId uint64, secret string, token *models.Token) error {
	claims := jwtPkg.MapClaims{
		"authorized": true,
		"token_uuid": token.UUID,
		"user_id":    userId,
		"exp":        token.Expires,
	}

	var err error
	at := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	token.Token, err = at.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return nil
}

func (jwt *jwt) ExtractTokenMetadata(tokenString string, tokenType TokenType) (*models.AccessDetails, error) {
	token, err := jwt.verifyToken(tokenString, tokenType)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwtPkg.MapClaims)
	if ok && token.Valid {
		tokenUuid, ok := claims["token_uuid"].(string)
		if !ok {
			return nil, err
		}

		userIdStr := fmt.Sprintf("%.f", claims["user_id"])
		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			return nil, err
		}

		return &models.AccessDetails{TokenUuid: tokenUuid, UserId: userId}, nil
	}

	return nil, errors.New("invalid token")
}

func (jwt *jwt) TokenValid(tokenString string, tokenType TokenType) error {
	token, err := jwt.verifyToken(tokenString, tokenType)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (jwt *jwt) verifyToken(tokenString string, tokenType TokenType) (*jwtPkg.Token, error) {
	var secret string
	if tokenType == Access {
		secret = jwt.config.AccessSecret
	} else {
		secret = jwt.config.RefreshSecret
	}

	token, err := jwtPkg.Parse(tokenString, jwt.checkSigningMethod(secret))
	if err != nil {
		return nil, err
	}

	return token, nil
}

//checkSigningMethod checks the token method conform to "SigningMethodHMAC"
func (jwt *jwt) checkSigningMethod(secret string) jwtPkg.Keyfunc {
	return func(token *jwtPkg.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtPkg.SigningMethodHMAC); !ok {
			s := token.Header["alg"]
			jwt.logger.Error("unexpected signing method", logger.Unknown("token", s))
			return nil, fmt.Errorf("unexpected signing method: %v", s)
		}

		return []byte(secret), nil
	}
}
