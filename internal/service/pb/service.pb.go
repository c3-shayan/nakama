// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type Post struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content              string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SpoilerOf            string               `protobuf:"bytes,4,opt,name=spoiler_of,json=spoilerOf,proto3" json:"spoiler_of,omitempty"`
	Nsfw                 bool                 `protobuf:"varint,5,opt,name=nsfw,proto3" json:"nsfw,omitempty"`
	LikesCount           int32                `protobuf:"varint,6,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	CommentsCount        int32                `protobuf:"varint,7,opt,name=comments_count,json=commentsCount,proto3" json:"comments_count,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	User                 *User                `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
	Mine                 bool                 `protobuf:"varint,10,opt,name=mine,proto3" json:"mine,omitempty"`
	Liked                bool                 `protobuf:"varint,11,opt,name=liked,proto3" json:"liked,omitempty"`
	Subscribed           bool                 `protobuf:"varint,12,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetSpoilerOf() string {
	if m != nil {
		return m.SpoilerOf
	}
	return ""
}

func (m *Post) GetNsfw() bool {
	if m != nil {
		return m.Nsfw
	}
	return false
}

func (m *Post) GetLikesCount() int32 {
	if m != nil {
		return m.LikesCount
	}
	return 0
}

func (m *Post) GetCommentsCount() int32 {
	if m != nil {
		return m.CommentsCount
	}
	return 0
}

func (m *Post) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Post) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Post) GetMine() bool {
	if m != nil {
		return m.Mine
	}
	return false
}

func (m *Post) GetLiked() bool {
	if m != nil {
		return m.Liked
	}
	return false
}

func (m *Post) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

type TimelineItem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId               int64    `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Post                 *Post    `protobuf:"bytes,4,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimelineItem) Reset()         { *m = TimelineItem{} }
func (m *TimelineItem) String() string { return proto.CompactTextString(m) }
func (*TimelineItem) ProtoMessage()    {}
func (*TimelineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *TimelineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimelineItem.Unmarshal(m, b)
}
func (m *TimelineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimelineItem.Marshal(b, m, deterministic)
}
func (m *TimelineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimelineItem.Merge(m, src)
}
func (m *TimelineItem) XXX_Size() int {
	return xxx_messageInfo_TimelineItem.Size(m)
}
func (m *TimelineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TimelineItem.DiscardUnknown(m)
}

var xxx_messageInfo_TimelineItem proto.InternalMessageInfo

func (m *TimelineItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimelineItem) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TimelineItem) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *TimelineItem) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type Comment struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId               int64                `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Content              string               `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	LikesCount           int32                `protobuf:"varint,5,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	User                 *User                `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	Mine                 bool                 `protobuf:"varint,8,opt,name=mine,proto3" json:"mine,omitempty"`
	Liked                bool                 `protobuf:"varint,9,opt,name=liked,proto3" json:"liked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Comment) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetLikesCount() int32 {
	if m != nil {
		return m.LikesCount
	}
	return 0
}

func (m *Comment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Comment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Comment) GetMine() bool {
	if m != nil {
		return m.Mine
	}
	return false
}

func (m *Comment) GetLiked() bool {
	if m != nil {
		return m.Liked
	}
	return false
}

type Notification struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Actors               []string             `protobuf:"bytes,3,rep,name=actors,proto3" json:"actors,omitempty"`
	Type                 string               `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	PostId               int64                `protobuf:"varint,5,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Read                 bool                 `protobuf:"varint,6,opt,name=read,proto3" json:"read,omitempty"`
	IssuedAt             *timestamp.Timestamp `protobuf:"bytes,7,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Notification) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Notification) GetActors() []string {
	if m != nil {
		return m.Actors
	}
	return nil
}

