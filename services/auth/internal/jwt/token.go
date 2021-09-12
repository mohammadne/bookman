package jwt

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	jwtPkg "github.com/golang-jwt/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/failures"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel/trace"
)

type TokenType uint8

const (
	Access TokenType = iota
	Refresh
)

type Jwt interface {
	CreateJwt(context.Context, uint64) (*models.Jwt, failures.Failure)
	ExtractTokenMetadata(context.Context, string, TokenType) (*models.AccessDetails, failures.Failure)
	TokenValid(context.Context, string, TokenType) failures.Failure
}

type jwt struct {
	config *Config
	logger logger.Logger
	tracer trace.Tracer
}

func New(cfg *Config, lg logger.Logger, tr trace.Tracer) Jwt {
	return &jwt{config: cfg, logger: lg, tracer: tr}
}

func (jwt *jwt) CreateJwt(ctx context.Context, userId uint64) (*models.Jwt, failures.Failure) {
	ctx, span := jwt.tracer.Start(ctx, "jwt.create_jwt")
	defer span.End()

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
		failure := Failure{}.NewUnprocessableEntity("unprocessable access token")
		jwt.logger.Error(failure.Message(), logger.Error(accessErr))
		span.RecordError(accessErr)
		return nil, failure
	}

	refreshErr := createToken(userId, jwt.config.RefreshSecret, tokenDetail.RefreshToken)
	if refreshErr != nil {
		failure := Failure{}.NewUnprocessableEntity("unprocessable refresh token")
		jwt.logger.Error(failure.Message(), logger.Error(refreshErr))
		span.RecordError(refreshErr)
		return nil, failure
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

// failureUnprocessableEntity
// 	failureUnautorized         = failures.Network{}.NewUnauthorized("unauthorized")
// failureUnautorized
func (jwt *jwt) ExtractTokenMetadata(ctx context.Context, tokenString string, tokenType TokenType) (*models.AccessDetails, failures.Failure) {
	ctx, span := jwt.tracer.Start(ctx, "jwt.extract_token_metadata")
	defer span.End()

	token, failure := jwt.verifyToken(tokenString, tokenType)
	if failure != nil {
		span.RecordError(failure)
		return nil, failure
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
			span.RecordError(err)
			return nil, err
		}

		return &models.AccessDetails{TokenUuid: tokenUuid, UserId: userId}, nil
	}

	err = errors.New("invalid token")
	span.RecordError(err)
	return nil, err
}

func (jwt *jwt) TokenValid(ctx context.Context, tokenString string, tokenType TokenType) failures.Failure {
	token, err := jwt.verifyToken(tokenString, tokenType)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (jwt *jwt) verifyToken(tokenString string, tokenType TokenType) (*jwtPkg.Token, failures.Failure) {
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

// checkSigningMethod checks the token method conform to "SigningMethodHMAC"
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
