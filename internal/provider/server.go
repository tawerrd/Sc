package provider

import (
	"context"
	"time"

	providerpb "github.com/synctv-org/synctv/proto/provider"
	"golang.org/x/oauth2"
)

type GRPCServer struct {
	providerpb.UnimplementedOauth2PluginServer
	Impl ProviderInterface
}

func (s *GRPCServer) Init(ctx context.Context, req *providerpb.InitReq) (*providerpb.Enpty, error) {
	s.Impl.Init(Oauth2Option{
		ClientID:     req.ClientId,
		ClientSecret: req.ClientSecret,
		RedirectURL:  req.RedirectUrl,
	})
	return &providerpb.Enpty{}, nil
}

func (s *GRPCServer) Provider(ctx context.Context, req *providerpb.Enpty) (*providerpb.ProviderResp, error) {
	return &providerpb.ProviderResp{Name: string(s.Impl.Provider())}, nil
}

func (s *GRPCServer) NewAuthURL(ctx context.Context, req *providerpb.NewAuthURLReq) (*providerpb.NewAuthURLResp, error) {
	return &providerpb.NewAuthURLResp{Url: s.Impl.NewAuthURL(req.State)}, nil
}

func (s *GRPCServer) GetToken(ctx context.Context, req *providerpb.GetTokenReq) (*providerpb.Token, error) {
	token, err := s.Impl.GetToken(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	return &providerpb.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry.Unix(),
	}, nil
}

func (s *GRPCServer) GetUserInfo(ctx context.Context, req *providerpb.GetUserInfoReq) (*providerpb.GetUserInfoResp, error) {
	userInfo, err := s.Impl.GetUserInfo(ctx, &oauth2.Token{
		AccessToken:  req.Token.AccessToken,
		TokenType:    req.Token.TokenType,
		Expiry:       time.Unix(req.Token.Expiry, 0),
		RefreshToken: req.Token.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
	resp := &providerpb.GetUserInfoResp{
		Username:       userInfo.Username,
		ProviderUserId: uint64(userInfo.ProviderUserID),
	}
	if userInfo.TokenRefreshed != nil {
		resp.TokenRefreshed = &providerpb.Token{
			AccessToken:  userInfo.TokenRefreshed.Token.AccessToken,
			TokenType:    userInfo.TokenRefreshed.Token.TokenType,
			RefreshToken: userInfo.TokenRefreshed.Token.RefreshToken,
			Expiry:       userInfo.TokenRefreshed.Token.Expiry.Unix(),
		}
	}
	return resp, nil
}
