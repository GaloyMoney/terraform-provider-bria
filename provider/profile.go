package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
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
		},
	}
}

func resourceBriaProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	profileName := d.Get("name").(string)

	resp, err := client.CreateProfile(profileName)
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

	d.Set("name", profile.Name)
	d.Set("id", profile.Id)

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
