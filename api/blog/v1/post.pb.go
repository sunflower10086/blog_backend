// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: blog/v1/post.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 获取博客列表的请求
type ListPostsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	mi := &file_blog_v1_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_blog_v1_post_proto_rawDescGZIP(), []int{0}
}

func (x *ListPostsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 获取博客列表的响应
type ListPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Posts         []*PostBaseInfo        `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	mi := &file_blog_v1_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_blog_v1_post_proto_rawDescGZIP(), []int{1}
}

func (x *ListPostsResponse) GetPosts() []*PostBaseInfo {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *ListPostsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 获取单个博客的请求
type GetPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_blog_v1_post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_blog_v1_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type PostBaseInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cover         string                 `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Tags          []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     int32                  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int32                  `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Categories    string                 `protobuf:"bytes,8,opt,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostBaseInfo) Reset() {
	*x = PostBaseInfo{}
	mi := &file_blog_v1_post_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBaseInfo) ProtoMessage() {}

func (x *PostBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_post_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBaseInfo.ProtoReflect.Descriptor instead.
func (*PostBaseInfo) Descriptor() ([]byte, []int) {
	return file_blog_v1_post_proto_rawDescGZIP(), []int{3}
}

func (x *PostBaseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostBaseInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostBaseInfo) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *PostBaseInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PostBaseInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostBaseInfo) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PostBaseInfo) GetUpdatedAt() int32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *PostBaseInfo) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

// 博客实体
type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BaseInfo      *PostBaseInfo          `protobuf:"bytes,1,opt,name=base_info,json=baseInfo,proto3" json:"base_info,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_blog_v1_post_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_post_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_blog_v1_post_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetBaseInfo() *PostBaseInfo {
	if x != nil {
		return x.BaseInfo
	}
	return nil
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_blog_v1_post_proto protoreflect.FileDescriptor

var file_blog_v1_post_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xb7,
	0x01, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x23, 0x5a, 0x21, 0x73, 0x75, 0x6e, 0x66,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_v1_post_proto_rawDescOnce sync.Once
	file_blog_v1_post_proto_rawDescData = file_blog_v1_post_proto_rawDesc
)

func file_blog_v1_post_proto_rawDescGZIP() []byte {
	file_blog_v1_post_proto_rawDescOnce.Do(func() {
		file_blog_v1_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_post_proto_rawDescData)
	})
	return file_blog_v1_post_proto_rawDescData
}

var file_blog_v1_post_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blog_v1_post_proto_goTypes = []any{
	(*ListPostsRequest)(nil),  // 0: blog.v1.ListPostsRequest
	(*ListPostsResponse)(nil), // 1: blog.v1.ListPostsResponse
	(*GetPostRequest)(nil),    // 2: blog.v1.GetPostRequest
	(*PostBaseInfo)(nil),      // 3: blog.v1.PostBaseInfo
	(*Post)(nil),              // 4: blog.v1.Post
}
var file_blog_v1_post_proto_depIdxs = []int32{
	3, // 0: blog.v1.ListPostsResponse.posts:type_name -> blog.v1.PostBaseInfo
	3, // 1: blog.v1.Post.base_info:type_name -> blog.v1.PostBaseInfo
	0, // 2: blog.v1.Poster.ListPosts:input_type -> blog.v1.ListPostsRequest
	2, // 3: blog.v1.Poster.GetPost:input_type -> blog.v1.GetPostRequest
	1, // 4: blog.v1.Poster.ListPosts:output_type -> blog.v1.ListPostsResponse
	4, // 5: blog.v1.Poster.GetPost:output_type -> blog.v1.Post
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blog_v1_post_proto_init() }
func file_blog_v1_post_proto_init() {
	if File_blog_v1_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_v1_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_v1_post_proto_goTypes,
		DependencyIndexes: file_blog_v1_post_proto_depIdxs,
		MessageInfos:      file_blog_v1_post_proto_msgTypes,
	}.Build()
	File_blog_v1_post_proto = out.File
	file_blog_v1_post_proto_rawDesc = nil
	file_blog_v1_post_proto_goTypes = nil
	file_blog_v1_post_proto_depIdxs = nil
}
