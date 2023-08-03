package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaWallet() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaWalletCreate,
		Update: resourceBriaWalletUpdate,
		Read:   resourceBriaWalletRead,
		Delete: resourceBriaWalletSoftDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the wallet.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the wallet.",
			},
			"keychain": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The keychain configuration for the wallet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wpkh": {
							Type:        schema.TypeList,
							Elem:        wpkhConfig(),
							Optional:    true,
							MaxItems:    1,
							Description: "A list of xpub reference IDs associated with the wallet.",
						},
						"descriptors": {
							Type:        schema.TypeList,
							Elem:        descriptorsConfig(),
							Optional:    true,
							MaxItems:    1,
							Description: "Descriptor configuration for the wallet.",
						},
					}}},
		},
	}
}

func wpkhConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"xpub": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The xpub-ref or xpub",
			},
			"derivation_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The derivation path for the xpub (if it is not a reference)",
			},
		},
	}
}

func descriptorsConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The external descriptor",
			},
			"internal": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The internal descriptor",
			},
		},
	}
}

func resourceBriaWalletCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	name := d.Get("name").(string)

	var keychainConfig *briav1.KeychainConfig

	// handle error if needed
	if keychainRaw, ok := d.GetOk("keychain"); ok {
		keychainList := keychainRaw.([]interface{})
		keychainMap := keychainList[0].(map[string]interface{})

		wpkhRaw, wpkhOk := keychainMap["wpkh"]
		descriptorsRaw, descriptorsOk := keychainMap["descriptors"]

		if wpkhOk && len(wpkhRaw.([]interface{})) > 0 {
			wpkhList := wpkhRaw.([]interface{})
			wpkhMap := wpkhList[0].(map[string]interface{})
			xpub := wpkhMap["xpub"].(string)
			derivationPath := wpkhMap["derivation_path"].(string)
			keychainConfig = &briav1.KeychainConfig{
				Config: &briav1.KeychainConfig_Wpkh_{
					Wpkh: &briav1.KeychainConfig_Wpkh{
						Xpub:           xpub,
						DerivationPath: &derivationPath,
					},
				},
			}
		} else if descriptorsOk && len(descriptorsRaw.([]interface{})) > 0 {
			descriptorsList := descriptorsRaw.([]interface{})
			descriptorsMap := descriptorsList[0].(map[string]interface{})
			external := descriptorsMap["external"].(string)
			internal := descriptorsMap["internal"].(string)
			keychainConfig = &briav1.KeychainConfig{
				Config: &briav1.KeychainConfig_Descriptors_{
					Descriptors: &briav1.KeychainConfig_Descriptors{
						External: external,
						Internal: internal,
					},
				},
			}
		}
	}

	fmt.Println(keychainConfig)
	resp, err := client.CreateWallet(name, keychainConfig)
	if err != nil {
		return fmt.Errorf("error creating Bria wallet: %w", err)
	}

	d.SetId(resp.Id)

	return resourceBriaWalletRead(d, meta)
}

func resourceBriaWalletRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_wallet resource
	// This can be a no-op if there is no way to read a wallet from the API
	return nil
}

func resourceBriaWalletUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("briaaccount_wallet resource does not support updates")
}

func resourceBriaWalletSoftDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the soft delete function for the bria_account_wallet resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}
