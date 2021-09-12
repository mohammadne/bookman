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
	// TokenValid(context.Context, string, TokenType) failures.Failure
}

type jwt struct {
	config *Config
	logger logger.Logger
	tracer trace.Tracer
}

func New(cfg *Config, lg logger.Logger, tr trace.Tracer) Jwt {
	return &jwt{config: cfg, logger: lg, tracer: tr}
}

var (
	createTokenFailure   = failures.Jwt{}.NewInternal("token creation failed, please try later")
	invalidClaimsFailure = failures.Jwt{}.NewInternal("token metadata is invalid, please try later")
	invalidTokenFailure  = failures.Jwt{}.NewInvalid("invalid token given")
)

func (jwt *jwt) CreateJwt(ctx context.Context, userId uint64) (*models.Jwt, failures.Failure) {
	_, span := jwt.tracer.Start(ctx, "jwt.create_jwt")
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
		jwt.logger.Error(createTokenFailure.Message(), logger.Error(accessErr))
		span.RecordError(accessErr)
		return nil, createTokenFailure
	}

	refreshErr := createToken(userId, jwt.config.RefreshSecret, tokenDetail.RefreshToken)
	if refreshErr != nil {
		jwt.logger.Error(createTokenFailure.Message(), logger.Error(refreshErr))
		span.RecordError(refreshErr)
		return nil, createTokenFailure
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

func (jwt *jwt) ExtractTokenMetadata(ctx context.Context, tokenString string, tokenType TokenType) (*models.AccessDetails, failures.Failure) {
	_, span := jwt.tracer.Start(ctx, "jwt.extract_token_metadata")
	defer span.End()

	token, err := jwt.verifyToken(tokenString, tokenType)
	if err != nil {
		jwt.logger.Error(invalidTokenFailure.Message(), logger.Error(err))
		span.RecordError(err)
		return nil, invalidTokenFailure
	}

	claims, ok := token.Claims.(jwtPkg.MapClaims)
	if !ok {
		err = fmt.Errorf("invalid map claims: %v", token.Claims)
		jwt.logger.Error(invalidClaimsFailure.Message(), logger.Error(err))
		span.RecordError(err)
		return nil, invalidClaimsFailure
	}

	tokenUuid, ok := claims["token_uuid"].(string)
	if !ok {
		err = fmt.Errorf("ivalid claims: %v", claims)
		jwt.logger.Error(invalidTokenFailure.Message(), logger.Error(err))
		span.RecordError(err)
		return nil, invalidTokenFailure
	}

	userIdStr := fmt.Sprintf("%.f", claims["user_id"])
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("ivalid claims: %v, err:%v", claims, err)
		jwt.logger.Error(invalidTokenFailure.Message(), logger.Error(err))
		span.RecordError(err)
		return nil, invalidTokenFailure
	}

	return &models.AccessDetails{TokenUuid: tokenUuid, UserId: userId}, nil
}

func (jwt *jwt) TokenValid(ctx context.Context, tokenString string, tokenType TokenType) failures.Failure {
	_, span := jwt.tracer.Start(ctx, "jwt.token_valid")
	defer span.End()

	_, err := jwt.verifyToken(tokenString, tokenType)
	if err != nil {
		jwt.logger.Error(invalidTokenFailure.Message(), logger.Error(err))
		span.RecordError(err)
		return invalidTokenFailure
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

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

// checkSigningMethod checks the token method conform to "SigningMethodHMAC"
func (jwt *jwt) checkSigningMethod(secret string) jwtPkg.Keyfunc {
	return func(token *jwtPkg.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtPkg.SigningMethodHMAC); !ok {
			s := token.Header["alg"]
			jwt.logger.Error("unexpected signing method", logger.Any("token", s))
			return nil, fmt.Errorf("unexpected signing method: %v", s)
		}

		return []byte(secret), nil
	}
}
