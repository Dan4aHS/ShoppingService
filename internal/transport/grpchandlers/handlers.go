package grpchandlers

import (
	fns_api "ShoppingExpensesService/internal/clients/fns-api"
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/models"
	"ShoppingExpensesService/internal/service"
	shopping "ShoppingExpensesService/pkg"
	"ShoppingExpensesService/pkg/logging"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type ServerAPI struct {
	shopping.UnimplementedShoppingServiceServer
	ss  service.IProductService
	cl  fns_api.IClient
	cfg *config.Config
}

func RegisterGRPCServer(server *grpc.Server, ss service.IProductService, cfg *config.Config) {
	cl := fns_api.NewClient(cfg)
	shopping.RegisterShoppingServiceServer(server, &ServerAPI{ss: ss, cl: cl, cfg: cfg})
}

func (s *ServerAPI) Add(ctx context.Context, req *shopping.AddRequest) (*shopping.AddResponse, error) {
	if err := ValidateAddRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	ctx = logging.WithLogUserID(ctx, req.GetTgId())
	slog.InfoContext(ctx, "Add Request received")
	products, err := s.cl.GetQRInfo(int(req.GetTgId()), req.GetQrInfo())
	if err != nil {
		slog.ErrorContext(ctx, "Get QR Info error", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	slog.InfoContext(ctx, "Got QR Info")
	count, err := s.ss.AddProducts(ctx, products)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	slog.InfoContext(ctx, fmt.Sprintf("Added %d products", count))
	return &shopping.AddResponse{
		Count: int32(count),
	}, nil
}

func (s *ServerAPI) List(ctx context.Context, req *shopping.ListRequest) (*shopping.ListResponse, error) {
	if err := ValidateListRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	ctx = logging.WithLogUserID(ctx, req.GetTgId())
	slog.InfoContext(ctx, "List Request received")
	products, err := s.ss.ListProducts(ctx, int(req.GetTgId()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	purchases := models.ProductsListDBToGRPC(products)
	slog.InfoContext(ctx, "Listing Products")
	return &shopping.ListResponse{
		Purchases: purchases,
	}, nil
}
