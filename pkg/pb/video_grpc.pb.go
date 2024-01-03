// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.10.0
// source: video.proto

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

const (
	VideoService_GetPublishVideoList_FullMethodName  = "/VideoService/GetPublishVideoList"
	VideoService_PublishVideo_FullMethodName         = "/VideoService/PublishVideo"
	VideoService_GetFeedList_FullMethodName          = "/VideoService/GetFeedList"
	VideoService_GetVideoInfoList_FullMethodName     = "/VideoService/GetVideoInfoList"
	VideoService_GetFavoriteVideoList_FullMethodName = "/VideoService/GetFavoriteVideoList"
	VideoService_UpdateFavoriteCount_FullMethodName  = "/VideoService/UpdateFavoriteCount"
	VideoService_UpdateCommentCount_FullMethodName   = "/VideoService/UpdateCommentCount"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	// 获取视频列表
	GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error)
	PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
	GetFeedList(ctx context.Context, in *GetFeedListRequest, opts ...grpc.CallOption) (*GetFeedListResponse, error)
	GetVideoInfoList(ctx context.Context, in *GetVideoInfoListReq, opts ...grpc.CallOption) (*GetVideoInfoListRsp, error)
	GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListReq, opts ...grpc.CallOption) (*GetFavoriteVideoListRsp, error)
	// 更新这个视频的获赞数
	UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountRsp, error)
	// 更新这个视频的评论数
	UpdateCommentCount(ctx context.Context, in *UpdateCommentCountReq, opts ...grpc.CallOption) (*UpdateCommentCountRsp, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error) {
	out := new(GetPublishVideoListResponse)
	err := c.cc.Invoke(ctx, VideoService_GetPublishVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	out := new(PublishVideoResponse)
	err := c.cc.Invoke(ctx, VideoService_PublishVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetFeedList(ctx context.Context, in *GetFeedListRequest, opts ...grpc.CallOption) (*GetFeedListResponse, error) {
	out := new(GetFeedListResponse)
	err := c.cc.Invoke(ctx, VideoService_GetFeedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideoInfoList(ctx context.Context, in *GetVideoInfoListReq, opts ...grpc.CallOption) (*GetVideoInfoListRsp, error) {
	out := new(GetVideoInfoListRsp)
	err := c.cc.Invoke(ctx, VideoService_GetVideoInfoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListReq, opts ...grpc.CallOption) (*GetFavoriteVideoListRsp, error) {
	out := new(GetFavoriteVideoListRsp)
	err := c.cc.Invoke(ctx, VideoService_GetFavoriteVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountRsp, error) {
	out := new(UpdateFavoriteCountRsp)
	err := c.cc.Invoke(ctx, VideoService_UpdateFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) UpdateCommentCount(ctx context.Context, in *UpdateCommentCountReq, opts ...grpc.CallOption) (*UpdateCommentCountRsp, error) {
	out := new(UpdateCommentCountRsp)
	err := c.cc.Invoke(ctx, VideoService_UpdateCommentCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations should embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	// 获取视频列表
	GetPublishVideoList(context.Context, *GetPublishVideoListRequest) (*GetPublishVideoListResponse, error)
	PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error)
	GetFeedList(context.Context, *GetFeedListRequest) (*GetFeedListResponse, error)
	GetVideoInfoList(context.Context, *GetVideoInfoListReq) (*GetVideoInfoListRsp, error)
	GetFavoriteVideoList(context.Context, *GetFavoriteVideoListReq) (*GetFavoriteVideoListRsp, error)
	// 更新这个视频的获赞数
	UpdateFavoriteCount(context.Context, *UpdateFavoriteCountReq) (*UpdateFavoriteCountRsp, error)
	// 更新这个视频的评论数
	UpdateCommentCount(context.Context, *UpdateCommentCountReq) (*UpdateCommentCountRsp, error)
}

// UnimplementedVideoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) GetPublishVideoList(context.Context, *GetPublishVideoListRequest) (*GetPublishVideoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishVideoList not implemented")
}
func (UnimplementedVideoServiceServer) PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedVideoServiceServer) GetFeedList(context.Context, *GetFeedListRequest) (*GetFeedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedList not implemented")
}
func (UnimplementedVideoServiceServer) GetVideoInfoList(context.Context, *GetVideoInfoListReq) (*GetVideoInfoListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfoList not implemented")
}
func (UnimplementedVideoServiceServer) GetFavoriteVideoList(context.Context, *GetFavoriteVideoListReq) (*GetFavoriteVideoListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoList not implemented")
}
func (UnimplementedVideoServiceServer) UpdateFavoriteCount(context.Context, *UpdateFavoriteCountReq) (*UpdateFavoriteCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavoriteCount not implemented")
}
func (UnimplementedVideoServiceServer) UpdateCommentCount(context.Context, *UpdateCommentCountReq) (*UpdateCommentCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommentCount not implemented")
}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_GetPublishVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishVideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetPublishVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetPublishVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetPublishVideoList(ctx, req.(*GetPublishVideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_PublishVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).PublishVideo(ctx, req.(*PublishVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetFeedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetFeedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetFeedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetFeedList(ctx, req.(*GetFeedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideoInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideoInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetVideoInfoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideoInfoList(ctx, req.(*GetVideoInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetFavoriteVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteVideoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetFavoriteVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetFavoriteVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetFavoriteVideoList(ctx, req.(*GetFavoriteVideoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_UpdateFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoriteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).UpdateFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_UpdateFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).UpdateFavoriteCount(ctx, req.(*UpdateFavoriteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_UpdateCommentCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).UpdateCommentCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_UpdateCommentCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).UpdateCommentCount(ctx, req.(*UpdateCommentCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublishVideoList",
			Handler:    _VideoService_GetPublishVideoList_Handler,
		},
		{
			MethodName: "PublishVideo",
			Handler:    _VideoService_PublishVideo_Handler,
		},
		{
			MethodName: "GetFeedList",
			Handler:    _VideoService_GetFeedList_Handler,
		},
		{
			MethodName: "GetVideoInfoList",
			Handler:    _VideoService_GetVideoInfoList_Handler,
		},
		{
			MethodName: "GetFavoriteVideoList",
			Handler:    _VideoService_GetFavoriteVideoList_Handler,
		},
		{
			MethodName: "UpdateFavoriteCount",
			Handler:    _VideoService_UpdateFavoriteCount_Handler,
		},
		{
			MethodName: "UpdateCommentCount",
			Handler:    _VideoService_UpdateCommentCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
