package grpchandlers

import (
	fns_api "ShoppingExpensesService/internal/clients/fns-api"
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/models"
	"ShoppingExpensesService/internal/service"
	shopping "ShoppingExpensesService/pkg"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type ServerAPI struct {
	shopping.UnimplementedShoppingServiceServer
	ss  service.IProductService
	cfg *config.Config
}

func RegisterGRPCServer(server *grpc.Server, ss service.IProductService, cfg *config.Config) {
	shopping.RegisterShoppingServiceServer(server, &ServerAPI{ss: ss, cfg: cfg})
}

func (s *ServerAPI) Add(ctx context.Context, req *shopping.AddRequest) (*shopping.AddResponse, error) {
	if err := ValidateAddRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	products, err := fns_api.GetQRInfo(int(req.GetTgId()), req.GetQrInfo(), s.cfg)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	log.Printf("Products: %v", products)
	count, err := s.ss.AddProducts(ctx, products)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	log.Printf("Added %d products", count)
	return &shopping.AddResponse{
		Count: int32(count),
	}, nil
}

func (s *ServerAPI) List(ctx context.Context, req *shopping.ListRequest) (*shopping.ListResponse, error) {
	if err := ValidateListRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	products, err := s.ss.ListProducts(ctx, int(req.GetTgId()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	purchases := models.ProductsListDBToGRPC(products)
	return &shopping.ListResponse{
		Purchases: purchases,
	}, nil
}
