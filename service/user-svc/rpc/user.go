package rpc

import (
	"context"
	proto "sisyphus/proto/user"
	"sisyphus/service/user-svc/handler"
)

type Service struct {
}

func (s *Service) AddUser(ctx context.Context, req *proto.AddUserRequest) (rsp *proto.AddUserResponse, err error) {
	reqUser := req.User
	user := handler.User{
		Username: reqUser.Username,
		Password: reqUser.Password,
		Email:    reqUser.Email,
		Phone:    reqUser.Phone,
		State:    int(reqUser.State),
		Profile: handler.Profile{
			Nickname: reqUser.Profile.Nickname,
			Age:      int8(reqUser.Profile.Age),
			Gender:   reqUser.Profile.Gender,
			Address:  reqUser.Profile.Address,
		},
	}
	err = user.Add()
	if err != nil {
		rsp = &proto.AddUserResponse{
			Success: false,
			Error: &proto.Error{
				Code:   500,
				Detail: err.Error(),
			},
		}

		return
	}
	rsp = &proto.AddUserResponse{
		Id:      user.ID,
		Success: true,
	}
	return rsp, nil
}
