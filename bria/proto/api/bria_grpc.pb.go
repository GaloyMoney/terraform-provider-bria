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
	BriaService_CreateProfile_FullMethodName            = "/services.bria.v1.BriaService/CreateProfile"
	BriaService_ListProfiles_FullMethodName             = "/services.bria.v1.BriaService/ListProfiles"
	BriaService_CreateProfileApiKey_FullMethodName      = "/services.bria.v1.BriaService/CreateProfileApiKey"
	BriaService_ImportXpub_FullMethodName               = "/services.bria.v1.BriaService/ImportXpub"
	BriaService_SetSignerConfig_FullMethodName          = "/services.bria.v1.BriaService/SetSignerConfig"
	BriaService_CreateWallet_FullMethodName             = "/services.bria.v1.BriaService/CreateWallet"
	BriaService_ListWallets_FullMethodName              = "/services.bria.v1.BriaService/ListWallets"
	BriaService_GetWalletBalanceSummary_FullMethodName  = "/services.bria.v1.BriaService/GetWalletBalanceSummary"
	BriaService_GetAccountBalanceSummary_FullMethodName = "/services.bria.v1.BriaService/GetAccountBalanceSummary"
	BriaService_NewAddress_FullMethodName               = "/services.bria.v1.BriaService/NewAddress"
	BriaService_UpdateAddress_FullMethodName            = "/services.bria.v1.BriaService/UpdateAddress"
	BriaService_ListAddresses_FullMethodName            = "/services.bria.v1.BriaService/ListAddresses"
	BriaService_ListUtxos_FullMethodName                = "/services.bria.v1.BriaService/ListUtxos"
	BriaService_CreateBatchGroup_FullMethodName         = "/services.bria.v1.BriaService/CreateBatchGroup"
	BriaService_ListBatchGroups_FullMethodName          = "/services.bria.v1.BriaService/ListBatchGroups"
	BriaService_QueuePayout_FullMethodName              = "/services.bria.v1.BriaService/QueuePayout"
	BriaService_ListPayouts_FullMethodName              = "/services.bria.v1.BriaService/ListPayouts"
	BriaService_ListSigningSessions_FullMethodName      = "/services.bria.v1.BriaService/ListSigningSessions"
	BriaService_SubscribeAll_FullMethodName             = "/services.bria.v1.BriaService/SubscribeAll"
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
	ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error)
	GetWalletBalanceSummary(ctx context.Context, in *GetWalletBalanceSummaryRequest, opts ...grpc.CallOption) (*GetWalletBalanceSummaryResponse, error)
	GetAccountBalanceSummary(ctx context.Context, in *GetAccountBalanceSummaryRequest, opts ...grpc.CallOption) (*GetAccountBalanceSummaryResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error)
	ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error)
	ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error)
	CreateBatchGroup(ctx context.Context, in *CreateBatchGroupRequest, opts ...grpc.CallOption) (*CreateBatchGroupResponse, error)
	ListBatchGroups(ctx context.Context, in *ListBatchGroupsRequest, opts ...grpc.CallOption) (*ListBatchGroupsResponse, error)
	QueuePayout(ctx context.Context, in *QueuePayoutRequest, opts ...grpc.CallOption) (*QueuePayoutResponse, error)
	ListPayouts(ctx context.Context, in *ListPayoutsRequest, opts ...grpc.CallOption) (*ListPayoutsResponse, error)
	ListSigningSessions(ctx context.Context, in *ListSigningSessionsRequest, opts ...grpc.CallOption) (*ListSigningSessionsResponse, error)
	SubscribeAll(ctx context.Context, in *SubscribeAllRequest, opts ...grpc.CallOption) (BriaService_SubscribeAllClient, error)
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

