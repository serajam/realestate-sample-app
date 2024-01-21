/* Copyright (C) Fedir Petryk */

package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/auth"
	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type Auth struct {
	usersRepository UsersRepository
	redis           *redis.Client
	accessToken     auth.TokenGenerator
	refreshToken    auth.TokenGenerator
	logger          *zap.SugaredLogger
}

func NewAuthSrv(
	usersRepository UsersRepository, redis *redis.Client, accessToken auth.TokenGenerator,
	refreshToken auth.TokenGenerator, logger *zap.SugaredLogger,
) Auth {
	return Auth{
		usersRepository: usersRepository, logger: logger, accessToken: accessToken, refreshToken: refreshToken,
		redis: redis,
	}
}

func (s Auth) Refresh(ctx context.Context, userID int, tokenUUID, deviceID string) (auth.AccessToken, auth.RefreshToken, error) {
	user, err := s.usersRepository.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", domainErrors.User{Msg: domainErrors.MsgUserNotExists}
		}

		s.logger.Errorw(err.Error(), "user_id", userID, "method", "Refresh")
		return "", "", domainErrors.Internal
	}

	tokenMetaStr, err := s.redis.Get(ctx, tokenUUID).Result()
	if err != nil {
		s.logger.Errorw(err.Error(), "token_uuid", tokenUUID, "method", "Refresh")
		return "", "", domainErrors.Internal
	}

	meta := auth.TokenMeta{}
	err = json.Unmarshal([]byte(tokenMetaStr), &meta)
	if err != nil {
		s.logger.Errorw(err.Error(), "token_uuid", tokenUUID, "method", "Refresh")
		return "", "", domainErrors.Internal
	}

	if meta.TokenType != auth.TypeRefreshToken {
		return "", "", domainErrors.User{Msg: domainErrors.MsgInvalidToken}
	}

	err = s.redis.Del(ctx, tokenUUID).Err()
	if err != nil {
		s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
		return "", "", domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return s.generateTokens(ctx, user, deviceID)
}

func (s Auth) Authenticate(ctx context.Context, email, pwd, deviceID string) (
	auth.AccessToken, auth.RefreshToken, error,
) {
	user, err := s.usersRepository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", domainErrors.User{Msg: domainErrors.MsgWrongEmailOrPwd}
		}

		s.logger.Errorw(err.Error(), "email", email, "method", "Authenticate")
		return "", "", domainErrors.User{Msg: domainErrors.MsgWrongEmailOrPwd}
	}

	if !user.CheckPasswordHash(pwd) {
		return "", "", domainErrors.User{Msg: domainErrors.MsgWrongEmailOrPwd}
	}

	return s.generateTokens(ctx, user, deviceID)
}

func (s Auth) SignOut(ctx context.Context, userID int, deviceID string) error {
	if deviceID == "" {
		return s.removeAllTokens(ctx, userID)
	}

	return s.removeDeviceTokens(ctx, userID, deviceID)
}

func (s Auth) SignOutAll(ctx context.Context, userID int) error {
	return s.removeAllTokens(ctx, userID)
}

func (s Auth) generateTokens(ctx context.Context, user *domain.User, deviceID string) (auth.AccessToken, auth.RefreshToken, error) {
	logger := s.logger.With("method", "generateTokens")
	accessToken, err := s.accessToken.Generate(user)
	if err != nil {
		logger.Errorw(err.Error(), "email", user.Email)
		return "", "", domainErrors.User{Msg: domainErrors.MsgWrongEmailOrPwd}
	}

	accessTokenMeta := auth.TokenMeta{
		UserID:    user.ID,
		TokenType: auth.TypeAccessToken,
		DeviceId:  deviceID,
		Token:     accessToken.Token,
	}
	accessTokenMetaBytes, err := json.Marshal(accessTokenMeta)
	if err != nil {
		logger.Error(err.Error())
		return "", "", domainErrors.Internal
	}

	err = s.redis.Set(ctx, accessToken.TokenUUID, accessTokenMetaBytes, accessToken.Ttl).Err()
	if err != nil {
		s.logger.Error(err.Error())
		return "", "", domainErrors.Internal
	}

	refreshToken, err := s.refreshToken.Generate(user)
	if err != nil {
		logger.Errorw(err.Error(), "email", user.Email)
		return "", "", domainErrors.User{Msg: domainErrors.MsgWrongEmailOrPwd}
	}

	refreshTokenMeta := auth.TokenMeta{
		UserID:    user.ID,
		TokenType: auth.TypeRefreshToken,
		DeviceId:  deviceID,
		Token:     refreshToken.Token,
	}
	refreshTokenMetaBytes, err := json.Marshal(refreshTokenMeta)
	if err != nil {
		logger.Error(err.Error())
		return "", "", domainErrors.Internal
	}

	err = s.redis.Set(ctx, refreshToken.TokenUUID, refreshTokenMetaBytes, refreshToken.Ttl).Err()
	if err != nil {
		logger.Error(err.Error())
		return "", "", domainErrors.Internal
	}

	err = s.redis.
		SAdd(ctx, strconv.Itoa(user.ID), refreshToken.TokenUUID, accessToken.TokenUUID).
		Err()
	if err != nil {
		logger.Error(err.Error())
		return "", "", domainErrors.Internal
	}

	return auth.AccessToken(accessToken.Token), auth.RefreshToken(refreshToken.Token), nil
}

