// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            []byte                `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	AccountId            string                `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	TotalPrice           float64               `protobuf:"fixed64,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Products             []*Order_OrderProduct `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedAt() []byte {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Order) GetTotalPrice() float64 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Order) GetProducts() []*Order_OrderProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

type Order_OrderProduct struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Descreption          string   `protobuf:"bytes,3,opt,name=descreption,proto3" json:"descreption,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             uint32   `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order_OrderProduct) Reset()         { *m = Order_OrderProduct{} }
func (m *Order_OrderProduct) String() string { return proto.CompactTextString(m) }
func (*Order_OrderProduct) ProtoMessage()    {}
func (*Order_OrderProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{0, 0}
}
func (m *Order_OrderProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order_OrderProduct.Unmarshal(m, b)
}
func (m *Order_OrderProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order_OrderProduct.Marshal(b, m, deterministic)
}
func (dst *Order_OrderProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order_OrderProduct.Merge(dst, src)
}
func (m *Order_OrderProduct) XXX_Size() int {
	return xxx_messageInfo_Order_OrderProduct.Size(m)
}
func (m *Order_OrderProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_Order_OrderProduct.DiscardUnknown(m)
}

var xxx_messageInfo_Order_OrderProduct proto.InternalMessageInfo

func (m *Order_OrderProduct) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order_OrderProduct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order_OrderProduct) GetDescreption() string {
	if m != nil {
		return m.Descreption
	}
	return ""
}

func (m *Order_OrderProduct) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order_OrderProduct) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type PostOrderRequest struct {
	AccountId            string                           `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Products             []*PostOrderRequest_OrderProduct `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PostOrderRequest) Reset()         { *m = PostOrderRequest{} }
