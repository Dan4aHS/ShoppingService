package grpchandlers

import (
	shopping "ShoppingExpensesService/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateAddRequest(req *shopping.AddRequest) error {
	if req.GetTgId() == 0 {
		return status.Errorf(codes.InvalidArgument, "TgId is required")
	}
	if req.GetQrInfo() == "" {
		return status.Errorf(codes.InvalidArgument, "QrInfo is required")
	}
	return nil
}

func ValidateListRequest(req *shopping.ListRequest) error {
	if req.GetTgId() == 0 {
		return status.Errorf(codes.InvalidArgument, "TgId is required")
	}
	if req.GetBeginTime() == "" {
		return status.Errorf(codes.InvalidArgument, "BeginTime is required")
	}
	if req.GetEndTime() == "" {
		return status.Errorf(codes.InvalidArgument, "EndTime is required")
	}
	return nil
}
