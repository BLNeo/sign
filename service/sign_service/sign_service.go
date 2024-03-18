package sign_service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sign/models"
	"sign/models/user"
	"sign/tool/redis"
	"sign/tool/util"
)

func NewSignService() ISignService {
	return &SignService{
		iUser:  user.NewIUser(),
		iRedis: redis.NewIRedis(),
	}
}

type ISignService interface {
	SignUp(in *user.SignUpRequest) error
	SignIn(in *user.SignInRequest) (*user.SignInRespond, error)
}

type SignService struct {
	iUser  user.IUser
	iRedis redis.IRedis
}

func (s *SignService) SignIn(in *user.SignInRequest) (*user.SignInRespond, error) {
	userInfo, err := s.iUser.GetByPhone(in.Phone)
	if err != nil {
		return nil, err
	}

	if util.EncodeMD5(userInfo.PassSalt+in.Password) != userInfo.Password {
		return nil, errors.New("手机号或密码错误")
	}

	// redis 校验
	tokenKey := fmt.Sprintf("token_%d", userInfo.ID)
	tokenValue := s.iRedis.Get(tokenKey)
	if tokenValue == "" {
		tokenValue, err = util.GenerateToken(userInfo.ID, userInfo.Name)
		if err != nil {
			return nil, err
		}
		err = s.iRedis.Set(tokenKey, tokenValue, util.TokenExpiresAt)
		if err != nil {
			return nil, err
		}
	}

	resp := &user.SignInRespond{
		Id:       userInfo.ID,
		Name:     userInfo.Name,
		Phone:    userInfo.Phone,
		Email:    userInfo.Email,
		Avatar:   userInfo.Avatar,
		Gender:   userInfo.Gender,
		Nickname: userInfo.Nickname,
		Token:    tokenValue,
	}
	return resp, nil
}

func (s *SignService) SignUp(in *user.SignUpRequest) error {
	userDao := user.NewIUser()
	// 校验手机是否存在
	exist, err := userDao.PhoneExist(in.Phone)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("手机号已注册")
	}

	uuidString := uuid.New().String()

	insertDate := &models.User{
		Name:     uuidString[24:],
		Phone:    in.Phone,
		Password: util.EncodeMD5(uuidString + in.Password),
		PassSalt: uuidString,
		Email:    "",
		Avatar:   "",
		Gender:   "",
		Nickname: "",
	}
	return userDao.Create(insertDate)
}
