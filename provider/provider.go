package provider

import (
	"github.com/GaloyMoney/terraform-provider-bria/bria"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_API_URL", "localhost:2742"),
				Description: "The API endpoint for Bria.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_API_KEY", ""),
				Description: "The API key for Bria.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bria_xpub":          resourceBriaAccountXpub(),
			"bria_wallet":        resourceBriaAccountWallet(),
			"bria_signer_config": resourceBriaAccountSignerConfig(),
			"bria_batch_group":   resourceBatchGroup(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)

	return bria.NewAccountClient(endpoint, apiKey)
}
