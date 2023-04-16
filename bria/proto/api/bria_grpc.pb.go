// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/bria.proto

package briav1

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
	BriaService_CreateProfile_FullMethodName           = "/services.bria.v1.BriaService/CreateProfile"
	BriaService_ListProfiles_FullMethodName            = "/services.bria.v1.BriaService/ListProfiles"
	BriaService_CreateProfileApiKey_FullMethodName     = "/services.bria.v1.BriaService/CreateProfileApiKey"
	BriaService_ImportXpub_FullMethodName              = "/services.bria.v1.BriaService/ImportXpub"
	BriaService_SetSignerConfig_FullMethodName         = "/services.bria.v1.BriaService/SetSignerConfig"
	BriaService_CreateWallet_FullMethodName            = "/services.bria.v1.BriaService/CreateWallet"
	BriaService_GetWalletBalanceSummary_FullMethodName = "/services.bria.v1.BriaService/GetWalletBalanceSummary"
	BriaService_NewAddress_FullMethodName              = "/services.bria.v1.BriaService/NewAddress"
	BriaService_ListUtxos_FullMethodName               = "/services.bria.v1.BriaService/ListUtxos"
	BriaService_CreateBatchGroup_FullMethodName        = "/services.bria.v1.BriaService/CreateBatchGroup"
	BriaService_QueuePayout_FullMethodName             = "/services.bria.v1.BriaService/QueuePayout"
	BriaService_ListPayouts_FullMethodName             = "/services.bria.v1.BriaService/ListPayouts"
)

// BriaServiceClient is the client API for BriaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BriaServiceClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error)
	ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error)
	CreateProfileApiKey(ctx context.Context, in *CreateProfileApiKeyRequest, opts ...grpc.CallOption) (*CreateProfileApiKeyResponse, error)
	ImportXpub(ctx context.Context, in *ImportXpubRequest, opts ...grpc.CallOption) (*ImportXpubResponse, error)
	SetSignerConfig(ctx context.Context, in *SetSignerConfigRequest, opts ...grpc.CallOption) (*SetSignerConfigResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	GetWalletBalanceSummary(ctx context.Context, in *GetWalletBalanceSummaryRequest, opts ...grpc.CallOption) (*GetWalletBalanceSummaryResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error)
	CreateBatchGroup(ctx context.Context, in *CreateBatchGroupRequest, opts ...grpc.CallOption) (*CreateBatchGroupResponse, error)
	QueuePayout(ctx context.Context, in *QueuePayoutRequest, opts ...grpc.CallOption) (*QueuePayoutResponse, error)
	ListPayouts(ctx context.Context, in *ListPayoutsRequest, opts ...grpc.CallOption) (*ListPayoutsResponse, error)
}

type briaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBriaServiceClient(cc grpc.ClientConnInterface) BriaServiceClient {
	return &briaServiceClient{cc}
}

