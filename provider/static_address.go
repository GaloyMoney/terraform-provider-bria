package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaStaticAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaStaticAddressCreate,
		Read:   resourceBriaStaticAddressRead,
		Update: resourceBriaStaticAddressUpdate,
		Delete: resourceBriaStaticAddressDelete,

		Schema: map[string]*schema.Schema{
			"wallet": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the wallet.",
			},
			"external_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the address.",
			},
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The address",
			},
		},
	}
}

func resourceBriaStaticAddressCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	external_id := d.Get("external_id").(string)
	wallet := d.Get("wallet").(string)

	resp, err := client.CreateStaticAddress(external_id, wallet)
	if err != nil {
		return fmt.Errorf("error creating Bria address: %w", err)
	}

	d.SetId(resp.Address)
	d.Set("address", resp.Address)

	return resourceBriaStaticAddressRead(d, meta)
}

func resourceBriaStaticAddressRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)
	address := d.Get("address").(string)

	resp, err := client.ReadStaticAddress(address)
	if err != nil {
		return fmt.Errorf("error reading Bria address: %w", err)
	}

	d.SetId(*resp.Address)
	d.Set("external_id", resp.ExternalId)
	d.Set("address", resp.Address)

	return nil
}

func resourceBriaStaticAddressDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)
	address := d.Get("address").(string)

	err := client.DeleteStaticAddress(address)
	if err != nil {
		return fmt.Errorf("error deleting Bria address: %w", err)
	}

	d.SetId("")

	return nil
}

func resourceBriaStaticAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	address := d.Get("address").(string)
	externalId := d.Get("external_id").(string)

	err := client.UpdateStaticAddress(address, externalId)
	if err != nil {
		return fmt.Errorf("error updating Bria address: %w", err)
	}

	return resourceBriaStaticAddressRead(d, meta)
}
