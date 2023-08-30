package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaProfileCreate,
		Read:   resourceBriaProfileRead,
		Update: resourceBriaProfileUpdate,
		Delete: resourceBriaProfileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the profile.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the profile.",
			},
			"spending_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_payout_addresses": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceBriaProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	profileName := d.Get("name").(string)
	spPolicyRaw := d.Get("spending_policy").([]interface{})
	var spendingPolicy *briav1.SpendingPolicy
	if len(spPolicyRaw) > 0 {
		spMap := spPolicyRaw[0].(map[string]interface{})
		allowedPayoutAddresses := make([]string, 0)
		for _, v := range spMap["allowed_payout_addresses"].(*schema.Set).List() {
			allowedPayoutAddresses = append(allowedPayoutAddresses, v.(string))
		}
		spendingPolicy = &briav1.SpendingPolicy{
			AllowedPayoutAddresses: allowedPayoutAddresses,
		}
	}

	resp, err := client.CreateProfile(profileName, spendingPolicy)
	if err != nil {
		return fmt.Errorf("error creating Bria profile: %w", err)
	}

	d.SetId(resp.Id)

	return resourceBriaProfileRead(d, meta)
}

func resourceBriaProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	profileID := d.Id()

	profile, err := client.ReadProfile(profileID)
	if err != nil {
		return fmt.Errorf("error reading Bria profile: %w", err)
	}

	if profile == nil {
		// Profile was deleted
		d.SetId("")
		return nil
	}

	d.Set("id", profile.Id)
	d.Set("name", profile.Name)
	if profile.SpendingPolicy != nil {
		sp := make([]map[string]interface{}, 1)
		sp[0] = map[string]interface{}{
			"allowed_payout_addresses": profile.SpendingPolicy.AllowedPayoutAddresses,
		}
		d.Set("spending_policy", sp)
	}

	return nil
}

func resourceBriaProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("bria_profile resource does not support updates")
}

func resourceBriaProfileDelete(d *schema.ResourceData, meta interface{}) error {
	// Soft delete: just remove the profile from the Terraform state
	d.SetId("")
	return nil
}
