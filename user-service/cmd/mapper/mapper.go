package mapper

import (
	"user-service/cmd/model"
	"user-service/cmd/proto"
)

func ProtoToModel(request *proto.UserRequest) model.User {
	user := model.User{
		Name:     request.GetName(),
		Surname:  request.GetSurname(),
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}
	return user
}

func ModelToProto(model *model.User) proto.UserResponse {
	proto := proto.UserResponse{
		Name:    model.Name,
		Surname: model.Surname,
		Id:      model.Id,
	}
	return proto
}
