package mapper

import (
	"address-service/cmd/model"
	"address-service/cmd/proto"
)

func ProtoToModel(req *proto.AddressRequest) model.Address {
	addressModel := model.Address{
		Location: req.Address,
	}
	return addressModel
}

func ModelToProto(model *model.Address) proto.AddressResponse {
	proto := proto.AddressResponse{
		Id:      model.Id,
		Address: model.Location,
	}
	return proto
}
