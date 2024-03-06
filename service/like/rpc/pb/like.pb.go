// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: like.proto

package pb

import (
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

type ThumbsUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId    string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`        // 业务id
	ObjId    int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`       // 点赞对象id
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
	LikeType int32  `protobuf:"varint,4,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *ThumbsUpRequest) Reset() {
	*x = ThumbsUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbsUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbsUpRequest) ProtoMessage() {}

func (x *ThumbsUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbsUpRequest.ProtoReflect.Descriptor instead.
func (*ThumbsUpRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *ThumbsUpRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *ThumbsUpRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *ThumbsUpRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ThumbsUpRequest) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

type ThumbsUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId      string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`            // 业务id
	ObjId      int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`           // 点赞对象id
	LikeNum    int64  `protobuf:"varint,3,opt,name=likeNum,proto3" json:"likeNum,omitempty"`       // 点赞数
	DislikeNum int64  `protobuf:"varint,4,opt,name=dislikeNum,proto3" json:"dislikeNum,omitempty"` // 点踩数
}

func (x *ThumbsUpResponse) Reset() {
	*x = ThumbsUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbsUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbsUpResponse) ProtoMessage() {}

func (x *ThumbsUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbsUpResponse.ProtoReflect.Descriptor instead.
func (*ThumbsUpResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *ThumbsUpResponse) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *ThumbsUpResponse) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *ThumbsUpResponse) GetLikeNum() int64 {
	if x != nil {
		return x.LikeNum
	}
	return 0
}

func (x *ThumbsUpResponse) GetDislikeNum() int64 {
	if x != nil {
		return x.DislikeNum
	}
	return 0
}

type IsThumbsUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId    string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`        // 业务id
	TargetId int64  `protobuf:"varint,2,opt,name=targetId,proto3" json:"targetId,omitempty"` // 点赞对象id
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
}

func (x *IsThumbsUpRequest) Reset() {
	*x = IsThumbsUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsThumbsUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsThumbsUpRequest) ProtoMessage() {}

func (x *IsThumbsUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsThumbsUpRequest.ProtoReflect.Descriptor instead.
func (*IsThumbsUpRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *IsThumbsUpRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *IsThumbsUpRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *IsThumbsUpRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsThumbsUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserThumbsUps map[int64]*UserThumbsUp `protobuf:"bytes,1,rep,name=userThumbsUps,proto3" json:"userThumbsUps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IsThumbsUpResponse) Reset() {
	*x = IsThumbsUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsThumbsUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsThumbsUpResponse) ProtoMessage() {}

func (x *IsThumbsUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsThumbsUpResponse.ProtoReflect.Descriptor instead.
func (*IsThumbsUpResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *IsThumbsUpResponse) GetUserThumbsUps() map[int64]*UserThumbsUp {
	if x != nil {
		return x.UserThumbsUps
	}
	return nil
}

type UserThumbsUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ThumbsUpTime int64 `protobuf:"varint,2,opt,name=ThumbsUpTime,proto3" json:"ThumbsUpTime,omitempty"`
	LikeType     int32 `protobuf:"varint,3,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *UserThumbsUp) Reset() {
	*x = UserThumbsUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserThumbsUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserThumbsUp) ProtoMessage() {}

func (x *UserThumbsUp) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserThumbsUp.ProtoReflect.Descriptor instead.
func (*UserThumbsUp) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *UserThumbsUp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserThumbsUp) GetThumbsUpTime() int64 {
	if x != nil {
		return x.ThumbsUpTime
	}
	return 0
}

func (x *UserThumbsUp) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f,
	0x62, 0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x10, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x73, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6b, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e,
	0x75, 0x6d, 0x22, 0x5d, 0x0a, 0x11, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xc3, 0x01, 0x0a, 0x12, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x73, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x73, 0x1a, 0x57,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32,
	0x8e, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x73, 0x55, 0x70, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x49, 0x73, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x73, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_like_proto_goTypes = []interface{}{
	(*ThumbsUpRequest)(nil),    // 0: service.ThumbsUpRequest
	(*ThumbsUpResponse)(nil),   // 1: service.ThumbsUpResponse
	(*IsThumbsUpRequest)(nil),  // 2: service.IsThumbsUpRequest
	(*IsThumbsUpResponse)(nil), // 3: service.IsThumbsUpResponse
	(*UserThumbsUp)(nil),       // 4: service.UserThumbsUp
	nil,                        // 5: service.IsThumbsUpResponse.UserThumbsUpsEntry
}
var file_like_proto_depIdxs = []int32{
	5, // 0: service.IsThumbsUpResponse.userThumbsUps:type_name -> service.IsThumbsUpResponse.UserThumbsUpsEntry
	4, // 1: service.IsThumbsUpResponse.UserThumbsUpsEntry.value:type_name -> service.UserThumbsUp
	0, // 2: service.Like.ThumbsUp:input_type -> service.ThumbsUpRequest
	2, // 3: service.Like.IsThumbsUp:input_type -> service.IsThumbsUpRequest
	1, // 4: service.Like.ThumbsUp:output_type -> service.ThumbsUpResponse
	3, // 5: service.Like.IsThumbsUp:output_type -> service.IsThumbsUpResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThumbsUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThumbsUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsThumbsUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsThumbsUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserThumbsUp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}