func (c *briaServiceClient) ListWallets(ctx context.Context, in *ListWalletsRequest, opts ...grpc.CallOption) (*ListWalletsResponse, error) {
	out := new(ListWalletsResponse)
	err := c.cc.Invoke(ctx, BriaService_ListWallets_FullMethodName, in, out, opts...)
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

func (c *briaServiceClient) GetAccountBalanceSummary(ctx context.Context, in *GetAccountBalanceSummaryRequest, opts ...grpc.CallOption) (*GetAccountBalanceSummaryResponse, error) {
	out := new(GetAccountBalanceSummaryResponse)
	err := c.cc.Invoke(ctx, BriaService_GetAccountBalanceSummary_FullMethodName, in, out, opts...)
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

func (c *briaServiceClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error) {
	out := new(UpdateAddressResponse)
	err := c.cc.Invoke(ctx, BriaService_UpdateAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error) {
	out := new(ListAddressesResponse)
	err := c.cc.Invoke(ctx, BriaService_ListAddresses_FullMethodName, in, out, opts...)
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

func (c *briaServiceClient) ListBatchGroups(ctx context.Context, in *ListBatchGroupsRequest, opts ...grpc.CallOption) (*ListBatchGroupsResponse, error) {
	out := new(ListBatchGroupsResponse)
	err := c.cc.Invoke(ctx, BriaService_ListBatchGroups_FullMethodName, in, out, opts...)
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

func (c *briaServiceClient) ListSigningSessions(ctx context.Context, in *ListSigningSessionsRequest, opts ...grpc.CallOption) (*ListSigningSessionsResponse, error) {
	out := new(ListSigningSessionsResponse)
	err := c.cc.Invoke(ctx, BriaService_ListSigningSessions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *briaServiceClient) SubscribeAll(ctx context.Context, in *SubscribeAllRequest, opts ...grpc.CallOption) (BriaService_SubscribeAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &BriaService_ServiceDesc.Streams[0], BriaService_SubscribeAll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &briaServiceSubscribeAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BriaService_SubscribeAllClient interface {
	Recv() (*BriaEvent, error)
	grpc.ClientStream
}

type briaServiceSubscribeAllClient struct {
	grpc.ClientStream
}

func (x *briaServiceSubscribeAllClient) Recv() (*BriaEvent, error) {
	m := new(BriaEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error)
	GetWalletBalanceSummary(context.Context, *GetWalletBalanceSummaryRequest) (*GetWalletBalanceSummaryResponse, error)
	GetAccountBalanceSummary(context.Context, *GetAccountBalanceSummaryRequest) (*GetAccountBalanceSummaryResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error)
	ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error)
	ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error)
	CreateBatchGroup(context.Context, *CreateBatchGroupRequest) (*CreateBatchGroupResponse, error)
	ListBatchGroups(context.Context, *ListBatchGroupsRequest) (*ListBatchGroupsResponse, error)
	QueuePayout(context.Context, *QueuePayoutRequest) (*QueuePayoutResponse, error)
	ListPayouts(context.Context, *ListPayoutsRequest) (*ListPayoutsResponse, error)
	ListSigningSessions(context.Context, *ListSigningSessionsRequest) (*ListSigningSessionsResponse, error)
	SubscribeAll(*SubscribeAllRequest, BriaService_SubscribeAllServer) error
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
func (UnimplementedBriaServiceServer) ListWallets(context.Context, *ListWalletsRequest) (*ListWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallets not implemented")
}
func (UnimplementedBriaServiceServer) GetWalletBalanceSummary(context.Context, *GetWalletBalanceSummaryRequest) (*GetWalletBalanceSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletBalanceSummary not implemented")
}
func (UnimplementedBriaServiceServer) GetAccountBalanceSummary(context.Context, *GetAccountBalanceSummaryRequest) (*GetAccountBalanceSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalanceSummary not implemented")
}
func (UnimplementedBriaServiceServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedBriaServiceServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedBriaServiceServer) ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddresses not implemented")
}
func (UnimplementedBriaServiceServer) ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtxos not implemented")
}
func (UnimplementedBriaServiceServer) CreateBatchGroup(context.Context, *CreateBatchGroupRequest) (*CreateBatchGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatchGroup not implemented")
}
func (UnimplementedBriaServiceServer) ListBatchGroups(context.Context, *ListBatchGroupsRequest) (*ListBatchGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatchGroups not implemented")
}
func (UnimplementedBriaServiceServer) QueuePayout(context.Context, *QueuePayoutRequest) (*QueuePayoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueuePayout not implemented")
}
func (UnimplementedBriaServiceServer) ListPayouts(context.Context, *ListPayoutsRequest) (*ListPayoutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayouts not implemented")
}
func (UnimplementedBriaServiceServer) ListSigningSessions(context.Context, *ListSigningSessionsRequest) (*ListSigningSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSigningSessions not implemented")
}
func (UnimplementedBriaServiceServer) SubscribeAll(*SubscribeAllRequest, BriaService_SubscribeAllServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeAll not implemented")
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

func _BriaService_ListWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListWallets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListWallets(ctx, req.(*ListWalletsRequest))
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

func _BriaService_GetAccountBalanceSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalanceSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).GetAccountBalanceSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_GetAccountBalanceSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).GetAccountBalanceSummary(ctx, req.(*GetAccountBalanceSummaryRequest))
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

func _BriaService_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_UpdateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_ListAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListAddresses(ctx, req.(*ListAddressesRequest))
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

func _BriaService_ListBatchGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListBatchGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListBatchGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListBatchGroups(ctx, req.(*ListBatchGroupsRequest))
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

func _BriaService_ListSigningSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSigningSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BriaServiceServer).ListSigningSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BriaService_ListSigningSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BriaServiceServer).ListSigningSessions(ctx, req.(*ListSigningSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BriaService_SubscribeAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BriaServiceServer).SubscribeAll(m, &briaServiceSubscribeAllServer{stream})
}

type BriaService_SubscribeAllServer interface {
	Send(*BriaEvent) error
	grpc.ServerStream
}

type briaServiceSubscribeAllServer struct {
	grpc.ServerStream
}

func (x *briaServiceSubscribeAllServer) Send(m *BriaEvent) error {
	return x.ServerStream.SendMsg(m)
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
			MethodName: "ListWallets",
			Handler:    _BriaService_ListWallets_Handler,
		},
		{
			MethodName: "GetWalletBalanceSummary",
			Handler:    _BriaService_GetWalletBalanceSummary_Handler,
		},
		{
			MethodName: "GetAccountBalanceSummary",
			Handler:    _BriaService_GetAccountBalanceSummary_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _BriaService_NewAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _BriaService_UpdateAddress_Handler,
		},
		{
			MethodName: "ListAddresses",
			Handler:    _BriaService_ListAddresses_Handler,
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
			MethodName: "ListBatchGroups",
			Handler:    _BriaService_ListBatchGroups_Handler,
		},
		{
			MethodName: "QueuePayout",
			Handler:    _BriaService_QueuePayout_Handler,
		},
		{
			MethodName: "ListPayouts",
			Handler:    _BriaService_ListPayouts_Handler,
		},
		{
			MethodName: "ListSigningSessions",
			Handler:    _BriaService_ListSigningSessions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeAll",
			Handler:       _BriaService_SubscribeAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/bria.proto",
}