func (m *PostOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PostOrderRequest) ProtoMessage()    {}
func (*PostOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{1}
}
func (m *PostOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOrderRequest.Unmarshal(m, b)
}
func (m *PostOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOrderRequest.Marshal(b, m, deterministic)
}
func (dst *PostOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOrderRequest.Merge(dst, src)
}
func (m *PostOrderRequest) XXX_Size() int {
	return xxx_messageInfo_PostOrderRequest.Size(m)
}
func (m *PostOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostOrderRequest proto.InternalMessageInfo

func (m *PostOrderRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PostOrderRequest) GetProducts() []*PostOrderRequest_OrderProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

type PostOrderRequest_OrderProduct struct {
	ProductId            string   `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity             uint32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostOrderRequest_OrderProduct) Reset()         { *m = PostOrderRequest_OrderProduct{} }
func (m *PostOrderRequest_OrderProduct) String() string { return proto.CompactTextString(m) }
func (*PostOrderRequest_OrderProduct) ProtoMessage()    {}
func (*PostOrderRequest_OrderProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{1, 0}
}
func (m *PostOrderRequest_OrderProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOrderRequest_OrderProduct.Unmarshal(m, b)
}
func (m *PostOrderRequest_OrderProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOrderRequest_OrderProduct.Marshal(b, m, deterministic)
}
func (dst *PostOrderRequest_OrderProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOrderRequest_OrderProduct.Merge(dst, src)
}
func (m *PostOrderRequest_OrderProduct) XXX_Size() int {
	return xxx_messageInfo_PostOrderRequest_OrderProduct.Size(m)
}
func (m *PostOrderRequest_OrderProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOrderRequest_OrderProduct.DiscardUnknown(m)
}

var xxx_messageInfo_PostOrderRequest_OrderProduct proto.InternalMessageInfo

func (m *PostOrderRequest_OrderProduct) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *PostOrderRequest_OrderProduct) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type PostOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostOrderResponse) Reset()         { *m = PostOrderResponse{} }
func (m *PostOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PostOrderResponse) ProtoMessage()    {}
func (*PostOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{2}
}
func (m *PostOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOrderResponse.Unmarshal(m, b)
}
func (m *PostOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOrderResponse.Marshal(b, m, deterministic)
}
func (dst *PostOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOrderResponse.Merge(dst, src)
}
func (m *PostOrderResponse) XXX_Size() int {
	return xxx_messageInfo_PostOrderResponse.Size(m)
}
func (m *PostOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostOrderResponse proto.InternalMessageInfo

func (m *PostOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type GetOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequest) Reset()         { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()    {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{3}
}
func (m *GetOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequest.Unmarshal(m, b)
}
func (m *GetOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequest.Marshal(b, m, deterministic)
}
func (dst *GetOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequest.Merge(dst, src)
}
func (m *GetOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequest.Size(m)
}
func (m *GetOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequest proto.InternalMessageInfo

func (m *GetOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderResponse) Reset()         { *m = GetOrderResponse{} }
func (m *GetOrderResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderResponse) ProtoMessage()    {}
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{4}
}
func (m *GetOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResponse.Unmarshal(m, b)
}
func (m *GetOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResponse.Marshal(b, m, deterministic)
}
func (dst *GetOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResponse.Merge(dst, src)
}
func (m *GetOrderResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderResponse.Size(m)
}
func (m *GetOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResponse proto.InternalMessageInfo

func (m *GetOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type GetOrdersForAccountRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrdersForAccountRequest) Reset()         { *m = GetOrdersForAccountRequest{} }
func (m *GetOrdersForAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrdersForAccountRequest) ProtoMessage()    {}
func (*GetOrdersForAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{5}
}
func (m *GetOrdersForAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrdersForAccountRequest.Unmarshal(m, b)
}
func (m *GetOrdersForAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrdersForAccountRequest.Marshal(b, m, deterministic)
}
func (dst *GetOrdersForAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrdersForAccountRequest.Merge(dst, src)
}
func (m *GetOrdersForAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrdersForAccountRequest.Size(m)
}
func (m *GetOrdersForAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrdersForAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrdersForAccountRequest proto.InternalMessageInfo

func (m *GetOrdersForAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type GetOrdersForAccountResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrdersForAccountResponse) Reset()         { *m = GetOrdersForAccountResponse{} }
func (m *GetOrdersForAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrdersForAccountResponse) ProtoMessage()    {}
func (*GetOrdersForAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_01c41da7c14ef50a, []int{6}
}
func (m *GetOrdersForAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrdersForAccountResponse.Unmarshal(m, b)
}
func (m *GetOrdersForAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrdersForAccountResponse.Marshal(b, m, deterministic)
}
func (dst *GetOrdersForAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrdersForAccountResponse.Merge(dst, src)
}
func (m *GetOrdersForAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrdersForAccountResponse.Size(m)
}
func (m *GetOrdersForAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrdersForAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrdersForAccountResponse proto.InternalMessageInfo

func (m *GetOrdersForAccountResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*Order_OrderProduct)(nil), "pb.Order.OrderProduct")
	proto.RegisterType((*PostOrderRequest)(nil), "pb.PostOrderRequest")
	proto.RegisterType((*PostOrderRequest_OrderProduct)(nil), "pb.PostOrderRequest.OrderProduct")
	proto.RegisterType((*PostOrderResponse)(nil), "pb.PostOrderResponse")
	proto.RegisterType((*GetOrderRequest)(nil), "pb.GetOrderRequest")
	proto.RegisterType((*GetOrderResponse)(nil), "pb.GetOrderResponse")
	proto.RegisterType((*GetOrdersForAccountRequest)(nil), "pb.GetOrdersForAccountRequest")
	proto.RegisterType((*GetOrdersForAccountResponse)(nil), "pb.GetOrdersForAccountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	PostOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error)
	GetOrdersForAccount(ctx context.Context, in *GetOrdersForAccountRequest, opts ...grpc.CallOption) (*GetOrdersForAccountResponse, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) PostOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error) {
	out := new(PostOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/PostOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrdersForAccount(ctx context.Context, in *GetOrdersForAccountRequest, opts ...grpc.CallOption) (*GetOrdersForAccountResponse, error) {
	out := new(GetOrdersForAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/GetOrdersForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	PostOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error)
	GetOrdersForAccount(context.Context, *GetOrdersForAccountRequest) (*GetOrdersForAccountResponse, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_PostOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PostOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/PostOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PostOrder(ctx, req.(*PostOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrdersForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersForAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrdersForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/GetOrdersForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrdersForAccount(ctx, req.(*GetOrdersForAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostOrder",
			Handler:    _OrderService_PostOrder_Handler,
		},
		{
			MethodName: "GetOrdersForAccount",
			Handler:    _OrderService_GetOrdersForAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_order_01c41da7c14ef50a) }

var fileDescriptor_order_01c41da7c14ef50a = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6e, 0x9b, 0x30,
	0x1c, 0x9e, 0x49, 0x88, 0xc2, 0x8f, 0x6c, 0xcb, 0xbc, 0x6c, 0x42, 0x2c, 0x4a, 0x08, 0x27, 0x4e,
	0x39, 0x90, 0x9d, 0x22, 0x4d, 0x5a, 0x2e, 0xfb, 0x73, 0x6a, 0xe4, 0x5e, 0x7a, 0x25, 0xe0, 0x03,
	0x52, 0x8b, 0x89, 0x6d, 0x2a, 0xf5, 0x05, 0xfa, 0x24, 0x7d, 0x80, 0x5e, 0xfb, 0x76, 0x15, 0x86,
	0x00, 0x71, 0x89, 0xd4, 0x4b, 0x14, 0x7f, 0x9f, 0xed, 0xef, 0xcf, 0xcf, 0x80, 0xcd, 0x78, 0x42,
	0xf9, 0x3a, 0xe7, 0x4c, 0x32, 0x6c, 0xe4, 0x07, 0xff, 0xd9, 0x00, 0xf3, 0xaa, 0xc4, 0xf0, 0x27,
	0x30, 0xd2, 0xc4, 0x41, 0x1e, 0x0a, 0x2c, 0x62, 0xa4, 0x09, 0x9e, 0x83, 0x15, 0x73, 0x1a, 0x49,
	0x9a, 0xec, 0xa4, 0x63, 0x78, 0x28, 0x98, 0x90, 0x16, 0x28, 0xd9, 0x28, 0x8e, 0x59, 0x91, 0xc9,
	0xff, 0x89, 0x33, 0x50, 0x87, 0x5a, 0x00, 0x2f, 0x00, 0x24, 0x93, 0xd1, 0xed, 0x9e, 0xa7, 0x31,
	0x75, 0x86, 0x1e, 0x0a, 0x10, 0xe9, 0x20, 0x38, 0x84, 0x71, 0xce, 0x59, 0x52, 0xc4, 0x52, 0x38,
	0xa6, 0x37, 0x08, 0xec, 0xf0, 0xfb, 0x3a, 0x3f, 0xac, 0x95, 0x91, 0xea, 0x77, 0x5f, 0xd1, 0xa4,
	0xd9, 0xe7, 0x3e, 0x22, 0x98, 0x74, 0xa9, 0x37, 0x86, 0x31, 0x0c, 0xb3, 0xe8, 0x8e, 0x2a, 0xaf,
	0x16, 0x51, 0xff, 0xb1, 0x07, 0x76, 0x42, 0x45, 0xcc, 0x69, 0x2e, 0x53, 0x96, 0xd5, 0x46, 0xbb,
	0x10, 0x9e, 0x81, 0x99, 0x77, 0x5c, 0x56, 0x0b, 0xec, 0xc2, 0xf8, 0x58, 0x44, 0x99, 0x4c, 0xe5,
	0x83, 0x63, 0x7a, 0x28, 0xf8, 0x48, 0x9a, 0xb5, 0xff, 0x82, 0x60, 0xba, 0x67, 0x42, 0x2a, 0x33,
	0x84, 0x1e, 0x0b, 0x2a, 0xb4, 0x3e, 0x0c, 0xbd, 0x8f, 0x5f, 0x9d, 0xbc, 0x43, 0x95, 0x77, 0x55,
	0xe6, 0xd5, 0x6f, 0xb9, 0x14, 0xfd, 0x9f, 0x96, 0x7c, 0x0e, 0x56, 0xcd, 0xb5, 0x62, 0x0d, 0x70,
	0xe6, 0x7d, 0xa0, 0x79, 0xff, 0x09, 0x5f, 0x3a, 0xa2, 0x22, 0x67, 0x99, 0xa0, 0x78, 0x09, 0xa6,
	0x7a, 0x16, 0xaa, 0x4b, 0x3b, 0xb4, 0x9a, 0x51, 0x90, 0x0a, 0xf7, 0x57, 0xf0, 0xf9, 0x2f, 0x3d,
	0xcf, 0xab, 0x95, 0xef, 0x6f, 0x60, 0xda, 0x6e, 0x79, 0xef, 0xbd, 0x5b, 0x70, 0x4f, 0x87, 0xc4,
	0x1f, 0xc6, 0x77, 0x55, 0x5f, 0xbd, 0x95, 0x22, 0xad, 0x52, 0xff, 0x37, 0xfc, 0xe8, 0x3d, 0x5b,
	0x6b, 0xaf, 0x60, 0xa4, 0x34, 0x84, 0x83, 0x54, 0xdf, 0x1d, 0xf1, 0x9a, 0x08, 0x9f, 0x4e, 0x0f,
	0xea, 0x9a, 0xf2, 0xfb, 0x72, 0xe8, 0x5b, 0xb0, 0x9a, 0x72, 0xf0, 0xac, 0x6f, 0x40, 0xee, 0x37,
	0x0d, 0xad, 0xd4, 0xfc, 0x0f, 0xf8, 0x06, 0xbe, 0xf6, 0xd8, 0xc1, 0x8b, 0x72, 0xff, 0xe5, 0x8c,
	0xee, 0xf2, 0x22, 0x7f, 0xba, 0xf9, 0x30, 0x52, 0x1f, 0xeb, 0xe6, 0x35, 0x00, 0x00, 0xff, 0xff,
	0xe1, 0xfc, 0xf0, 0x81, 0xbb, 0x03, 0x00, 0x00,
}