func (c *briaServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error) {
	out := new(CreateProfileResponse)
	err := c.cc.Invoke(ctx, BriaService_CreateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error) {
	out := new(ListProfilesResponse)
	err := c.cc.Invoke(ctx, BriaService_ListProfiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) CreateProfileApiKey(ctx context.Context, in *CreateProfileApiKeyRequest, opts ...grpc.CallOption) (*CreateProfileApiKeyResponse, error) {
	out := new(CreateProfileApiKeyResponse)
	err := c.cc.Invoke(ctx, BriaService_CreateProfileApiKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) ImportXpub(ctx context.Context, in *ImportXpubRequest, opts ...grpc.CallOption) (*ImportXpubResponse, error) {
	out := new(ImportXpubResponse)
	err := c.cc.Invoke(ctx, BriaService_ImportXpub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) SetSignerConfig(ctx context.Context, in *SetSignerConfigRequest, opts ...grpc.CallOption) (*SetSignerConfigResponse, error) {
	out := new(SetSignerConfigResponse)
	err := c.cc.Invoke(ctx, BriaService_SetSignerConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, BriaService_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) GetWalletBalanceSummary(ctx context.Context, in *GetWalletBalanceSummaryRequest, opts ...grpc.CallOption) (*GetWalletBalanceSummaryResponse, error) {
	out := new(GetWalletBalanceSummaryResponse)
	err := c.cc.Invoke(ctx, BriaService_GetWalletBalanceSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, BriaService_NewAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error) {
	out := new(ListUtxosResponse)
	err := c.cc.Invoke(ctx, BriaService_ListUtxos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) CreateBatchGroup(ctx context.Context, in *CreateBatchGroupRequest, opts ...grpc.CallOption) (*CreateBatchGroupResponse, error) {
	out := new(CreateBatchGroupResponse)
	err := c.cc.Invoke(ctx, BriaService_CreateBatchGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) QueuePayout(ctx context.Context, in *QueuePayoutRequest, opts ...grpc.CallOption) (*QueuePayoutResponse, error) {
	out := new(QueuePayoutResponse)
	err := c.cc.Invoke(ctx, BriaService_QueuePayout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) ListPayouts(ctx context.Context, in *ListPayoutsRequest, opts ...grpc.CallOption) (*ListPayoutsResponse, error) {
	out := new(ListPayoutsResponse)
	err := c.cc.Invoke(ctx, BriaService_ListPayouts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BriaServiceServer is the server API for BriaService service.
// All implementations must embed UnimplementedBriaServiceServer
// for forward compatibility
type BriaServiceServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
	ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error)
	CreateProfileApiKey(context.Context, *CreateProfileApiKeyRequest) (*CreateProfileApiKeyResponse, error)
	ImportXpub(context.Context, *ImportXpubRequest) (*ImportXpubResponse, error)
	SetSignerConfig(context.Context, *SetSignerConfigRequest) (*SetSignerConfigResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWalletBalanceSummary(context.Context, *GetWalletBalanceSummaryRequest) (*GetWalletBalanceSummaryResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error)
	CreateBatchGroup(context.Context, *CreateBatchGroupRequest) (*CreateBatchGroupResponse, error)
	QueuePayout(context.Context, *QueuePayoutRequest) (*QueuePayoutResponse, error)
	ListPayouts(context.Context, *ListPayoutsRequest) (*ListPayoutsResponse, error)
	mustEmbedUnimplementedBriaServiceServer()
}

// UnimplementedBriaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBriaServiceServer struct {
}

func (UnimplementedBriaServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedBriaServiceServer) ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfiles not implemented")
}
func (UnimplementedBriaServiceServer) CreateProfileApiKey(context.Context, *CreateProfileApiKeyRequest) (*CreateProfileApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfileApiKey not implemented")
}
func (UnimplementedBriaServiceServer) ImportXpub(context.Context, *ImportXpubRequest) (*ImportXpubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportXpub not implemented")
}
func (UnimplementedBriaServiceServer) SetSignerConfig(context.Context, *SetSignerConfigRequest) (*SetSignerConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignerConfig not implemented")
}
func (UnimplementedBriaServiceServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedBriaServiceServer) GetWalletBalanceSummary(context.Context, *GetWalletBalanceSummaryRequest) (*GetWalletBalanceSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletBalanceSummary not implemented")
}
func (UnimplementedBriaServiceServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedBriaServiceServer) ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtxos not implemented")
}
func (UnimplementedBriaServiceServer) CreateBatchGroup(context.Context, *CreateBatchGroupRequest) (*CreateBatchGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatchGroup not implemented")
}
func (UnimplementedBriaServiceServer) QueuePayout(context.Context, *QueuePayoutRequest) (*QueuePayoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueuePayout not implemented")
}
func (UnimplementedBriaServiceServer) ListPayouts(context.Context, *ListPayoutsRequest) (*ListPayoutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayouts not implemented")
}
func (UnimplementedBriaServiceServer) mustEmbedUnimplementedBriaServiceServer() {}

// UnsafeBriaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BriaServiceServer will
// result in compilation errors.
type UnsafeBriaServiceServer interface {
	mustEmbedUnimplementedBriaServiceServer()
}

func RegisterBriaServiceServer(s grpc.ServiceRegistrar, srv BriaServiceServer) {
	s.RegisterService(&BriaService_ServiceDesc, srv)
}

func _BriaService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_ListProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListProfiles(ctx, req.(*ListProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_CreateProfileApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).CreateProfileApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_CreateProfileApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).CreateProfileApiKey(ctx, req.(*CreateProfileApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_ImportXpub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportXpubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ImportXpub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ImportXpub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ImportXpub(ctx, req.(*ImportXpubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_SetSignerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSignerConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).SetSignerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_SetSignerConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).SetSignerConfig(ctx, req.(*SetSignerConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_GetWalletBalanceSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletBalanceSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).GetWalletBalanceSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_GetWalletBalanceSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).GetWalletBalanceSummary(ctx, req.(*GetWalletBalanceSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_NewAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_ListUtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListUtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListUtxos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListUtxos(ctx, req.(*ListUtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_CreateBatchGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBatchGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).CreateBatchGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_CreateBatchGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).CreateBatchGroup(ctx, req.(*CreateBatchGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_QueuePayout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueuePayoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).QueuePayout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_QueuePayout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).QueuePayout(ctx, req.(*QueuePayoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_ListPayouts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPayoutsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListPayouts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListPayouts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListPayouts(ctx, req.(*ListPayoutsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BriaService_ServiceDesc is the grpc.ServiceDesc for BriaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BriaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.bria.v1.BriaService",
	HandlerType: (*BriaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _BriaService_CreateProfile_Handler,
		},
		{
			MethodName: "ListProfiles",
			Handler:    _BriaService_ListProfiles_Handler,
		},
		{
			MethodName: "CreateProfileApiKey",
			Handler:    _BriaService_CreateProfileApiKey_Handler,
		},
		{
			MethodName: "ImportXpub",
			Handler:    _BriaService_ImportXpub_Handler,
		},
		{
			MethodName: "SetSignerConfig",
			Handler:    _BriaService_SetSignerConfig_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _BriaService_CreateWallet_Handler,
		},
		{
			MethodName: "GetWalletBalanceSummary",
			Handler:    _BriaService_GetWalletBalanceSummary_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _BriaService_NewAddress_Handler,
		},
		{
			MethodName: "ListUtxos",
			Handler:    _BriaService_ListUtxos_Handler,
		},
		{
			MethodName: "CreateBatchGroup",
			Handler:    _BriaService_CreateBatchGroup_Handler,
		},
		{
			MethodName: "QueuePayout",
			Handler:    _BriaService_QueuePayout_Handler,
		},
		{
			MethodName: "ListPayouts",
			Handler:    _BriaService_ListPayouts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bria.proto",
}
