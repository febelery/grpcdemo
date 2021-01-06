package auth

import (
	"context"
	"errors"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenInfo struct {
	ID    string
	Roles []string
}

func Interceptor(ctx context.Context) (context.Context, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}

	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)

	return newCtx, nil
}

func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}

		return tokenInfo, nil
	}

	return tokenInfo, errors.New("Token invalid: bearer " + token)
}
