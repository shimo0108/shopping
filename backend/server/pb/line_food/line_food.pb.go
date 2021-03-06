// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/line_food.proto

package server

import (
	reflect "reflect"
	sync "sync"

	pb_restaurant "github.com/shimo0108/shopping/backend/server/pb/restaurant"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LineFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId string                 `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	FoodId       string                 `protobuf:"bytes,3,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	OrderId      string                 `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Count        int32                  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Active       bool                   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *LineFood) Reset() {
	*x = LineFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineFood) ProtoMessage() {}

func (x *LineFood) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineFood.ProtoReflect.Descriptor instead.
func (*LineFood) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{0}
}

func (x *LineFood) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LineFood) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *LineFood) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

func (x *LineFood) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *LineFood) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LineFood) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *LineFood) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LineFood) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListLineFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLineFoodRequest) Reset() {
	*x = ListLineFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLineFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLineFoodRequest) ProtoMessage() {}

func (x *ListLineFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLineFoodRequest.ProtoReflect.Descriptor instead.
func (*ListLineFoodRequest) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{1}
}

type ListLineFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineFoodIds []string                  `protobuf:"bytes,1,rep,name=line_food_ids,json=lineFoodIds,proto3" json:"line_food_ids,omitempty"`
	Restaurant  *pb_restaurant.Restaurant `protobuf:"bytes,2,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	Count       int32                     `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Amount      int32                     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ListLineFoodResponse) Reset() {
	*x = ListLineFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLineFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLineFoodResponse) ProtoMessage() {}

func (x *ListLineFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLineFoodResponse.ProtoReflect.Descriptor instead.
func (*ListLineFoodResponse) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{2}
}

func (x *ListLineFoodResponse) GetLineFoodIds() []string {
	if x != nil {
		return x.LineFoodIds
	}
	return nil
}

