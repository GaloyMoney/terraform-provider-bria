package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaApiKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaApiKeyCreate,
		Read:   resourceBriaApiKeyRead,
		Update: resourceBriaApiKeyUpdate,
		Delete: resourceBriaApiKeyDelete,

		Schema: map[string]*schema.Schema{
			"profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the profile.",
			},
			"key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The secret key.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the key.",
			},
		},
	}
}

func resourceBriaApiKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	profileName := d.Get("profile").(string)

	resp, err := client.CreateApiKey(profileName)
	if err != nil {
		return fmt.Errorf("error creating Bria api key: %w", err)
	}

	d.SetId(resp.Id)
	d.Set("key", resp.Key)

	return resourceBriaApiKeyRead(d, meta)
}

func resourceBriaApiKeyRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_wallet resource
	// This can be a no-op if there is no way to read a wallet from the API
	return nil
}

func resourceBriaApiKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("bria_profile resource does not support updates")
}

func resourceBriaApiKeyDelete(d *schema.ResourceData, meta interface{}) error {
	// Soft delete: just remove the profile from the Terraform state
	d.SetId("")
	return nil
}
