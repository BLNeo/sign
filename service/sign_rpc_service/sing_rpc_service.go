package sign_rpc_service

import (
	"context"
	"errors"
	signPb "github.com/BLNeo/protobuf-grpc-file/sign"
	"sign/tool/util"
	"time"
)

func NewRealSignRpc() signPb.SignServer {
	return &RealSignRpc{}
}

type RealSignRpc struct{}

func (r *RealSignRpc) VerifyToken(ctx context.Context, request *signPb.VerifyTokenRequest) (*signPb.VerifyTokenRespond, error) {
	resp := &signPb.VerifyTokenRespond{}

	claims, err := util.ParseToken(request.Token)
	if err != nil {
		return resp, err
	} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return resp, errors.New("token have expired")
	}

	resp.Enabled = true
	return resp, nil
}
