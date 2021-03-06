// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Unary
	LoadProducts(ctx context.Context, in *LoadProductsRequest, opts ...grpc.CallOption) (*LoadProductsResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	// Bi-directional steaming
	GetProductsInfinite(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductsInfiniteClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) LoadProducts(ctx context.Context, in *LoadProductsRequest, opts ...grpc.CallOption) (*LoadProductsResponse, error) {
	out := new(LoadProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/LoadProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductsInfinite(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductsInfiniteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/product.ProductService/GetProductsInfinite", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetProductsInfiniteClient{stream}
	return x, nil
}

type ProductService_GetProductsInfiniteClient interface {
	Send(*GetProductsRequest) error
	Recv() (*GetProductsResponse, error)
	grpc.ClientStream
}

type productServiceGetProductsInfiniteClient struct {
	grpc.ClientStream
}

func (x *productServiceGetProductsInfiniteClient) Send(m *GetProductsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceGetProductsInfiniteClient) Recv() (*GetProductsResponse, error) {
	m := new(GetProductsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// Unary
	LoadProducts(context.Context, *LoadProductsRequest) (*LoadProductsResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	// Bi-directional steaming
	GetProductsInfinite(ProductService_GetProductsInfiniteServer) error
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) LoadProducts(context.Context, *LoadProductsRequest) (*LoadProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadProducts not implemented")
}
func (UnimplementedProductServiceServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductServiceServer) GetProductsInfinite(ProductService_GetProductsInfiniteServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductsInfinite not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_LoadProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).LoadProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/LoadProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).LoadProducts(ctx, req.(*LoadProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductsInfinite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).GetProductsInfinite(&productServiceGetProductsInfiniteServer{stream})
}

type ProductService_GetProductsInfiniteServer interface {
	Send(*GetProductsResponse) error
	Recv() (*GetProductsRequest, error)
	grpc.ServerStream
}

type productServiceGetProductsInfiniteServer struct {
	grpc.ServerStream
}

func (x *productServiceGetProductsInfiniteServer) Send(m *GetProductsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceGetProductsInfiniteServer) Recv() (*GetProductsRequest, error) {
	m := new(GetProductsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadProducts",
			Handler:    _ProductService_LoadProducts_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProductsInfinite",
			Handler:       _ProductService_GetProductsInfinite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "product.proto",
}
