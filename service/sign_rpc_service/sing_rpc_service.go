package sign_rpc_service

import (
	"context"
	"errors"
	"sign/models/user"
	signPb "sign/proto/sign"
	"sign/tool/util"
	"time"
)

func NewRealSignRpc() signPb.SignServer {
	return &RealSignRpc{}
}

type RealSignRpc struct{}

func (r *RealSignRpc) UserInfo(ctx context.Context, request *signPb.UserInfoRequest) (*signPb.UserInfoRespond, error) {
	if request.UserId == 0 {
		return nil, errors.New("userId is empty")
	}
	_, err := user.NewIUser().Get(request.UserId)
	if err != nil {
		return nil, err
	}
	resp := &signPb.UserInfoRespond{
		//Id:       info.ID,
		//Name:     info.Name,
		//Phone:    info.Phone,
		//Email:    info.Email,
		//Avatar:   info.Avatar,
		//Gender:   info.Gender,
		//Nickname: info.Nickname,
	}
	return resp, err
}

func (r *RealSignRpc) VerifyToken(ctx context.Context, request *signPb.VerifyTokenRequest) (*signPb.VerifyTokenRespond, error) {

	claims, err := util.ParseToken(request.Token)
	if err != nil {
		return nil, err
	} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, errors.New("token have expired")
	}
	resp := &signPb.VerifyTokenRespond{
		//UserId:   claims.UserId,
		//UserName: claims.UserName,
	}
	return resp, nil
}
