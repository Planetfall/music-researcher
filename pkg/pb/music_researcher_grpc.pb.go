// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package musicresearcher

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

// MusicResearcherClient is the client API for MusicResearcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicResearcherClient interface {
	Search(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Results, error)
	GetGenreList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenreList, error)
}

type musicResearcherClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicResearcherClient(cc grpc.ClientConnInterface) MusicResearcherClient {
	return &musicResearcherClient{cc}
}

func (c *musicResearcherClient) Search(ctx context.Context, in *Parameters, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/musicresearcher.MusicResearcher/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicResearcherClient) GetGenreList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenreList, error) {
	out := new(GenreList)
	err := c.cc.Invoke(ctx, "/musicresearcher.MusicResearcher/GetGenreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicResearcherServer is the server API for MusicResearcher service.
// All implementations must embed UnimplementedMusicResearcherServer
// for forward compatibility
type MusicResearcherServer interface {
	Search(context.Context, *Parameters) (*Results, error)
	GetGenreList(context.Context, *Empty) (*GenreList, error)
	mustEmbedUnimplementedMusicResearcherServer()
}

// UnimplementedMusicResearcherServer must be embedded to have forward compatible implementations.
type UnimplementedMusicResearcherServer struct {
}

func (UnimplementedMusicResearcherServer) Search(context.Context, *Parameters) (*Results, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedMusicResearcherServer) GetGenreList(context.Context, *Empty) (*GenreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenreList not implemented")
}
func (UnimplementedMusicResearcherServer) mustEmbedUnimplementedMusicResearcherServer() {}

// UnsafeMusicResearcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicResearcherServer will
// result in compilation errors.
type UnsafeMusicResearcherServer interface {
	mustEmbedUnimplementedMusicResearcherServer()
}

func RegisterMusicResearcherServer(s grpc.ServiceRegistrar, srv MusicResearcherServer) {
	s.RegisterService(&MusicResearcher_ServiceDesc, srv)
}

func _MusicResearcher_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicResearcherServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/musicresearcher.MusicResearcher/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicResearcherServer).Search(ctx, req.(*Parameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicResearcher_GetGenreList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicResearcherServer).GetGenreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/musicresearcher.MusicResearcher/GetGenreList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicResearcherServer).GetGenreList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicResearcher_ServiceDesc is the grpc.ServiceDesc for MusicResearcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicResearcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "musicresearcher.MusicResearcher",
	HandlerType: (*MusicResearcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _MusicResearcher_Search_Handler,
		},
		{
			MethodName: "GetGenreList",
			Handler:    _MusicResearcher_GetGenreList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Music-Researcher/api/music_researcher.proto",
}