func (s Auth) removeDeviceTokens(ctx context.Context, userID int, deviceID string) error {
	genKey := strconv.Itoa(userID)
	tokens, err := s.redis.SMembers(ctx, genKey).Result()
	if err != nil {
		s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
		return domainErrors.Internal
	}

	if len(tokens) == 0 {
		return nil
	}

	for _, token := range tokens {
		res, err := s.redis.Get(ctx, token).Result()
		if err != nil {
			s.logger.Debugw(err.Error(), "method", "removeDeviceTokens")
			err = s.redis.SRem(ctx, genKey, tokens).Err()
			if err != nil {
				s.logger.Errorw(err.Error(), "userID", userID, "method", "removeDeviceTokens")
			}
		}

		meta := auth.TokenMeta{}
		err = json.Unmarshal([]byte(res), &meta)
		if err != nil {
			s.logger.Debugw(err.Error(), "method", "removeDeviceTokens")
			err = s.redis.SRem(ctx, genKey, tokens).Err()
			if err != nil {
				s.logger.Errorw(err.Error(), "userID", userID, "method", "removeDeviceTokens")
			}
		}

		if meta.DeviceId != deviceID {
			continue
		}

		err = s.redis.Del(ctx, token).Err()
		if err != nil {
			s.logger.Errorw(err.Error(), "userID", userID, "method", "removeDeviceTokens")
			return domainErrors.Internal
		}

		err = s.redis.SRem(ctx, genKey, token).Err()
		if err != nil {
			s.logger.Errorw(err.Error(), "userID", userID, "method", "removeDeviceTokens")
			return domainErrors.Internal
		}
	}

	// err = s.redis.Del(ctx, strconv.Itoa(userID)).Err()
	// if err != nil {
	// 	s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
	// 	return domainErrors.Internal
	// }

	return nil
}

func (s Auth) removeAllTokens(ctx context.Context, userID int) error {
	tokens, err := s.redis.SMembers(ctx, strconv.Itoa(userID)).Result()
	if err != nil {
		s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
		return domainErrors.Internal
	}

	if len(tokens) == 0 {
		return nil
	}

	for _, token := range tokens {
		err = s.redis.Del(ctx, token).Err()
		if err != nil {
			s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
			return domainErrors.Internal
		}
	}

	err = s.redis.Del(ctx, strconv.Itoa(userID)).Err()
	if err != nil {
		s.logger.Errorw(err.Error(), "userID", userID, "method", "SignOut")
		return domainErrors.Internal
	}

	return nil
}

func (s Auth) Validate(ctx context.Context, tokenStr string) (auth.Token, error) {
	s.logger.Debugw("token", "token", tokenStr)

	claims, ok, err := s.accessToken.VerifyAndExtract(tokenStr)
	s.logger.Debugw("token", "claims", claims)
	if err != nil {
		s.logger.Debugw(err.Error(), "method", "Validate")
		return auth.Token{}, domainErrors.User{Msg: domainErrors.MsgInvalidToken}
	}

	if !ok {
		s.logger.Debugw("token is not valid", "method", "Validate")
		return auth.Token{}, domainErrors.User{Msg: domainErrors.MsgInvalidToken}
	}

	token := auth.FromClaims(claims)

	_, err = s.redis.Get(ctx, claims.TokenUUID).Result()

	if errors.Is(err, redis.Nil) {
		s.logger.Debugw("token not found in redis", "method", "Validate")
		return auth.Token{}, domainErrors.User{Msg: domainErrors.MsgInvalidToken}
	}

	if err != nil {
		s.logger.Debugw(err.Error(), "method", "Validate")
		return auth.Token{}, domainErrors.Internal
	}

	_, err = s.usersRepository.Get(ctx, claims.ID)
	if err != nil {
		s.logger.Debugw(err.Error(), "method", "Validate")
		return auth.Token{}, domainErrors.User{Msg: domainErrors.MsgUserNotExists}
	}

	return token, nil
}
