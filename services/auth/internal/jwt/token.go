package jwt

import (
	"os"
	"time"

	jwtPkg "github.com/golang-jwt/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/go-pkgs/logger"
	uuid "github.com/satori/go.uuid"
)

type Jwt interface {
	CreateTokenDetail(userId uint64) (*models.TokenDetails, error)
}

type jwt struct {
	config *Config
	logger logger.Logger
}

func New(config *Config, logger logger.Logger) Jwt {
	return &jwt{config: config, logger: logger}
}

func (jwt *jwt) CreateTokenDetail(userId uint64) (*models.TokenDetails, error) {
	accessExpires := time.Duration(jwt.config.AccessExpires) * time.Hour
	refreshExpires := time.Duration(jwt.config.RefreshExpires) * time.Hour

	tokenDetail := &models.TokenDetails{
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

func createToken(userId uint64, secretKey string, token *models.Token) error {
	claims := jwtPkg.MapClaims{
		"authorized": true,
		"token_uuid": token.UUID,
		"user_id":    userId,
		"exp":        token.Expires,
	}

	var err error
	at := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	token.Token, err = at.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		return err
	}

	return nil
}
