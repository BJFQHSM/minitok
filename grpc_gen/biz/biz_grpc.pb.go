// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: idl/biz.proto

package biz

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

// BizServiceClient is the client API for BizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BizServiceClient interface {
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	QueryInfo(ctx context.Context, in *QueryInfoRequest, opts ...grpc.CallOption) (*QueryInfoResponse, error)
	PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error)
	QueryPublishList(ctx context.Context, in *QueryPublishListRequest, opts ...grpc.CallOption) (*QueryPublishListResponse, error)
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error)
	QueryFavoriteList(ctx context.Context, in *QueryFavoriteListRequest, opts ...grpc.CallOption) (*QueryFavoriteListResponse, error)
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	QueryCommentList(ctx context.Context, in *QueryCommentListRequest, opts ...grpc.CallOption) (*QueryCommentListResponse, error)
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
	QueryFollowList(ctx context.Context, in *QueryFollowListRequest, opts ...grpc.CallOption) (*QueryFollowListResponse, error)
	QueryFollowerList(ctx context.Context, in *QueryFollowerListRequest, opts ...grpc.CallOption) (*QueryFollowerListResponse, error)
}

type bizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBizServiceClient(cc grpc.ClientConnInterface) BizServiceClient {
	return &bizServiceClient{cc}
}

func (c *bizServiceClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryInfo(ctx context.Context, in *QueryInfoRequest, opts ...grpc.CallOption) (*QueryInfoResponse, error) {
	out := new(QueryInfoResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error) {
	out := new(PublishActionResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/PublishAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryPublishList(ctx context.Context, in *QueryPublishListRequest, opts ...grpc.CallOption) (*QueryPublishListResponse, error) {
	out := new(QueryPublishListResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryPublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error) {
	out := new(FavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryFavoriteList(ctx context.Context, in *QueryFavoriteListRequest, opts ...grpc.CallOption) (*QueryFavoriteListResponse, error) {
	out := new(QueryFavoriteListResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryFavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryCommentList(ctx context.Context, in *QueryCommentListRequest, opts ...grpc.CallOption) (*QueryCommentListResponse, error) {
	out := new(QueryCommentListResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	out := new(RelationActionResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/RelationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryFollowList(ctx context.Context, in *QueryFollowListRequest, opts ...grpc.CallOption) (*QueryFollowListResponse, error) {
	out := new(QueryFollowListResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryFollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) QueryFollowerList(ctx context.Context, in *QueryFollowerListRequest, opts ...grpc.CallOption) (*QueryFollowerListResponse, error) {
	out := new(QueryFollowerListResponse)
	err := c.cc.Invoke(ctx, "/grpc_biz.BizService/QueryFollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BizServiceServer is the server API for BizService service.
// All implementations must embed UnimplementedBizServiceServer
// for forward compatibility
type BizServiceServer interface {
	Feed(context.Context, *FeedRequest) (*FeedResponse, error)
	QueryInfo(context.Context, *QueryInfoRequest) (*QueryInfoResponse, error)
	PublishAction(context.Context, *PublishActionRequest) (*PublishActionResponse, error)
	QueryPublishList(context.Context, *QueryPublishListRequest) (*QueryPublishListResponse, error)
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error)
	QueryFavoriteList(context.Context, *QueryFavoriteListRequest) (*QueryFavoriteListResponse, error)
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	QueryCommentList(context.Context, *QueryCommentListRequest) (*QueryCommentListResponse, error)
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error)
	QueryFollowList(context.Context, *QueryFollowListRequest) (*QueryFollowListResponse, error)
	QueryFollowerList(context.Context, *QueryFollowerListRequest) (*QueryFollowerListResponse, error)
	mustEmbedUnimplementedBizServiceServer()
}

// UnimplementedBizServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBizServiceServer struct {
}

func (UnimplementedBizServiceServer) Feed(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedBizServiceServer) QueryInfo(context.Context, *QueryInfoRequest) (*QueryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInfo not implemented")
}
func (UnimplementedBizServiceServer) PublishAction(context.Context, *PublishActionRequest) (*PublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedBizServiceServer) QueryPublishList(context.Context, *QueryPublishListRequest) (*QueryPublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPublishList not implemented")
}
func (UnimplementedBizServiceServer) FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedBizServiceServer) QueryFavoriteList(context.Context, *QueryFavoriteListRequest) (*QueryFavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFavoriteList not implemented")
}
func (UnimplementedBizServiceServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedBizServiceServer) QueryCommentList(context.Context, *QueryCommentListRequest) (*QueryCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCommentList not implemented")
}
func (UnimplementedBizServiceServer) RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedBizServiceServer) QueryFollowList(context.Context, *QueryFollowListRequest) (*QueryFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFollowList not implemented")
}
func (UnimplementedBizServiceServer) QueryFollowerList(context.Context, *QueryFollowerListRequest) (*QueryFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFollowerList not implemented")
}
func (UnimplementedBizServiceServer) mustEmbedUnimplementedBizServiceServer() {}

// UnsafeBizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BizServiceServer will
// result in compilation errors.
type UnsafeBizServiceServer interface {
	mustEmbedUnimplementedBizServiceServer()
}

func RegisterBizServiceServer(s grpc.ServiceRegistrar, srv BizServiceServer) {
	s.RegisterService(&BizService_ServiceDesc, srv)
}

func _BizService_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryInfo(ctx, req.(*QueryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/PublishAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).PublishAction(ctx, req.(*PublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryPublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryPublishList(ctx, req.(*QueryPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).FavoriteAction(ctx, req.(*FavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryFavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryFavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryFavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryFavoriteList(ctx, req.(*QueryFavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryCommentList(ctx, req.(*QueryCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/RelationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryFollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryFollowList(ctx, req.(*QueryFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_QueryFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).QueryFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_biz.BizService/QueryFollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).QueryFollowerList(ctx, req.(*QueryFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BizService_ServiceDesc is the grpc.ServiceDesc for BizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_biz.BizService",
	HandlerType: (*BizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _BizService_Feed_Handler,
		},
		{
			MethodName: "QueryInfo",
			Handler:    _BizService_QueryInfo_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _BizService_PublishAction_Handler,
		},
		{
			MethodName: "QueryPublishList",
			Handler:    _BizService_QueryPublishList_Handler,
		},
		{
			MethodName: "FavoriteAction",
			Handler:    _BizService_FavoriteAction_Handler,
		},
		{
			MethodName: "QueryFavoriteList",
			Handler:    _BizService_QueryFavoriteList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _BizService_CommentAction_Handler,
		},
		{
			MethodName: "QueryCommentList",
			Handler:    _BizService_QueryCommentList_Handler,
		},
		{
			MethodName: "RelationAction",
			Handler:    _BizService_RelationAction_Handler,
		},
		{
			MethodName: "QueryFollowList",
			Handler:    _BizService_QueryFollowList_Handler,
		},
		{
			MethodName: "QueryFollowerList",
			Handler:    _BizService_QueryFollowerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/biz.proto",
}