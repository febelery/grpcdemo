package auth

import "context"

type Token struct {
	AppID     string
	AppSecret string
}

func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_id": t.AppID, "app_secret": t.AppSecret}, nil
}

func (t *Token) RequireTransportSecurity() bool {
	return true
}
