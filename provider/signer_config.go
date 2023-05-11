package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaSignerConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaSignerConfigCreate,
		Update: resourceBriaSignerConfigUpdate,
		Read:   resourceBriaSignerConfigRead,
		Delete: resourceBriaSignerConfigDelete,

		Schema: map[string]*schema.Schema{
			"xpub": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the xpub.",
			},
			"lnd": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				Elem:          lndConfigElem(),
				Description:   "LND signer configuration.",
				ConflictsWith: []string{"bitcoind"},
				AtLeastOneOf:  []string{"lnd", "bitcoind"},
			},
			"bitcoind": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				Elem:          bitcoindConfigElem(),
				Description:   "BitcoinD signer configuration.",
				ConflictsWith: []string{"lnd"},
				AtLeastOneOf:  []string{"lnd", "bitcoind"},
			},
		},
	}
}

func lndConfigElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The LND endpoint.",
			},
			"macaroon_base64": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The base64 encoded macaroon.",
				Sensitive:   true,
			},
			"cert": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The LND certificate.",
				Sensitive:   true,
			},
		},
	}
}

func bitcoindConfigElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Bitcoind endpoint.",
			},
			"rpc_user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user for the Bitcoind RPC.",
				Sensitive:   true,
			},
			"rpc_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password for the Bitcoind RPC.",
				Sensitive:   true,
			},
		},
	}
}

func resourceBriaSignerConfigCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	xpub := d.Get("xpub").(string)

	var err error
	if lndConfig, ok := d.GetOk("lnd"); ok {
		err = client.SetLndSignerConfig(xpub, lndConfig.([]interface{}))
	} else if bitcoindConfig, ok := d.GetOk("bitcoind"); ok {
		err = client.SetBitcoindSignerConfig(xpub, bitcoindConfig.([]interface{}))
	} else {
		return fmt.Errorf("lnd block must be provided")
	}

	if err != nil {
		return fmt.Errorf("error creating Bria signer config: %w", err)
	}

	d.SetId(xpub)

	return resourceBriaSignerConfigRead(d, meta)
}

func resourceBriaSignerConfigRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_signer_config resource
	// This can be a no-op if there is no way to read a signer config from the API
	return nil
}

func resourceBriaSignerConfigDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the delete function for the bria_account_signer_config resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}

func resourceBriaSignerConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	xpub := d.Get("xpub").(string)

	var err error
	if lndConfig, ok := d.GetOk("lnd"); ok {
		err = client.SetLndSignerConfig(xpub, lndConfig.([]interface{}))
	} else if bitcoindConfig, ok := d.GetOk("bitcoind"); ok {
		err = client.SetBitcoindSignerConfig(xpub, bitcoindConfig.([]interface{}))
	} else {
		return fmt.Errorf("lnd block must be provided")
	}

	if err != nil {
		return fmt.Errorf("error creating Bria signer config: %w", err)
	}

	d.SetId(xpub)

	return resourceBriaSignerConfigRead(d, meta)
}
