// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: product.proto

package product

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

// ProductesClient is the client API for Productes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductesClient interface {
	CreateProduct(ctx context.Context, in *NewProduct, opts ...grpc.CallOption) (*InsertResp, error)
	GetProductById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdateResp, error)
	DeleteProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	FetchProducts(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Products, error)
}

type productesClient struct {
	cc grpc.ClientConnInterface
}

func NewProductesClient(cc grpc.ClientConnInterface) ProductesClient {
	return &productesClient{cc}
}

func (c *productesClient) CreateProduct(ctx context.Context, in *NewProduct, opts ...grpc.CallOption) (*InsertResp, error) {
	out := new(InsertResp)
	err := c.cc.Invoke(ctx, "/product.Productes/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productesClient) GetProductById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.Productes/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productesClient) UpdateProduct(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/product.Productes/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productesClient) DeleteProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/product.Productes/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productesClient) FetchProducts(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/product.Productes/FetchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductesServer is the server API for Productes service.
// All implementations must embed UnimplementedProductesServer
// for forward compatibility
type ProductesServer interface {
	CreateProduct(context.Context, *NewProduct) (*InsertResp, error)
	GetProductById(context.Context, *Id) (*Product, error)
	UpdateProduct(context.Context, *NewData) (*UpdateResp, error)
	DeleteProduct(context.Context, *Id) (*Void, error)
	FetchProducts(context.Context, *Filter) (*Products, error)
	mustEmbedUnimplementedProductesServer()
}

// UnimplementedProductesServer must be embedded to have forward compatible implementations.
type UnimplementedProductesServer struct {
}

func (UnimplementedProductesServer) CreateProduct(context.Context, *NewProduct) (*InsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductesServer) GetProductById(context.Context, *Id) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProductesServer) UpdateProduct(context.Context, *NewData) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductesServer) DeleteProduct(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductesServer) FetchProducts(context.Context, *Filter) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchProducts not implemented")
}
func (UnimplementedProductesServer) mustEmbedUnimplementedProductesServer() {}

// UnsafeProductesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductesServer will
// result in compilation errors.
type UnsafeProductesServer interface {
	mustEmbedUnimplementedProductesServer()
}

func RegisterProductesServer(s grpc.ServiceRegistrar, srv ProductesServer) {
	s.RegisterService(&Productes_ServiceDesc, srv)
}

func _Productes_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductesServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Productes/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductesServer).CreateProduct(ctx, req.(*NewProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productes_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductesServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Productes/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductesServer).GetProductById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productes_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductesServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Productes/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductesServer).UpdateProduct(ctx, req.(*NewData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productes_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductesServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Productes/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductesServer).DeleteProduct(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productes_FetchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductesServer).FetchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Productes/FetchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductesServer).FetchProducts(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Productes_ServiceDesc is the grpc.ServiceDesc for Productes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Productes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Productes",
	HandlerType: (*ProductesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Productes_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _Productes_GetProductById_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Productes_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Productes_DeleteProduct_Handler,
		},
		{
			MethodName: "FetchProducts",
			Handler:    _Productes_FetchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
