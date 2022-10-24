package mapper

import (
	"user-registration/cmd/model"
	"user-registration/cmd/proto"
	proto_address "user-registration/cmd/proto/address-service"
	proto_user "user-registration/cmd/proto/user-service"
)

func RegReqToUsrReq(req *proto.RegistrationRequest) *proto_user.UserRequest {
	userRequst := proto_user.UserRequest{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
	}

	return &userRequst
}

func RegReqToAdrReq(req *proto.RegistrationRequest) proto_address.AddressRequest {
	addressRequst := proto_address.AddressRequest{
		Address: req.Address,
	}

	return addressRequst
}

func ProtoToModel(req *proto.RegistrationRequest) model.User {
	userModel := model.User{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
	}

	return userModel
}
