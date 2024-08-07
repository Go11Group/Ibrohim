// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: collaborations.proto

package collaborations

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

// CollaborationsClient is the client API for Collaborations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollaborationsClient interface {
	CreateInvitation(ctx context.Context, in *CreateInvite, opts ...grpc.CallOption) (*ID, error)
	RespondInvitation(ctx context.Context, in *CreateCollaboration, opts ...grpc.CallOption) (*ID, error)
	GetCollaboratorsByPodcastId(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Collaborators, error)
	UpdateCollaboratorByPodcastId(ctx context.Context, in *UpdateCollaborator, opts ...grpc.CallOption) (*Void, error)
	DeleteCollaboratorByPodcastId(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Void, error)
}

type collaborationsClient struct {
	cc grpc.ClientConnInterface
}

func NewCollaborationsClient(cc grpc.ClientConnInterface) CollaborationsClient {
	return &collaborationsClient{cc}
}

func (c *collaborationsClient) CreateInvitation(ctx context.Context, in *CreateInvite, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/Collaborations/CreateInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationsClient) RespondInvitation(ctx context.Context, in *CreateCollaboration, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/Collaborations/RespondInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationsClient) GetCollaboratorsByPodcastId(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/Collaborations/GetCollaboratorsByPodcastId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationsClient) UpdateCollaboratorByPodcastId(ctx context.Context, in *UpdateCollaborator, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Collaborations/UpdateCollaboratorByPodcastId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationsClient) DeleteCollaboratorByPodcastId(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Collaborations/DeleteCollaboratorByPodcastId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollaborationsServer is the server API for Collaborations service.
// All implementations must embed UnimplementedCollaborationsServer
// for forward compatibility
type CollaborationsServer interface {
	CreateInvitation(context.Context, *CreateInvite) (*ID, error)
	RespondInvitation(context.Context, *CreateCollaboration) (*ID, error)
	GetCollaboratorsByPodcastId(context.Context, *ID) (*Collaborators, error)
	UpdateCollaboratorByPodcastId(context.Context, *UpdateCollaborator) (*Void, error)
	DeleteCollaboratorByPodcastId(context.Context, *Ids) (*Void, error)
	mustEmbedUnimplementedCollaborationsServer()
}

// UnimplementedCollaborationsServer must be embedded to have forward compatible implementations.
type UnimplementedCollaborationsServer struct {
}

func (UnimplementedCollaborationsServer) CreateInvitation(context.Context, *CreateInvite) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvitation not implemented")
}
func (UnimplementedCollaborationsServer) RespondInvitation(context.Context, *CreateCollaboration) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondInvitation not implemented")
}
func (UnimplementedCollaborationsServer) GetCollaboratorsByPodcastId(context.Context, *ID) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaboratorsByPodcastId not implemented")
}
func (UnimplementedCollaborationsServer) UpdateCollaboratorByPodcastId(context.Context, *UpdateCollaborator) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollaboratorByPodcastId not implemented")
}
func (UnimplementedCollaborationsServer) DeleteCollaboratorByPodcastId(context.Context, *Ids) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollaboratorByPodcastId not implemented")
}
func (UnimplementedCollaborationsServer) mustEmbedUnimplementedCollaborationsServer() {}

// UnsafeCollaborationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollaborationsServer will
// result in compilation errors.
type UnsafeCollaborationsServer interface {
	mustEmbedUnimplementedCollaborationsServer()
}

func RegisterCollaborationsServer(s grpc.ServiceRegistrar, srv CollaborationsServer) {
	s.RegisterService(&Collaborations_ServiceDesc, srv)
}

func _Collaborations_CreateInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationsServer).CreateInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Collaborations/CreateInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationsServer).CreateInvitation(ctx, req.(*CreateInvite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collaborations_RespondInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollaboration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationsServer).RespondInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Collaborations/RespondInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationsServer).RespondInvitation(ctx, req.(*CreateCollaboration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collaborations_GetCollaboratorsByPodcastId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationsServer).GetCollaboratorsByPodcastId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Collaborations/GetCollaboratorsByPodcastId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationsServer).GetCollaboratorsByPodcastId(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collaborations_UpdateCollaboratorByPodcastId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollaborator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationsServer).UpdateCollaboratorByPodcastId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Collaborations/UpdateCollaboratorByPodcastId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationsServer).UpdateCollaboratorByPodcastId(ctx, req.(*UpdateCollaborator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collaborations_DeleteCollaboratorByPodcastId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationsServer).DeleteCollaboratorByPodcastId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Collaborations/DeleteCollaboratorByPodcastId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationsServer).DeleteCollaboratorByPodcastId(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

// Collaborations_ServiceDesc is the grpc.ServiceDesc for Collaborations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collaborations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Collaborations",
	HandlerType: (*CollaborationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvitation",
			Handler:    _Collaborations_CreateInvitation_Handler,
		},
		{
			MethodName: "RespondInvitation",
			Handler:    _Collaborations_RespondInvitation_Handler,
		},
		{
			MethodName: "GetCollaboratorsByPodcastId",
			Handler:    _Collaborations_GetCollaboratorsByPodcastId_Handler,
		},
		{
			MethodName: "UpdateCollaboratorByPodcastId",
			Handler:    _Collaborations_UpdateCollaboratorByPodcastId_Handler,
		},
		{
			MethodName: "DeleteCollaboratorByPodcastId",
			Handler:    _Collaborations_DeleteCollaboratorByPodcastId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collaborations.proto",
}