func (m *Notification) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Notification) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *Notification) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *Notification) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Post)(nil), "pb.Post")
	proto.RegisterType((*TimelineItem)(nil), "pb.TimelineItem")
	proto.RegisterType((*Comment)(nil), "pb.Comment")
	proto.RegisterType((*Notification)(nil), "pb.Notification")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x6a, 0xdc, 0x3e,
	0x10, 0xc7, 0xf1, 0x9f, 0x5d, 0xdb, 0xb3, 0x9b, 0x1c, 0xc4, 0x8f, 0x5f, 0xc4, 0xd2, 0x36, 0x66,
	0xa1, 0xe0, 0x93, 0x03, 0xed, 0xa1, 0xf4, 0x18, 0x72, 0xda, 0x4b, 0xff, 0x88, 0xe6, 0x6c, 0x64,
	0x5b, 0x0e, 0xa2, 0xb6, 0x64, 0x24, 0x39, 0xa5, 0x8f, 0xd0, 0x47, 0xeb, 0x93, 0xf4, 0x35, 0x8a,
	0x24, 0xbb, 0xec, 0x36, 0xa5, 0x0d, 0xa5, 0xb7, 0x99, 0xef, 0x8c, 0x66, 0x46, 0x9f, 0x2f, 0x9c,
	0x69, 0xa6, 0xee, 0x79, 0xc3, 0xca, 0x51, 0x49, 0x23, 0x51, 0x38, 0xd6, 0xbb, 0xcb, 0x3b, 0x29,
	0xef, 0x7a, 0x76, 0xe5, 0x94, 0x7a, 0xea, 0xae, 0x0c, 0x1f, 0x98, 0x36, 0x74, 0x18, 0x7d, 0xd3,
	0xfe, 0x3d, 0xc4, 0xb7, 0x9a, 0x29, 0x74, 0x0e, 0x21, 0x6f, 0x71, 0x90, 0x07, 0x45, 0x44, 0x42,
	0xde, 0xa2, 0x1d, 0xa4, 0x93, 0x66, 0x4a, 0xd0, 0x81, 0xe1, 0x30, 0x0f, 0x8a, 0x8c, 0xfc, 0xc8,
	0xd1, 0x53, 0x00, 0x7a, 0x4f, 0x0d, 0x55, 0xd5, 0xa4, 0x7a, 0x1c, 0xb9, 0x6a, 0xe6, 0x95, 0x5b,
	0xd5, 0xef, 0xbf, 0x85, 0x10, 0xbf, 0x93, 0xda, 0x3c, 0x98, 0x79, 0x01, 0x89, 0x9d, 0x51, 0xf1,
	0xd6, 0x8d, 0x8c, 0xc8, 0xda, 0xa6, 0x87, 0x16, 0x61, 0x48, 0x1a, 0x29, 0x0c, 0x13, 0x66, 0x9e,
	0xb6, 0xa4, 0x76, 0x95, 0x1e, 0x25, 0xef, 0x99, 0xaa, 0x64, 0x87, 0x63, 0xbf, 0x6a, 0x56, 0xde,
	0x76, 0x08, 0x41, 0x2c, 0x74, 0xf7, 0x09, 0xaf, 0xf2, 0xa0, 0x48, 0x89, 0x8b, 0xd1, 0x25, 0x6c,
	0x7a, 0xfe, 0x91, 0xe9, 0xaa, 0x91, 0x93, 0x30, 0x78, 0x9d, 0x07, 0xc5, 0x8a, 0x80, 0x93, 0x6e,
	0xac, 0x82, 0x9e, 0xc3, 0x79, 0x23, 0x87, 0x81, 0x09, 0xb3, 0xf4, 0x24, 0xae, 0xe7, 0x6c, 0x51,
	0x7d, 0xdb, 0x6b, 0x80, 0x46, 0x31, 0x6a, 0x58, 0x5b, 0x51, 0x83, 0xd3, 0x3c, 0x28, 0x36, 0x2f,
	0x76, 0xa5, 0xe7, 0x59, 0x2e, 0x3c, 0xcb, 0x0f, 0x0b, 0x4f, 0x92, 0xcd, 0xdd, 0xd7, 0x06, 0x3d,
	0x81, 0xd8, 0xfe, 0x0c, 0x67, 0xee, 0x51, 0x5a, 0x8e, 0x75, 0x69, 0x21, 0x13, 0xa7, 0xda, 0xa3,
	0x07, 0x2e, 0x18, 0x06, 0x7f, 0xb4, 0x8d, 0xd1, 0x7f, 0xb0, 0xb2, 0x17, 0xb6, 0x78, 0xe3, 0x44,
	0x9f, 0xa0, 0x67, 0x00, 0x7a, 0xaa, 0x75, 0xa3, 0x78, 0xcd, 0x5a, 0xbc, 0x75, 0xa5, 0x23, 0x65,
	0x2f, 0x60, 0x6b, 0xf7, 0xf7, 0x5c, 0xb0, 0x83, 0x61, 0xc3, 0xe3, 0x81, 0x5f, 0x40, 0x32, 0x4a,
	0x6d, 0x6c, 0x21, 0xf2, 0x05, 0x9b, 0x1e, 0x5a, 0x7b, 0xb9, 0x8d, 0x1c, 0xe9, 0xf9, 0x72, 0x6b,
	0x25, 0x71, 0xea, 0xfe, 0x4b, 0x08, 0xc9, 0x8d, 0x87, 0xf4, 0x0f, 0x76, 0x1d, 0xb9, 0x1e, 0x9f,
	0xba, 0xfe, 0x93, 0x85, 0xab, 0x07, 0x16, 0x9e, 0x7a, 0xb3, 0xfe, 0x1b, 0x6f, 0x92, 0xdf, 0x7a,
	0x93, 0xfe, 0xca, 0x9b, 0xec, 0xc8, 0x9b, 0xfd, 0xd7, 0x00, 0xb6, 0x6f, 0xa4, 0xe1, 0x1d, 0x6f,
	0xa8, 0xe1, 0x52, 0x3c, 0x1e, 0xc8, 0xff, 0xb0, 0xa6, 0x8d, 0x91, 0x4a, 0xe3, 0x28, 0x8f, 0x8a,
	0x8c, 0xcc, 0x99, 0xdd, 0x6d, 0x3e, 0x8f, 0x6c, 0x86, 0xe1, 0xe2, 0x63, 0x78, 0xab, 0x13, 0x78,
	0x08, 0x62, 0xc5, 0x68, 0xeb, 0xfe, 0x9e, 0x12, 0x17, 0xa3, 0x57, 0x90, 0x71, 0xad, 0x27, 0x0f,
	0x25, 0xf9, 0x23, 0x94, 0xd4, 0x37, 0x5f, 0x9b, 0x7a, 0xed, 0xaa, 0x2f, 0xbf, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x09, 0x8b, 0x21, 0x41, 0x04, 0x00, 0x00,
}
