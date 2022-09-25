package extauthz

import (
	"context"
	authv3 "github.com/Dimss/envoy-authz/gen/proto/go/envoy/service/auth/v3"
)

type Service struct {
	authv3.UnimplementedAuthorizationServer
}

func (s *Service) Check(ctx context.Context, request *authv3.CheckRequest) (*authv3.CheckResponse, error) {

	rep := &authv3.CheckResponse{Status: &authv3.OkHttpResponse{}}
	return rep, nil
}
