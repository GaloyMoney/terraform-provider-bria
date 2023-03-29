package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBatchGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceBatchGroupCreate,
		Read:   resourceBatchGroupRead,
		Update: resourceBatchGroupUpdate,
		Delete: resourceBatchGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_priority": {
							Type:     schema.TypeString,
							Required: true,
						},
						"consolidate_deprecated_keychains": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"trigger": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval_secs": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceBatchGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*bria.AccountClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	configData := d.Get("config").([]interface{})[0].(map[string]interface{})

	config := &briav1.BatchGroupConfig{
		TxPriority:                     briav1.TxPriority(briav1.TxPriority_value[configData["tx_priority"].(string)]),
		ConsolidateDeprecatedKeychains: configData["consolidate_deprecated_keychains"].(bool),
	}

	trigger := configData["trigger"].(string)
	switch trigger {
	case "manual":
		config.Trigger = &briav1.BatchGroupConfig_Manual{Manual: true}
	case "immediate":
		config.Trigger = &briav1.BatchGroupConfig_Immediate{Immediate: true}
	default:
		config.Trigger = &briav1.BatchGroupConfig_IntervalSecs{IntervalSecs: uint32(configData["interval_secs"].(int))}
	}

	res, err := client.CreateBatchGroup(name, description, config)
	if err != nil {
		return fmt.Errorf("error creating Bria batch group: %w", err)
	}

	d.SetId(res.Id)

	return resourceBatchGroupRead(d, m)
}

func resourceBatchGroupRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_signer_config resource
	// This can be a no-op if there is no way to read a signer config from the API
	return nil
}

func resourceBatchGroupDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the delete function for the bria_account_signer_config resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}

func resourceBatchGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("bri_update resource does not support updates")
}
