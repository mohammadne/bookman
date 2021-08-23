package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	uuid "github.com/satori/go.uuid"
)

const (
	accessSecretKey  = "ACCESS_SECRET"
	aceessExpireTime = time.Minute * 15

	refreshSecretKey  = "REFRESH_SECRET"
	refreshExpireTime = time.Hour * 24 * 7
)

func CreateTokenDetail(userId uint64) (*models.TokenDetails, error) {
	tokenDetail := &models.TokenDetails{
		AccessToken: &models.Token{
			UUID:    uuid.NewV4().String(),
			Expires: time.Now().Add(aceessExpireTime).Unix(),
		},
		RefreshToken: &models.Token{
			UUID:    uuid.NewV4().String(),
			Expires: time.Now().Add(refreshExpireTime).Unix(),
		},
	}

	accessErr := createToken(userId, accessSecretKey, tokenDetail.AccessToken)
	if accessErr != nil {
		return nil, accessErr
	}

	refreshErr := createToken(userId, refreshSecretKey, tokenDetail.RefreshToken)
	if refreshErr != nil {
		return nil, refreshErr
	}

	return tokenDetail, nil
}

func createToken(userId uint64, secretKey string, token *models.Token) error {
	claims := jwt.MapClaims{
		"authorized": true,
		"token_uuid": token.UUID,
		"user_id":    userId,
		"exp":        token.Expires,
	}

	var err error
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Token, err = at.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		return err
	}

	return nil
}