func (x *ListLineFoodResponse) GetRestaurant() *pb_restaurant.Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *ListLineFoodResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListLineFoodResponse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateLineFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	FoodId       string `protobuf:"bytes,2,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	Count        int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateLineFoodRequest) Reset() {
	*x = CreateLineFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLineFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLineFoodRequest) ProtoMessage() {}

func (x *CreateLineFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLineFoodRequest.ProtoReflect.Descriptor instead.
func (*CreateLineFoodRequest) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLineFoodRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *CreateLineFoodRequest) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

func (x *CreateLineFoodRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateLineFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineFood *LineFood `protobuf:"bytes,1,opt,name=LineFood,proto3" json:"LineFood,omitempty"`
}

func (x *CreateLineFoodResponse) Reset() {
	*x = CreateLineFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLineFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLineFoodResponse) ProtoMessage() {}

func (x *CreateLineFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLineFoodResponse.ProtoReflect.Descriptor instead.
func (*CreateLineFoodResponse) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLineFoodResponse) GetLineFood() *LineFood {
	if x != nil {
		return x.LineFood
	}
	return nil
}

type ReplaceLineFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
}

func (x *ReplaceLineFoodRequest) Reset() {
	*x = ReplaceLineFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceLineFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceLineFoodRequest) ProtoMessage() {}

func (x *ReplaceLineFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceLineFoodRequest.ProtoReflect.Descriptor instead.
func (*ReplaceLineFoodRequest) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{5}
}

func (x *ReplaceLineFoodRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

type ReplaceLineFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineFoods []*LineFood `protobuf:"bytes,1,rep,name=line_foods,json=lineFoods,proto3" json:"line_foods,omitempty"`
}

func (x *ReplaceLineFoodResponse) Reset() {
	*x = ReplaceLineFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_line_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceLineFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceLineFoodResponse) ProtoMessage() {}

func (x *ReplaceLineFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_line_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceLineFoodResponse.ProtoReflect.Descriptor instead.
func (*ReplaceLineFoodResponse) Descriptor() ([]byte, []int) {
	return file_proto_line_food_proto_rawDescGZIP(), []int{6}
}

func (x *ReplaceLineFoodResponse) GetLineFoods() []*LineFood {
	if x != nil {
		return x.LineFoods
	}
	return nil
}

var File_proto_line_food_proto protoreflect.FileDescriptor

var file_proto_line_food_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x49, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x45, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x4c,
	0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x08,
	0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x22, 0x3d, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f,
	0x64, 0x73, 0x32, 0xe2, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69,
	0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x6b, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x0f, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x6d, 0x6f, 0x30, 0x31, 0x30, 0x38, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_line_food_proto_rawDescOnce sync.Once
	file_proto_line_food_proto_rawDescData = file_proto_line_food_proto_rawDesc
)

func file_proto_line_food_proto_rawDescGZIP() []byte {
	file_proto_line_food_proto_rawDescOnce.Do(func() {
		file_proto_line_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_line_food_proto_rawDescData)
	})
	return file_proto_line_food_proto_rawDescData
}

var file_proto_line_food_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_line_food_proto_goTypes = []interface{}{
	(*LineFood)(nil),                 // 0: proto.LineFood
	(*ListLineFoodRequest)(nil),      // 1: proto.ListLineFoodRequest
	(*ListLineFoodResponse)(nil),     // 2: proto.ListLineFoodResponse
	(*CreateLineFoodRequest)(nil),    // 3: proto.CreateLineFoodRequest
	(*CreateLineFoodResponse)(nil),   // 4: proto.CreateLineFoodResponse
	(*ReplaceLineFoodRequest)(nil),   // 5: proto.ReplaceLineFoodRequest
	(*ReplaceLineFoodResponse)(nil),  // 6: proto.ReplaceLineFoodResponse
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*pb_restaurant.Restaurant)(nil), // 8: proto.Restaurant
}
var file_proto_line_food_proto_depIdxs = []int32{
	7, // 0: proto.LineFood.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: proto.LineFood.updated_at:type_name -> google.protobuf.Timestamp
	8, // 2: proto.ListLineFoodResponse.restaurant:type_name -> proto.Restaurant
	0, // 3: proto.CreateLineFoodResponse.LineFood:type_name -> proto.LineFood
	0, // 4: proto.ReplaceLineFoodResponse.line_foods:type_name -> proto.LineFood
	1, // 5: proto.LineFoodService.ListLineFoods:input_type -> proto.ListLineFoodRequest
	3, // 6: proto.LineFoodService.CreateLineFood:input_type -> proto.CreateLineFoodRequest
	5, // 7: proto.LineFoodService.ReplaceLineFood:input_type -> proto.ReplaceLineFoodRequest
	2, // 8: proto.LineFoodService.ListLineFoods:output_type -> proto.ListLineFoodResponse
	4, // 9: proto.LineFoodService.CreateLineFood:output_type -> proto.CreateLineFoodResponse
	6, // 10: proto.LineFoodService.ReplaceLineFood:output_type -> proto.ReplaceLineFoodResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_line_food_proto_init() }
func file_proto_line_food_proto_init() {
	if File_proto_line_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_line_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineFood); i {
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
		file_proto_line_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLineFoodRequest); i {
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
		file_proto_line_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLineFoodResponse); i {
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
		file_proto_line_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLineFoodRequest); i {
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
		file_proto_line_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLineFoodResponse); i {
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
		file_proto_line_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceLineFoodRequest); i {
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
		file_proto_line_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceLineFoodResponse); i {
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
			RawDescriptor: file_proto_line_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_line_food_proto_goTypes,
		DependencyIndexes: file_proto_line_food_proto_depIdxs,
		MessageInfos:      file_proto_line_food_proto_msgTypes,
	}.Build()
	File_proto_line_food_proto = out.File
	file_proto_line_food_proto_rawDesc = nil
	file_proto_line_food_proto_goTypes = nil
	file_proto_line_food_proto_depIdxs = nil
}
