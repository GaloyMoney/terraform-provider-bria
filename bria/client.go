package bria

import (
	"context"
	"encoding/base64"
	"fmt"

	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AccountClient struct {
	conn    *grpc.ClientConn
	service briav1.BriaServiceClient
}

func NewAccountClient(endpoint string, apiKey string) (*AccountClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Inject API key as an HTTP header
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, "x-bria-api-key", apiKey)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}

	client := briav1.NewBriaServiceClient(conn)

	return &AccountClient{
		conn:    conn,
		service: client,
	}, nil
}

func (c *AccountClient) Close() {
	c.conn.Close()
}

func (c *AccountClient) CreateProfile(name string) (*briav1.CreateProfileResponse, error) {
	req := &briav1.CreateProfileRequest{
		Name: name,
	}
	ctx := context.Background()
	res, err := c.service.CreateProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) ReadProfile(profileID string) (*briav1.Profile, error) {
	ctx := context.Background()

	listProfilesRequest := &briav1.ListProfilesRequest{}
	listProfilesResponse, err := c.service.ListProfiles(ctx, listProfilesRequest)
	if err != nil {
		return nil, fmt.Errorf("error fetching profiles: %w", err)
	}

	var foundProfile *briav1.Profile
	for _, profile := range listProfilesResponse.Profiles {
		if profile.Id == profileID {
			foundProfile = profile
			break
		}
	}

	return foundProfile, nil
}

func (c *AccountClient) CreateApiKey(name string) (*briav1.CreateProfileApiKeyResponse, error) {
	req := &briav1.CreateProfileApiKeyRequest{
		ProfileName: name,
	}
	ctx := context.Background()
	res, err := c.service.CreateProfileApiKey(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) ImportXpub(name, xpub, derivation string) (*briav1.ImportXpubResponse, error) {
	req := &briav1.ImportXpubRequest{
		Name:       name,
		Xpub:       xpub,
		Derivation: derivation,
	}
	ctx := context.Background()
	res, err := c.service.ImportXpub(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) SetLndSignerConfig(xpub string, lndConfig []interface{}) error {
	lnd := lndConfig[0].(map[string]interface{})

	certBase64 := base64.StdEncoding.EncodeToString([]byte(lnd["cert"].(string)))

	lndSignerConfig := &briav1.LndSignerConfig{
		Endpoint:       lnd["endpoint"].(string),
		CertBase64:     certBase64,
		MacaroonBase64: lnd["macaroon_base64"].(string),
	}

	req := &briav1.SetSignerConfigRequest{
		XpubRef: xpub,
		Config: &briav1.SetSignerConfigRequest_Lnd{
			Lnd: lndSignerConfig,
		},
	}

	_, err := c.service.SetSignerConfig(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error setting LND signer config: %w", err)
	}

	return nil
}

func (c *AccountClient) SetBitcoindSignerConfig(xpub string, bitcoindConfig []interface{}) error {
	bitcoind := bitcoindConfig[0].(map[string]interface{})

	bitcoindSignerConfig := &briav1.BitcoindSignerConfig{
		Endpoint:    bitcoind["endpoint"].(string),
		RpcUser:     bitcoind["rpc_user"].(string),
		RpcPassword: bitcoind["rpc_password"].(string),
	}

	req := &briav1.SetSignerConfigRequest{
		XpubRef: xpub,
		Config: &briav1.SetSignerConfigRequest_Bitcoind{
			Bitcoind: bitcoindSignerConfig,
		},
	}

	_, err := c.service.SetSignerConfig(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error setting LND signer config: %w", err)
	}

	return nil
}

func (c *AccountClient) CreateWallet(name string, keychainConfig *briav1.KeychainConfig) (*briav1.CreateWalletResponse, error) {
	req := &briav1.CreateWalletRequest{
		Name:           name,
		KeychainConfig: keychainConfig,
	}
	ctx := context.Background()
	res, err := c.service.CreateWallet(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) CreatePayoutQueue(name string, description string, config *briav1.PayoutQueueConfig) (*briav1.CreatePayoutQueueResponse, error) {
	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}
	req := &briav1.CreatePayoutQueueRequest{
		Name:        name,
		Description: descriptionPtr,
		Config:      config,
	}
	ctx := context.Background()
	res, err := c.service.CreatePayoutQueue(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AccountClient) ReadPayoutQueue(queueId string) (*briav1.PayoutQueue, error) {
	ctx := context.Background()

	listPayoutQueuesRequest := &briav1.ListPayoutQueuesRequest{}
	listPayoutQueuesResponse, err := c.service.ListPayoutQueues(ctx, listPayoutQueuesRequest)
	if err != nil {
		return nil, fmt.Errorf("error fetching queues: %w", err)
	}

	var foundPayoutQueue *briav1.PayoutQueue
	for _, queue := range listPayoutQueuesResponse.PayoutQueues {
		if queue.Id == queueId {
			foundPayoutQueue = queue
			break
		}
	}

	return foundPayoutQueue, nil
}

func (c *AccountClient) UpdatePayoutQueue(id string, description string, config *briav1.PayoutQueueConfig) (*briav1.UpdatePayoutQueueResponse, error) {
	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}
	req := &briav1.UpdatePayoutQueueRequest{
		Id:             id,
		NewDescription: descriptionPtr,
		NewConfig:      config,
	}
	ctx := context.Background()
	res, err := c.service.UpdatePayoutQueue(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
