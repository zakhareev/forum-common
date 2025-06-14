// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: forum.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Forum_CreateCategory_FullMethodName = "/forum.Forum/CreateCategory"
	Forum_ListCategories_FullMethodName = "/forum.Forum/ListCategories"
	Forum_CreateTopic_FullMethodName    = "/forum.Forum/CreateTopic"
	Forum_ListTopics_FullMethodName     = "/forum.Forum/ListTopics"
	Forum_CreatePost_FullMethodName     = "/forum.Forum/CreatePost"
	Forum_ListPosts_FullMethodName      = "/forum.Forum/ListPosts"
	Forum_CreateComment_FullMethodName  = "/forum.Forum/CreateComment"
	Forum_ListComments_FullMethodName   = "/forum.Forum/ListComments"
)

// ForumClient is the client API for Forum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForumClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	ListCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
}

type forumClient struct {
	cc grpc.ClientConnInterface
}

func NewForumClient(cc grpc.ClientConnInterface) ForumClient {
	return &forumClient{cc}
}

func (c *forumClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, Forum_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) ListCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, Forum_ListCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, Forum_CreateTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, Forum_ListTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, Forum_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, Forum_ListPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, Forum_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCommentsResponse)
	err := c.cc.Invoke(ctx, Forum_ListComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForumServer is the server API for Forum service.
// All implementations must embed UnimplementedForumServer
// for forward compatibility.
type ForumServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)
	ListCategories(context.Context, *Empty) (*ListCategoriesResponse, error)
	CreateTopic(context.Context, *CreateTopicRequest) (*TopicResponse, error)
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error)
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CommentResponse, error)
	ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error)
	mustEmbedUnimplementedForumServer()
}

// UnimplementedForumServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedForumServer struct{}

func (UnimplementedForumServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedForumServer) ListCategories(context.Context, *Empty) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedForumServer) CreateTopic(context.Context, *CreateTopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedForumServer) ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedForumServer) CreatePost(context.Context, *CreatePostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedForumServer) ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedForumServer) CreateComment(context.Context, *CreateCommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedForumServer) ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedForumServer) mustEmbedUnimplementedForumServer() {}
func (UnimplementedForumServer) testEmbeddedByValue()               {}

// UnsafeForumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumServer will
// result in compilation errors.
type UnsafeForumServer interface {
	mustEmbedUnimplementedForumServer()
}

func RegisterForumServer(s grpc.ServiceRegistrar, srv ForumServer) {
	// If the following call pancis, it indicates UnimplementedForumServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Forum_ServiceDesc, srv)
}

func _Forum_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_ListCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).ListCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).ListPosts(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forum_ListComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).ListComments(ctx, req.(*ListCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Forum_ServiceDesc is the grpc.ServiceDesc for Forum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Forum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forum.Forum",
	HandlerType: (*ForumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _Forum_CreateCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _Forum_ListCategories_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _Forum_CreateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _Forum_ListTopics_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Forum_CreatePost_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _Forum_ListPosts_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Forum_CreateComment_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _Forum_ListComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forum.proto",
}
